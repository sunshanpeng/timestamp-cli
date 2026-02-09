package parser

import (
	"testing"
	"time"
)

func TestParseRelativeTime(t *testing.T) {
	loc := time.Local
	now := time.Now().In(loc)

	tests := []struct {
		input    string
		expected time.Time
		hasError bool
	}{
		{"-5m", now.Add(-5 * time.Minute), false},
		{"+1h", now.Add(1 * time.Hour), false},
		{"30s", now.Add(30 * time.Second), false},
		{"-1d", now.Add(-24 * time.Hour), false},
		{"invalid", time.Time{}, true},
		{"123", time.Time{}, true}, // Pure number should be handled by isNumeric check in ParseInput, but here testing parseRelativeTime directly which expects units
		{"-", time.Time{}, true},
	}

	for _, test := range tests {
		inputType, result, err := parseRelativeTime(test.input, loc)
		if test.hasError {
			if err == nil {
				t.Errorf("expected error for input %s, but got nil", test.input)
			}
			continue
		}

		if err != nil {
			t.Errorf("unexpected error for input %s: %v", test.input, err)
			continue
		}

		if inputType != RelativeTime {
			t.Errorf("expected input type RelativeTime, but got %v", inputType)
		}

		// Allow slight time difference due to execution time
		diff := result.Sub(test.expected)
		if diff < -time.Second || diff > time.Second {
			t.Errorf("expected time close to %v, but got %v (diff: %v)", test.expected, result, diff)
		}
	}
}

func TestParseInput_RelativeTime(t *testing.T) {
	loc := time.Local
	
	tests := []struct {
		input string
		isRelative bool
	}{
		{"-5m", true},
		{"+1h", true},
		{"30s", true},
		{"1768809600", false}, // Timestamp
		{"2026-01-19", false}, // DateString
	}

	for _, test := range tests {
		inputType, _, err := ParseInput(test.input, loc)
		if err != nil {
			t.Errorf("unexpected error for input %s: %v", test.input, err)
			continue
		}

		if test.isRelative && inputType != RelativeTime {
			t.Errorf("expected RelativeTime for input %s, but got %v", test.input, inputType)
		}
		if !test.isRelative && inputType == RelativeTime {
			t.Errorf("expected NOT RelativeTime for input %s, but got %v", test.input, inputType)
		}
	}
}
