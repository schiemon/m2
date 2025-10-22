# M2 Exam Stats Scraper And Visualizer

This is a tool that scrapes and visualizes statistics on the German medical licensing examination M2 (Zweiter Abschnitt der Ärztlichen Prüfung). The data is fetched from [MEDI-LEARN](https://www.medi-learn.de/statistik/stex/fragenstatistiken/index.php?ex=S-F24-Tag%201&gruppe=Gruppe%20A).

## Overview

This project scrapes and computes statistics from correct answer rates from M2 exams from 2016 up until 2024.

## Exam Statistics (2016 to 2024)

```
┌───────┬───────┬───────┬───────┬───────┬────────┬───────────┐
│ EXAM  │  MIN  │  MAX  │  AVG  │  P50  │ STDDEV │ UNCOUNTED │
├───────┼───────┼───────┼───────┼───────┼────────┼───────────┤
│ 2016F │ 0.127 │ 0.994 │ 0.794 │ 0.860 │ 0.199  │ 4         │
│ 2016H │ 0.030 │ 0.996 │ 0.799 │ 0.870 │ 0.192  │ 8         │
│ 2017F │ 0.052 │ 0.996 │ 0.774 │ 0.826 │ 0.200  │ 12        │
│ 2017H │ 0.094 │ 0.995 │ 0.776 │ 0.864 │ 0.207  │ 20        │
│ 2018F │ 0.111 │ 0.996 │ 0.774 │ 0.837 │ 0.200  │ 26        │
│ 2018H │ 0.135 │ 0.994 │ 0.790 │ 0.852 │ 0.191  │ 24        │
│ 2019F │ 0.105 │ 0.993 │ 0.751 │ 0.809 │ 0.206  │ 26        │
│ 2019H │ 0.033 │ 0.993 │ 0.738 │ 0.787 │ 0.208  │ 10        │
│ 2020F │ 0.054 │ 0.995 │ 0.726 │ 0.805 │ 0.232  │ 20        │
│ 2020H │ 0.034 │ 0.992 │ 0.752 │ 0.803 │ 0.215  │ 20        │
│ 2021F │ 0.033 │ 0.994 │ 0.721 │ 0.785 │ 0.233  │ 30        │
│ 2021H │ 0.000 │ 0.994 │ 0.733 │ 0.786 │ 0.226  │ 24        │
│ 2022F │ 0.026 │ 0.994 │ 0.747 │ 0.800 │ 0.209  │ 4         │
│ 2022H │ 0.094 │ 0.997 │ 0.740 │ 0.817 │ 0.228  │ 20        │
│ 2023F │ 0.000 │ 0.995 │ 0.742 │ 0.821 │ 0.227  │ 14        │
│ 2023H │ 0.052 │ 0.994 │ 0.740 │ 0.828 │ 0.242  │ 12        │
│ 2024F │ 0.077 │ 0.994 │ 0.768 │ 0.851 │ 0.212  │ 24        │
│ 2024H │ 0.018 │ 0.996 │ 0.710 │ 0.786 │ 0.249  │ 10        │
└───────┴───────┴───────┴───────┴───────┴────────┴───────────┘
                EXAM STATISTICS (2016 to 2024)
```

```
┌─────────────────────────────────────────────────────────────────────────────┬───────┬───────┬───────┬────────┬────────────────────────────────────────────┐
│                       SUBJECT NAME  ( SUBSPECIALTY )                        │ COUNT │  AVG  │  P50  │ STDDEV │                CORRECTNESS                 │
├─────────────────────────────────────────────────────────────────────────────┼───────┼───────┼───────┼────────┼────────────────────────────────────────────┤
│ Auge                                                                        │ 2     │ 0.961 │ 0.961 │ 0.000  │ [======================================  ] │
│ Rechtsmedizin                                                               │ 22    │ 0.947 │ 0.953 │ 0.042  │ [=====================================   ] │
│ Anästhesie (Schmerzth. und Notfallmed.)                                     │ 358   │ 0.823 │ 0.888 │ 0.175  │ [================================        ] │
│ Pädiatrie (Infektionskrankheiten, Immunologie, Neurologie, Blut)            │ 222   │ 0.811 │ 0.881 │ 0.173  │ [================================        ] │
│ Urologie                                                                    │ 68    │ 0.810 │ 0.861 │ 0.173  │ [================================        ] │
│ Psychiatrie (Psychopathologie, Psychosen, Neurosen)                         │ 258   │ 0.805 │ 0.855 │ 0.191  │ [================================        ] │
│ Gynäkologie (Geburtshilfe)                                                  │ 280   │ 0.803 │ 0.860 │ 0.184  │ [================================        ] │
│ Rechtsmedizin                                                               │ 192   │ 0.792 │ 0.877 │ 0.225  │ [===============================         ] │
│ Psychiatrie (Abhängigkeit, Schizophrenie, Affektive Psychosen)              │ 334   │ 0.790 │ 0.881 │ 0.225  │ [===============================         ] │
│ Innere (Lunge)                                                              │ 508   │ 0.787 │ 0.862 │ 0.207  │ [===============================         ] │
│ Innere (Diabetes / Rheumatologie)                                           │ 368   │ 0.786 │ 0.855 │ 0.204  │ [===============================         ] │
│ Innere (Kardiologie, Kreislauf)                                             │ 348   │ 0.779 │ 0.860 │ 0.207  │ [===============================         ] │
│ Innere (Kardiologie)                                                        │ 350   │ 0.779 │ 0.837 │ 0.185  │ [===============================         ] │
│ Orthopädie (Wirbelsäule, obere Extremiät)                                   │ 230   │ 0.778 │ 0.828 │ 0.201  │ [===============================         ] │
│ Gynäkologie (Gynäkologie)                                                   │ 284   │ 0.775 │ 0.842 │ 0.209  │ [==============================          ] │
│ Neurologie (entzündliche Prozesse, Traumen, Gefäßkrankheiten, Anfälle)      │ 382   │ 0.773 │ 0.834 │ 0.204  │ [==============================          ] │
│ Innere (Gastroenterologie)                                                  │ 322   │ 0.767 │ 0.827 │ 0.212  │ [==============================          ] │
│ Chriurgie (Oberbauch)                                                       │ 312   │ 0.766 │ 0.826 │ 0.224  │ [==============================          ] │
│ Sozialmedizin                                                               │ 220   │ 0.765 │ 0.810 │ 0.190  │ [==============================          ] │
│ Chriurgie (Allgemeine Chirurgie)                                            │ 150   │ 0.765 │ 0.794 │ 0.203  │ [==============================          ] │
│ Pharmakologie (Zytostatika, Immunsuppressiva, ZNS)                          │ 188   │ 0.764 │ 0.817 │ 0.210  │ [==============================          ] │
│ Chriurgie (Traumatologie)                                                   │ 232   │ 0.762 │ 0.838 │ 0.214  │ [==============================          ] │
│ Pharmakologie (Herz-Kreislauf, Gerinnung)                                   │ 154   │ 0.762 │ 0.823 │ 0.204  │ [==============================          ] │
│ Innere (Endokrinologie)                                                     │ 272   │ 0.761 │ 0.844 │ 0.218  │ [==============================          ] │
│ Urologie                                                                    │ 220   │ 0.756 │ 0.828 │ 0.212  │ [==============================          ] │
│ Chriurgie (Unterbauch)                                                      │ 268   │ 0.756 │ 0.842 │ 0.236  │ [==============================          ] │
│ Infektionskrankheiten (Hygiene)                                             │ 426   │ 0.752 │ 0.824 │ 0.220  │ [==============================          ] │
│ Dermatologie (sonstige Hauterkrankungen)                                    │ 152   │ 0.751 │ 0.789 │ 0.176  │ [==============================          ] │
│ Pädiatrie (Stoffwechsel, Endokrinologie, Herz/Kreislauf, Atmung, Verdauung) │ 366   │ 0.748 │ 0.837 │ 0.236  │ [=============================           ] │
│ HNO                                                                         │ 222   │ 0.746 │ 0.806 │ 0.233  │ [=============================           ] │
│ Radiologie (Thorax und Gefäße)                                              │ 12    │ 0.735 │ 0.804 │ 0.208  │ [=============================           ] │
│ Arbeitsmedizin                                                              │ 128   │ 0.735 │ 0.820 │ 0.232  │ [=============================           ] │
│ Neurologie (Raumforderungen, degenerative Prozesse, Meningitis)             │ 358   │ 0.729 │ 0.795 │ 0.228  │ [=============================           ] │
│ Infektionskrankheiten (Erkrankungen)                                        │ 256   │ 0.727 │ 0.808 │ 0.234  │ [=============================           ] │
│ Innere (Hämatologie)                                                        │ 318   │ 0.727 │ 0.800 │ 0.217  │ [=============================           ] │
│ Sozialmedizin                                                               │ 148   │ 0.726 │ 0.769 │ 0.217  │ [=============================           ] │
│ Pharmakologie (Lunge, Infektionskrankheiten, Endokrine Organe)              │ 212   │ 0.726 │ 0.786 │ 0.222  │ [=============================           ] │
│ Orthopädie (Untere Extremität, Tumoren)                                     │ 210   │ 0.722 │ 0.775 │ 0.221  │ [============================            ] │
│ Innere (Niere und Wasserhaushalt)                                           │ 86    │ 0.721 │ 0.773 │ 0.168  │ [============================            ] │
│ Neurologie (Muskelkrankheiten)                                              │ 260   │ 0.718 │ 0.781 │ 0.232  │ [============================            ] │
│ Pathologie                                                                  │ 32    │ 0.716 │ 0.789 │ 0.193  │ [============================            ] │
│ Pädiatrie (Wachstum, Prä-u. Perinatalperiode)                               │ 192   │ 0.711 │ 0.766 │ 0.237  │ [============================            ] │
│ Natruheilverfahren                                                          │ 18    │ 0.709 │ 0.863 │ 0.275  │ [============================            ] │
│ Anästhesie (Anästhesie / Intensivmed.)                                      │ 146   │ 0.703 │ 0.755 │ 0.229  │ [============================            ] │
│ Dermatologie (Infektiäse, immunologische und erbliche Hauterkrankungen)     │ 282   │ 0.698 │ 0.743 │ 0.230  │ [===========================             ] │
│ Neurologie (peripheres NS, Polyneuropathien)                                │ 302   │ 0.692 │ 0.749 │ 0.220  │ [===========================             ] │
│ Augenheilkunde                                                              │ 320   │ 0.687 │ 0.723 │ 0.220  │ [===========================             ] │
│ Urologie                                                                    │ 2     │ 0.678 │ 0.678 │ 0.000  │ [===========================             ] │
│ Epidemiologie                                                               │ 22    │ 0.651 │ 0.748 │ 0.236  │ [==========================              ] │
│ Radiologie (Knochen und Abdomen)                                            │ 20    │ 0.649 │ 0.721 │ 0.240  │ [=========================               ] │
│ Humangenetik                                                                │ 154   │ 0.624 │ 0.627 │ 0.246  │ [========================                ] │
│ Natruheilverfahren                                                          │ 2     │ 0.621 │ 0.621 │ 0.000  │ [========================                ] │
│ Radiologie                                                                  │ 10    │ 0.554 │ 0.438 │ 0.222  │ [======================                  ] │
│ Pathologie                                                                  │ 32    │ 0.548 │ 0.523 │ 0.255  │ [=====================                   ] │
│ Urologie                                                                    │ 6     │ 0.515 │ 0.407 │ 0.271  │ [====================                    ] │
│ Dermatologie                                                                │ 8     │ 0.469 │ 0.567 │ 0.288  │ [==================                      ] │
└─────────────────────────────────────────────────────────────────────────────┴───────┴───────┴───────┴────────┴────────────────────────────────────────────┘
                                                            CORRECTNESS BY SUBJECT (2016 to 2024)                                                            
```


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
