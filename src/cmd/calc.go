/*
Copyright © 2023 Andreas Wachs <andreas@wachs.dk>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/andreaswachs/track/pkg/hours"
	"github.com/andreaswachs/track/pkg/sheet"
	"github.com/spf13/cobra"
)

// calcCmd represents the calc command
var calcCmd = &cobra.Command{
	Use:   "calc",
	Short: "calculates the hours worked given a time sheet",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No files provided to calculate hours for")
			os.Exit(2)
		}

		for _, file := range args {
			if err := CheckFile(file); err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
		}

		var sheets []*sheet.Sheet
		for _, file := range args {
			sheet, err := sheet.Load(file)
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}

			sheets = append(sheets, sheet)
		}

		var lunchTimeDeduct time.Duration
		shouldSubtractLunch, err := cmd.PersistentFlags().GetBool("substractbreak")
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		if shouldSubtractLunch {
			lunchTimeDeduct, err = cmd.PersistentFlags().GetDuration("break")
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
		}

		var workedTime time.Duration
		for _, singleSheet := range sheets {
			for _, entry := range singleSheet.Entries {
				if entry.Start == "" || entry.End == "" {
					continue
				}
				duration, err := hours.TimeDifference(entry.Start, entry.End)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}

				workedTime += duration - lunchTimeDeduct
			}
		}

		fmt.Printf("Worked hours: %.02f\n", workedTime.Hours())
	},
}

func init() {
	rootCmd.AddCommand(calcCmd)
	calcCmd.PersistentFlags().BoolP("break", "b", false, "Subtract break from each day worked for unpaid lunch")
	calcCmd.PersistentFlags().DurationP("breaktime", "t", 30*time.Minute, "Time to deduct for break")
}

func CheckFile(path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file does not exist: %s", path)
		}
		return fmt.Errorf("error retrieving file information: %v", err)
	}

	if fileInfo.Size() == 0 {
		return fmt.Errorf("file is empty: %s", path)
	}

	return nil
}
