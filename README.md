# Timestamp CLI

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Release](https://img.shields.io/github/v/release/sunshanpeng/timestamp-cli)](https://github.com/sunshanpeng/timestamp-cli/releases)

A simple and fast command-line tool for handling timestamp conversions.

[中文文档](README_CN.md)

## Features

- ✅ Get current timestamp (seconds/milliseconds)
- ✅ Convert timestamp to human-readable format
- ✅ Convert date string to timestamp
- ✅ Smart recognition (10-digit = seconds, 13-digit = milliseconds)
- ✅ UTC timezone support
- ✅ Multiple date format support

## Installation

### Using Go Install (Recommended)

```bash
go install github.com/sunshanpeng/timestamp-cli/cmd/timestamp@latest
```

### Download Binary

Download pre-built binaries from [GitHub Releases](https://github.com/sunshanpeng/timestamp-cli/releases).

### Build from Source

```bash
git clone https://github.com/sunshanpeng/timestamp-cli.git
cd timestamp-cli
go build -o timestamp
```

## Usage

### Get Current Time Information

```bash
# Show full information (timezone, local time, timestamps)
timestamp

# Output:
# Timezone: Asia/Shanghai (CST, UTC+8:00)
# Local Time: 2026-02-09 16:48:30
# Timestamp (s): 1738227510
# Timestamp (ms): 1738227510000
```

### Get Timestamp Only

```bash
# Get seconds timestamp
timestamp -s
# Output: 1738227510

# Get milliseconds timestamp
timestamp --ms
# Output: 1738227510000
```

### Convert Timestamp to Date

```bash
# 10-digit timestamp (seconds)
timestamp 1768809600
# Output: 2026-01-19 16:00:00

# 13-digit timestamp (milliseconds, auto-detected)
timestamp 1768809600000
# Output: 2026-01-19 16:00:00
```

### Convert Date to Timestamp

```bash
# Date string to seconds timestamp
timestamp "2026-01-19 16:00:00"
# Output: 1768809600

# Date string to milliseconds timestamp
timestamp "2026-01-19 16:00:00" --ms
# Output: 1768809600000

# Support multiple date formats
timestamp "2026-01-19"           # YYYY-MM-DD
timestamp "2026/01/19 16:00:00"  # YYYY/MM/DD HH:MM:SS
timestamp "2026/01/19"           # YYYY/MM/DD
```

### UTC Timezone Support

```bash
# Show UTC timezone information
timestamp --utc
# Output: Timezone: UTC (UTC, UTC+0:00)...

# Get UTC timestamp
timestamp -s --utc

# Parse timestamp in UTC
timestamp 1768809600 --utc

# Parse date string in UTC
timestamp "2026-01-19 16:00:00" --utc
```

### Show Timezone Information

```bash
timestamp --tz
# Output: Asia/Shanghai (CST, UTC+8:00)
```

### Version Information

```bash
timestamp -v
# Output: timestamp version 1.0.0
```

## Command-Line Options

| Option | Short | Description |
|--------|-------|-------------|
| `--help` | `-h` | Show help information |
| `--version` | `-v` | Show version information |
| `--second` | `-s` | Output seconds timestamp only |
| `--ms` | | Output/use milliseconds timestamp |
| `--utc` | | Use UTC timezone (default: local timezone) |
| `--tz` | | Show timezone information |

## Supported Date Formats

- `2006-01-02 15:04:05` (Standard datetime)
- `2006-01-02` (Date only)
- `2006/01/02 15:04:05` (Datetime with slashes)
- `2006/01/02` (Date with slashes)

## Use Cases

- 🔧 Development & Debugging: Quick timestamp lookups
- 📝 Script Writing: Get timestamps in shell scripts
- 📊 Log Analysis: Convert timestamps in log files
- ⏰ Time Calculation: Convert between different time formats

## Examples

```bash
# Get current timestamp for a script
NOW=$(timestamp -s)
echo "Current timestamp: $NOW"

# Convert log timestamp to readable format
timestamp 1768809600
# Output: 2026-01-19 16:00:00

# Get UTC timestamp for API calls
timestamp -s --utc
```

## Development

### Run Tests

```bash
go test -v ./...
```

### Run Tests with Coverage

```bash
go test -cover ./...
```

### Build

```bash
go build -o timestamp
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

**sunshanpeng** - [sunshanpeng@outlook.com](mailto:sunshanpeng@outlook.com)

## Repository

https://github.com/sunshanpeng/timestamp-cli
