# M2 Exam Statistics

Analysis tool for the German medical licensing examination M2 (Zweiter Abschnitt der Ärztlichen Prüfung).

## Overview

This project scrapes and analyzes statistics from M2 exam questions, tracking correct answer rates across different exam seasons and years (2016-2024).

## Season Statistics (2016-2024)

```
Season | Min   | Max   | Avg   | P50   | StdDev | Uncounted
-------|-------|-------|-------|-------|--------|----------
 2016F | 0.127 | 0.993 | 0.816 | 0.888 | 0.190  | 2
 2016H | 0.030 | 0.996 | 0.810 | 0.888 | 0.194  | 6
 2017F | 0.052 | 0.994 | 0.769 | 0.814 | 0.200  | 4
 2017H | 0.094 | 0.995 | 0.795 | 0.884 | 0.206  | 6
 2018F | 0.174 | 0.996 | 0.784 | 0.845 | 0.196  | 18
 2018H | 0.135 | 0.994 | 0.783 | 0.843 | 0.191  | 16
 2019F | 0.105 | 0.993 | 0.751 | 0.809 | 0.212  | 14
 2019H | 0.033 | 0.992 | 0.725 | 0.771 | 0.214  | 4
 2020F | 0.054 | 0.993 | 0.743 | 0.808 | 0.220  | 16
 2020H | 0.034 | 0.991 | 0.737 | 0.788 | 0.226  | 14
 2021F | 0.033 | 0.994 | 0.715 | 0.776 | 0.226  | 20
 2021H | 0.000 | 0.989 | 0.716 | 0.783 | 0.237  | 14
 2022F | 0.026 | 0.994 | 0.741 | 0.800 | 0.216  | 2
 2022H | 0.094 | 0.997 | 0.739 | 0.808 | 0.227  | 16
 2023F | 0.000 | 0.995 | 0.745 | 0.821 | 0.225  | 8
 2023H | 0.052 | 0.990 | 0.734 | 0.827 | 0.253  | 8
 2024F | 0.111 | 0.993 | 0.769 | 0.851 | 0.210  | 16
 2024H | 0.018 | 0.996 | 0.712 | 0.805 | 0.252  | 6
```

## Features

- **Scrape**: Collect M2 exam data from online sources
- **Analyze**: Generate statistics by season including min, max, average, median, and standard deviation of correct answer rates

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
