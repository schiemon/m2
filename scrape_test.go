package main

import (
	"net/url"
	"strings"
	"testing"
)

func TestBuildURL(t *testing.T) {
	tests := []struct {
		name     string
		season   string
		year     int
		day      int
		group    string
		wantURL  string
		validate func(t *testing.T, got string)
	}{
		{
			name:    "Spring 2024 Day 1 Group A",
			season:  "F",
			year:    2024,
			day:     1,
			group:   "A",
			wantURL: baseURL + "?ex=S-F24-Tag%201&gruppe=Gruppe%20A",
		},
		{
			name:    "Autumn 2024 Day 2 Group B",
			season:  "H",
			year:    2024,
			day:     2,
			group:   "B",
			wantURL: baseURL + "?ex=S-H24-Tag%202&gruppe=Gruppe%20B",
		},
		{
			name:    "Spring 2016 Day 1 Group A",
			season:  "F",
			year:    2016,
			day:     1,
			group:   "A",
			wantURL: baseURL + "?ex=S-F16-Tag%201&gruppe=Gruppe%20A",
		},
		{
			name:    "Autumn 2023 Day 2 Group B",
			season:  "H",
			year:    2023,
			day:     2,
			group:   "B",
			wantURL: baseURL + "?ex=S-H23-Tag%202&gruppe=Gruppe%20B",
		},
		{
			name:   "URL is properly encoded",
			season: "F",
			year:   2024,
			day:    1,
			group:  "A",
			validate: func(t *testing.T, got string) {
				// Verify the URL is valid
				_, err := url.Parse(got)
				if err != nil {
					t.Errorf("buildURL returned invalid URL: %v", err)
				}

				// Verify base URL is present
				if !strings.HasPrefix(got, baseURL) {
					t.Errorf("URL should start with BaseURL: %s, got: %s", baseURL, got)
				}

				// Verify query parameters are present
				if !strings.Contains(got, "ex=S-F24-Tag%201") {
					t.Errorf("URL should contain ex parameter with season and year")
				}
				if !strings.Contains(got, "gruppe=Gruppe%20A") {
					t.Errorf("URL should contain gruppe parameter with group")
				}
			},
		},
		{
			name:   "Year modulo 100 for 2021",
			season: "F",
			year:   2021,
			day:    1,
			group:  "A",
			validate: func(t *testing.T, got string) {
				// Year 2021 should be encoded as 21
				if !strings.Contains(got, "S-F21") {
					t.Errorf("Year 2021 should be encoded as 21 in URL, got: %s", got)
				}
			},
		},
		{
			name:   "Day 1 is URL encoded properly",
			season: "F",
			year:   2024,
			day:    1,
			group:  "A",
			validate: func(t *testing.T, got string) {
				if !strings.Contains(got, "Tag%201") {
					t.Errorf("Day 1 should be URL encoded as Tag%%201, got: %s", got)
				}
			},
		},
		{
			name:   "Day 2 is URL encoded properly",
			season: "F",
			year:   2024,
			day:    2,
			group:  "A",
			validate: func(t *testing.T, got string) {
				if !strings.Contains(got, "Tag%202") {
					t.Errorf("Day 2 should be URL encoded as Tag%%202, got: %s", got)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildURL(tt.season, tt.year, tt.day, tt.group)

			if tt.wantURL != "" {
				if got != tt.wantURL {
					t.Errorf("buildURL() = %s, want %s", got, tt.wantURL)
				}
			}

			if tt.validate != nil {
				tt.validate(t, got)
			}
		})
	}
}
