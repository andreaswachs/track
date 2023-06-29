package sheet

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Sheet struct {
	Entries []Entry `yaml:"entires"`
}

type Entry struct {
	Date  string `yaml:"date"`
	Day   string `yaml:"day"` // This is not used, but nice when editing a sheet
	Start string `yaml:"start"`
	End   string `yaml:"end"`
}

func Load(path string) (*Sheet, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var sheet Sheet
	if err := yaml.Unmarshal(contents, &sheet); err != nil {
		return nil, err
	}

	return &sheet, nil
}

func EntriesInMonth(date time.Time) []Entry {
	// Reset time to the start of the day
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	startMonth := date.Month()

	entries := []Entry{}

	for date.Month() == startMonth {
		entry := Entry{Date: date.Format("02-01-2006"), Day: date.Weekday().String()}
		entries = append(entries, entry)
		date = date.AddDate(0, 0, 1) // Add one day
	}

	return entries
}
