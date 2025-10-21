# M2 Exam Statistics

Analysis tool for the German medical licensing examination M2 (Zweiter Abschnitt der Ärztlichen Prüfung).

## Overview

This project scrapes and analyzes statistics from M2 exam questions, tracking correct answer rates across different exam seasons and years (2016-2024).

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
