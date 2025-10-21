package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
)

type SeasonStatsForAllQuestions struct {
	Label        string  `json:"label"` // e.g., "2024F" or "2024H"
	Min          float64 `json:"min"`
	Max          float64 `json:"max"`
	Avg          float64 `json:"avg"`
	P50          float64 `json:"p50"`
	StdDev       float64 `json:"stdDev"`
	NumUncounted int     `json:"numUncounted"`
}

func AnalyzeRange(fromYear int, toYear int, inputFolder string) {
	seasonStats := calculateSeasonStats(fromYear, toYear, inputFolder)
	printSeasonStats(seasonStats)
}

func readQuestionStatsJSON(inputFolder string, season string, year int, day int, group string) examStats {
	filename := fmt.Sprintf("%d_%s_%d_%s.json", year, season, day, group)
	filepath := filepath.Join(inputFolder, filename)

	file, err := os.Open(filepath)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file %s: %v", filepath, err))
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(fmt.Sprintf("Failed to close file %s: %v", filepath, err))
		}
	}()

	var stats examStats
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&stats)
	if err != nil {
		panic(fmt.Sprintf("Failed to decode JSON from %s: %v", filepath, err))
	}

	return stats
}

func calculateSeasonStats(fromYear int, toYear int, inputFolder string) []SeasonStatsForAllQuestions {
	var result []SeasonStatsForAllQuestions

	for year := fromYear; year <= toYear; year++ {
		for _, season := range []string{Spring, Autumn} {
			var allCorrectPercs []float64
			var totalUncounted int

			// Collect all correct percentages and uncounted questions for this season
			for day := 1; day <= 3; day++ {
				for _, group := range []string{groupA, groupB} {
					stats := readQuestionStatsJSON(inputFolder, season, year, day, group)

					for _, question := range stats {
						if question == nil {
							// Question was removed from scoring
							totalUncounted++
						} else {
							// Get the percentage for the correct answer
							var correctPerc float64
							switch question.CorrectAnswer {
							case "A":
								correctPerc = question.A
							case "B":
								correctPerc = question.B
							case "C":
								correctPerc = question.C
							case "D":
								correctPerc = question.D
							case "E":
								correctPerc = question.E
							}
							allCorrectPercs = append(allCorrectPercs, correctPerc)
						}
					}
				}
			}

			if len(allCorrectPercs) > 0 {
				// Calculate statistics
				min := allCorrectPercs[0]
				max := allCorrectPercs[0]
				sum := 0.0

				for _, p := range allCorrectPercs {
					if p < min {
						min = p
					}
					if p > max {
						max = p
					}
					sum += p
				}

				avg := sum / float64(len(allCorrectPercs))

				// Calculate median (p50)
				sorted := make([]float64, len(allCorrectPercs))
				copy(sorted, allCorrectPercs)
				sort.Float64s(sorted)

				var p50 float64
				n := len(sorted)
				if n%2 == 0 {
					p50 = (sorted[n/2-1] + sorted[n/2]) / 2
				} else {
					p50 = sorted[n/2]
				}

				// Calculate standard deviation
				variance := 0.0
				for _, p := range allCorrectPercs {
					variance += (p - avg) * (p - avg)
				}
				variance /= float64(len(allCorrectPercs))
				stdDev := math.Sqrt(variance)

				// Create label like "2024F" or "2024H"
				seasonLabel := fmt.Sprintf("%d%s", year, season)

				result = append(result, SeasonStatsForAllQuestions{
					Label:        seasonLabel,
					Min:          min,
					Max:          max,
					Avg:          avg,
					P50:          p50,
					StdDev:       stdDev,
					NumUncounted: totalUncounted,
				})
			}
		}
	}

	return result
}

func printSeasonStats(stats []SeasonStatsForAllQuestions) {
	fmt.Println("=== Season Statistics ===")
	fmt.Println()
	fmt.Println("Season | Min   | Max   | Avg   | P50   | StdDev | Uncounted")
	fmt.Println("-------|-------|-------|-------|-------|--------|----------")

	for _, seasonStat := range stats {
		fmt.Printf("%6s | %.3f | %.3f | %.3f | %.3f | %.3f  | %d\n",
			seasonStat.Label,
			seasonStat.Min,
			seasonStat.Max,
			seasonStat.Avg,
			seasonStat.P50,
			seasonStat.StdDev,
			seasonStat.NumUncounted,
		)
	}
	fmt.Println()
}
