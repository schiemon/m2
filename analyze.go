package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
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
	printSeasonStats(seasonStats, fromYear, toYear)

	fmt.Println()
	subjectStats := calculateSubjectStats(fromYear, toYear, inputFolder)
	printSubjectStats(subjectStats, fromYear, toYear)
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

func printSeasonStats(stats []SeasonStatsForAllQuestions, fromYear int, toYear int) {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header("Exam", "Min", "Max", "Avg", "P50", "STDDEV", "Uncounted")

	for _, seasonStat := range stats {
		_ = table.Append(
			seasonStat.Label,
			fmt.Sprintf("%.3f", seasonStat.Min),
			fmt.Sprintf("%.3f", seasonStat.Max),
			fmt.Sprintf("%.3f", seasonStat.Avg),
			fmt.Sprintf("%.3f", seasonStat.P50),
			fmt.Sprintf("%.3f", seasonStat.StdDev),
			fmt.Sprintf("%d", seasonStat.NumUncounted),
		)
	}

	table.Caption(tw.Caption{
		Text: fmt.Sprintf("EXAM STATISTICS (%d to %d)", fromYear, toYear),
	})
	if err := table.Render(); err != nil {
		fmt.Fprintf(os.Stderr, "Error rendering table: %v\n", err)
	}
}

type SubjectStats struct {
	Tag               string
	SubjectName       string
	SubSpecialtyName  *string
	Count             int
	CountMin          int
	CountMax          int
	CountAvg          float64
	CountMedian       float64
	CountStdDev       float64
	CorrectnessAvg    float64
	CorrectnessP50    float64
	CorrectnessStdDev float64
}

func tagToSubjectName(tag string) string {
	// Map known tags to subject names
	subjectMap := map[string]string{
		"Aug": "Auge",
		"Rec": "Rechtsmedizin",
		"Pat": "Pathologie",
		"Epi": "Epidemiologie",
	}

	if name, exists := subjectMap[tag]; exists {
		return name
	}
	return tag // Fallback to tag if unknown
}

func calculateSubjectStats(fromYear int, toYear int, inputFolder string) []SubjectStats {
	// Map to track per-subject data
	type SubjectData struct {
		correctnesses []float64
		name          string
		subspecialty  *string
	}
	subjectDataMap := make(map[string]*SubjectData) // tag -> subject data

	for year := fromYear; year <= toYear; year++ {
		for _, season := range []string{Spring, Autumn} {
			for day := 1; day <= 3; day++ {
				for _, group := range []string{groupA, groupB} {
					stats := readQuestionStatsJSON(inputFolder, season, year, day, group)

					for _, question := range stats {
						if question != nil && question.Subject != nil {
							tag := question.Subject.Tag
							subjectName := question.Subject.Name

							if tag == "" {
								continue
							}

							if subjectName == "" {
								subjectName = tagToSubjectName(tag)
							}

							// Initialize if not seen before
							if _, exists := subjectDataMap[tag]; !exists {
								subjectDataMap[tag] = &SubjectData{
									correctnesses: []float64{},
									name:          subjectName,
									subspecialty:  question.Subject.SubSubjectName,
								}
							}

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

							subjectDataMap[tag].correctnesses = append(subjectDataMap[tag].correctnesses, correctPerc)
						}
					}
				}
			}
		}
	}

	// Convert to slice with statistics
	var result []SubjectStats
	for tag, data := range subjectDataMap {
		result = append(result, SubjectStats{
			Tag:              tag,
			SubjectName:      data.name,
			SubSpecialtyName: data.subspecialty,
			Count:            len(data.correctnesses),
		})
	}

	// Calculate statistics for each subject
	for i, s := range result {
		difficulties := subjectDataMap[s.Tag].correctnesses
		if len(difficulties) > 0 {
			// Count statistics
			countMin := s.Count
			countMax := s.Count
			countAvg := float64(s.Count)
			countMedian := float64(s.Count)

			// Difficulty statistics
			minDiff := difficulties[0]
			maxDiff := difficulties[0]
			sumDiff := 0.0

			for _, d := range difficulties {
				if d < minDiff {
					minDiff = d
				}
				if d > maxDiff {
					maxDiff = d
				}
				sumDiff += d
			}

			avgDiff := sumDiff / float64(len(difficulties))

			// Calculate median difficulty
			sortedDiff := make([]float64, len(difficulties))
			copy(sortedDiff, difficulties)
			sort.Float64s(sortedDiff)

			var medianDiff float64
			n := len(sortedDiff)
			if n%2 == 0 {
				medianDiff = (sortedDiff[n/2-1] + sortedDiff[n/2]) / 2
			} else {
				medianDiff = sortedDiff[n/2]
			}

			// Calculate standard deviation of difficulty
			varianceDiff := 0.0
			for _, d := range difficulties {
				varianceDiff += (d - avgDiff) * (d - avgDiff)
			}
			varianceDiff /= float64(len(difficulties))
			stdDevDiff := math.Sqrt(varianceDiff)

			result[i].CountMin = countMin
			result[i].CountMax = countMax
			result[i].CountAvg = countAvg
			result[i].CountMedian = countMedian
			result[i].CountStdDev = 0
			result[i].CorrectnessAvg = avgDiff
			result[i].CorrectnessP50 = medianDiff
			result[i].CorrectnessStdDev = stdDevDiff
		}
	}

	// Sort by correctness in descending order (highest first = most correct)
	sort.Slice(result, func(i, j int) bool {
		return result[i].CorrectnessAvg > result[j].CorrectnessAvg
	})

	return result
}

func printSubjectStats(stats []SubjectStats, fromYear int, toYear int) {
	if len(stats) == 0 {
		return
	}

	// Create table
	table := tablewriter.NewWriter(os.Stdout)
	table.Header("Subject Name (Subspecialty)", "Count", "Avg", "P50", "STDDEV", "Correctness")
	// Set caption as title
	table.Caption(tw.Caption{
		Text: fmt.Sprintf("CORRECTNESS BY SUBJECT (%d to %d)", fromYear, toYear),
	})

	for _, s := range stats {
		// Build the display name
		displayName := s.SubjectName
		if s.SubSpecialtyName != nil {
			displayName = fmt.Sprintf("%s (%s)", s.SubjectName, *s.SubSpecialtyName)
		}

		// Calculate histogram bar position based on average correctness
		// Lower (0.0) = harder questions, Higher (1.0) = easier questions
		histogramWidth := 40
		histWidth := int(s.CorrectnessAvg * float64(histogramWidth))
		if histWidth < 0 {
			histWidth = 0
		}
		if histWidth > histogramWidth {
			histWidth = histogramWidth
		}

		histogram := "[" + strings.Repeat("=", histWidth) + strings.Repeat(" ", histogramWidth-histWidth) + "]"

		_ = table.Append(
			displayName,
			fmt.Sprintf("%d", s.Count),
			fmt.Sprintf("%.3f", s.CorrectnessAvg),
			fmt.Sprintf("%.3f", s.CorrectnessP50),
			fmt.Sprintf("%.3f", s.CorrectnessStdDev),
			histogram,
		)
	}

	if err := table.Render(); err != nil {
		fmt.Fprintf(os.Stderr, "Error rendering table: %v\n", err)
	}
}
