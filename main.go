package main

import (
	"flag"
	"fmt"
	"os"
)

const defaultMinYear = 2016
const defaultMaxYear = 2024
const defaultOutputFolder = "./output"

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	command := os.Args[1]

	switch command {
	case "help", "--help", "-h":
		printHelp()
	case "scrape":
		handleScrape(os.Args[2:])
	case "analyze":
		handleAnalyze(os.Args[2:])
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printHelp()
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("Usage: main <command> [options]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  help              Show this help message")
	fmt.Println("  scrape            Scrape exam data")
	fmt.Println("  analyze           Analyze exam statistics")
	fmt.Println()
	fmt.Println("Scrape options:")
	fmt.Println("  --from YEAR       Start year (default: 2016)")
	fmt.Println("  --to YEAR         End year (default: 2024)")
	fmt.Println("  --out PATH        Output folder (default: ./output)")
	fmt.Println()
	fmt.Println("Analyze options:")
	fmt.Println("  --from YEAR       Start year (default: 2016)")
	fmt.Println("  --to YEAR         End year (default: 2024)")
	fmt.Println("  --in PATH         Input folder (default: ./output)")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  main --help")
	fmt.Println("  main scrape --from 2016 --to 2024 --out ./output")
	fmt.Println("  main analyze --from 2016 --to 2024 --in ./output")
}

func handleScrape(args []string) {
	scrapeCmd := flag.NewFlagSet("scrape", flag.ExitOnError)
	from := scrapeCmd.Int("from", defaultMinYear, "Start year")
	to := scrapeCmd.Int("to", defaultMaxYear, "End year")
	out := scrapeCmd.String("out", defaultOutputFolder, "Output folder")

	if err := scrapeCmd.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing scrape arguments: %v\n", err)
		os.Exit(1)
	}

	ScrapeRange(*from, *to, *out)
}

func handleAnalyze(args []string) {
	analyzeCmd := flag.NewFlagSet("analyze", flag.ExitOnError)
	from := analyzeCmd.Int("from", defaultMinYear, "Start year")
	to := analyzeCmd.Int("to", defaultMaxYear, "End year")
	in := analyzeCmd.String("in", defaultOutputFolder, "Input folder")

	if err := analyzeCmd.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing analyze arguments: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Analyzing from %d to %d, input from %s\n", *from, *to, *in)
	AnalyzeRange(*from, *to, *in)
}
