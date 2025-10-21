# M2 Exam Stats Scraper And Visualizer

This is a tool that scrapes and visualizes statistics on the German medical licensing examination M2 (Zweiter Abschnitt der Ärztlichen Prüfung). The data is fetched from [MEDI-LEARN](https://www.medi-learn.de/statistik/stex/fragenstatistiken/index.php?ex=S-F24-Tag%201&gruppe=Gruppe%20A).

## Overview

This project scrapes and computes statistics from correct answer rates from M2 exams from 2016 up until 2024.

## Exam Statistics (2016-2024)


Season | Min   | Max   | Avg   | P50   | StdDev | Uncounted
-------|-------|-------|-------|-------|--------|----------
 2016F | 0.127 | 0.994 | 0.794 | 0.860 | 0.199  | 2
 2016H | 0.030 | 0.996 | 0.799 | 0.870 | 0.192  | 6
 2017F | 0.052 | 0.996 | 0.774 | 0.826 | 0.200  | 10
 2017H | 0.094 | 0.995 | 0.776 | 0.864 | 0.207  | 18
 2018F | 0.111 | 0.996 | 0.774 | 0.837 | 0.200  | 24
 2018H | 0.135 | 0.994 | 0.790 | 0.852 | 0.191  | 22
 2019F | 0.105 | 0.993 | 0.751 | 0.809 | 0.206  | 24
 2019H | 0.033 | 0.993 | 0.738 | 0.787 | 0.208  | 8
 2020F | 0.054 | 0.995 | 0.726 | 0.805 | 0.232  | 18
 2020H | 0.034 | 0.992 | 0.752 | 0.803 | 0.215  | 18
 2021F | 0.033 | 0.994 | 0.721 | 0.785 | 0.233  | 28
 2021H | 0.000 | 0.994 | 0.733 | 0.786 | 0.226  | 22
 2022F | 0.026 | 0.994 | 0.747 | 0.800 | 0.209  | 2
 2022H | 0.094 | 0.997 | 0.740 | 0.817 | 0.228  | 18
 2023F | 0.000 | 0.995 | 0.742 | 0.821 | 0.227  | 12
 2023H | 0.052 | 0.994 | 0.740 | 0.828 | 0.242  | 10
 2024F | 0.077 | 0.994 | 0.768 | 0.851 | 0.212  | 22
 2024H | 0.018 | 0.996 | 0.710 | 0.786 | 0.249  | 8


## Features

- **Scrape**: Collect M2 exam data from online sources
- **Analyze**: Computes statistics by season including min, max, average, median, and standard deviation of correct answer rates

## Usage

```bash
# Show help
go run . --help

# Scrape exam data (2016-2024)
go run . scrape --from 2016 --to 2024 --out ./output

# Analyze statistics
go run . analyze --from 2016 --to 2024 --in ./output

# Analyze specific year range
go run . analyze --from 2020 --to 2024
```

## Output

The analyzer generates season-based statistics showing:
- **Season**: Year and season (F = Frühling/Spring, H = Herbst/Autumn)
- **Min/Max**: Lowest and highest correct answer percentages
- **Avg**: Average correct answer rate
- **P50**: Median correct answer rate
- **StdDev**: Standard deviation
- **Uncounted**: Number of questions removed from scoring

## Project Structure

- `main.go` - CLI entry point with command handling
- `scrape.go` - Web scraping logic
- `analyze.go` - Statistics calculation
- `output/` - Default folder for downloaded exam statistic data (JSON format)

## Development

### Setting up Git Hooks

This project uses a pre-commit hook to automatically check code formatting and potential issues before committing.

**Setup (one-time):**
```bash
chmod +x .git/hooks/pre-commit
```

**Prerequisites:**
```bash
# Install golangci-lint
brew install golangci-lint
```

The hook will automatically:
- Run `gofmt` to check/fix code formatting
- Run `go vet` to detect potential bugs
- Run `golangci-lint` for comprehensive linting (if installed)

If `gofmt` finds formatting issues, it will fix them and ask you to review and stage the changes before committing.

### CI/CD Pipeline

Linting checks are also enforced in GitHub Actions (`.github/workflows/lint.yml`):
- `gofmt` - Code formatting
- `go vet` - Bug detection
- `golangci-lint` - Comprehensive linting

The pipeline runs automatically on push and pull requests to the `main` branch.
