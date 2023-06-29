# Track

`track` is a command line tool that helps you track your worked hours from month to month. It allows you to generate and manage time sheets, which are saved in YAML format for easy reading and editing.

## Installation

(Provide installation instructions here)

## Usage

### General 

`track` provides several commands for managing your time sheets:

```bash
track [command]
```

### Commands

#### `calc`

Calculates the hours worked given a time sheet.

```bash
track calc
```

#### `gen`

Generates a new time sheet for a given month, optionally for a user-defined year. If no month is provided, the current month is assumed. 

```bash
track gen [month] --year [year]
```

For example, to generate a time sheet for January 2023:

```bash
track gen January --year 2023
```
