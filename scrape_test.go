package main

import (
	"testing"
)

func TestBuildURL(t *testing.T) {
	tests := []struct {
		name    string
		season  string
		year    int
		day     int
		group   string
		wantURL string
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildURL(tt.season, tt.year, tt.day, tt.group)

			if got != tt.wantURL {
				t.Errorf("buildURL() = %s, want %s", got, tt.wantURL)
			}
		})
	}
}
