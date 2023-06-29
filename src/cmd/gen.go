/*
Copyright © 2023 Andreas Wachs <andreas@wachs.dk>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/andreaswachs/track/pkg/sheet"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a new time sheet for a given month, optionally for a user defined year",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var month time.Month

		if len(args) > 0 {
			monthName := args[0]
			today, err := time.Parse("January", monthName)
			if err != nil {
				fmt.Println("Invalid month name:", monthName)
				return
			}
			month = today.Month()
		} else {
			now := time.Now()
			month = now.Month()
		}

		year, err := cmd.PersistentFlags().GetInt("year")
		if err != nil || year == 0 {
			year = time.Now().Year()
		}

		day := 1
		t := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

		newSheet := &sheet.Sheet{Entries: sheet.EntriesInMonth(t)}

		buffer, err := yaml.Marshal(newSheet)
		if err != nil {
			fmt.Printf("Error occurred: %s", err)
			os.Exit(1)
		}

		fileName := fmt.Sprintf("%s-%d.yaml", strings.ToLower(month.String()), year)
		if err := os.WriteFile(fileName, buffer, os.ModePerm); err != nil {
			fmt.Printf("Error occurred: %s", err)
			os.Exit(1)
		}

		fmt.Printf("New sheet generated: %s\n", fileName)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.PersistentFlags().IntP("year", "y", 0, "Year for which entries should be generated")
}
