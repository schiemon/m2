package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	Spring = "F" // Frühjahr
	Autumn = "H" // Herbst
)

const (
	groupA = "A"
	groupB = "B"
)

const fetchDelaySeconds = 5

const baseURL = "https://www.medi-learn.de/statistik/stex/fragenstatistiken/index.php"

type questionStats struct {
	CorrectAnswer string  `json:"correctAnswer"`
	A             float64 `json:"A"`
	B             float64 `json:"B"`
	C             float64 `json:"C"`
	D             float64 `json:"D"`
	E             float64 `json:"E"`
}

type examStats map[string]*questionStats // key is question number as string

func ScrapeRange(fromYear int, toYear int, outputFolder string) {
	for year := fromYear; year <= toYear; year++ {
		if year != fromYear {
			// Delay between years to be polite to the server
			time.Sleep(fetchDelaySeconds * time.Second)
		}

		for _, season := range []string{Spring, Autumn} {
			// Each season has 3 days
			for day := 1; day <= 3; day++ {
				for _, group := range []string{groupA, groupB} {
					fmt.Printf("Scraping %d%s, day %d, group %s...\n", year, season, day, group)
					scrape(outputFolder, season, year, day, group)
				}
			}
		}
	}
}

func scrape(outputFolder string, season string, year int, day int, group string) {
	// Create output folder if it doesn't exist
	err := os.MkdirAll(outputFolder, 0755)
	if err != nil {
		panic(fmt.Sprintf("Failed to create output folder: %v", err))
	}

	url := buildURL(season, year, day, group)

	stats := fetchAndParse(url)

	// Create filename: 2024_H_1_A.json
	filename := fmt.Sprintf("%d_%s_%d_%s.json", year, season, day, group)
	filepath := filepath.Join(outputFolder, filename)

	// Write to JSON file
	writeJSON(filepath, stats)
}

func buildURL(season string, year int, day int, group string) string {
	// Format: S-H24-Tag 1
	examParam := fmt.Sprintf("S-%s%02d-Tag%%20%d", season, year%100, day)
	groupParam := fmt.Sprintf("Gruppe%%20%s", group)
	return fmt.Sprintf("%s?ex=%s&gruppe=%s", baseURL, examParam, groupParam)
}

func fetchAndParse(url string) examStats {
	// Fetch the page
	resp, err := http.Get(url)
	if err != nil {
		panic(fmt.Sprintf("Failed to fetch URL: %v", err))
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(fmt.Sprintf("Failed to close response body: %v", err))
		}
	}()

	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("HTTP error: %d %s", resp.StatusCode, resp.Status))
	}

	// Read the body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("Failed to read response body: %v", err))
	}

	// Parse HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		panic(fmt.Sprintf("Failed to parse HTML: %v", err))
	}

	return parseTable(doc)
}

func parseTable(doc *goquery.Document) examStats {
	stats := make(examStats)

	// Find the table and iterate through rows
	doc.Find("table tr").Each(func(i int, row *goquery.Selection) {
		// Skip header row
		if i == 0 {
			return
		}

		cells := row.Find("td")
		if cells.Length() < 2 {
			return
		}

		// Get question number
		questionNum := strings.TrimSpace(cells.Eq(0).Text())

		// Check if question was removed from scoring
		if cells.Length() == 2 && strings.Contains(cells.Eq(1).Text(), "aus der Wertung genommen") {
			stats[questionNum] = nil
			return
		}

		// Parse correct answer and percentages (columns 2-7)
		if cells.Length() >= 8 {
			correctAnswer := strings.TrimSpace(cells.Eq(2).Find("div").Text())

			questionStats := &questionStats{
				CorrectAnswer: correctAnswer,
				A:             parsePercentage(cells.Eq(3).Text()),
				B:             parsePercentage(cells.Eq(4).Text()),
				C:             parsePercentage(cells.Eq(5).Text()),
				D:             parsePercentage(cells.Eq(6).Text()),
				E:             parsePercentage(cells.Eq(7).Text()),
			}
			stats[questionNum] = questionStats
		}
	})

	return stats
}

func parsePercentage(text string) float64 {
	// Remove whitespace, &nbsp; and % sign
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, "\u00a0", "") // &nbsp;
	text = strings.ReplaceAll(text, " ", "")
	text = strings.ReplaceAll(text, "%", "")

	// Parse as float
	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse percentage '%s': %v", text, err))
	}

	// Convert to decimal (96.92 -> 0.9692)
	return value / 100.0
}

func writeJSON(filename string, stats examStats) {
	file, err := os.Create(filename)
	if err != nil {
		panic(fmt.Sprintf("Failed to create file: %v", err))
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(fmt.Sprintf("Failed to close file: %v", err))
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(stats)
	if err != nil {
		panic(fmt.Sprintf("Failed to encode JSON: %v", err))
	}
}
