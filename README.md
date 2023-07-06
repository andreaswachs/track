# Track

`track` is a command line tool that helps you track your worked hours from month to month. It allows you to generate and manage time sheets, which are saved in YAML format for easy reading and editing.

`track`'s reason for being is that I am tracking my work hours at my student job and I can't figure out how to make a table in Excel that can correctly subtract timestamps from each other and produce the number of hours spent working. 

## Installation

To install `track`, you need to have a Golang compiler on your local machine. Navigate to the `src/` folder and execute

```bash
go build
```

and move the built `track` binary to a folder located in your $PATH.

## Usage

### General 

`track` provides several commands for managing your time sheets:

```bash
track [command]
```

### Commands

#### `gen`

Generates a new time sheet for a given month, optionally for a user-defined year. If no month is provided, the current month is assumed. 

```bash
track gen [month] --year [year]
```

For example, to generate a time sheet for January 2023:

```bash
track gen January --year 2023
```

#### `calc`

Calculates the hours worked given a time sheet.

```bash
track calc [sheet | sheets]
```

The basic calculation sums up the hours worked from the given sheet. A lunch break of 30 minutes is standard. A lunch break can be subtracted from a days work when using `calc` if you supply the command with the `-b` flag. Change the duration of the break with the flag `-t` 


