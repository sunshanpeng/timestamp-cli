package formatter

import (
	"strings"
	"testing"
	"time"
)

func TestFormatFull(t *testing.T) {
	loc := time.UTC

	result := FormatFull(loc)

	// 验证输出包含必要的信息
	requiredStrings := []string{
		"Timezone:",
		"Local Time:",
		"Timestamp (s):",
		"Timestamp (ms):",
		"UTC",
	}

	for _, required := range requiredStrings {
		if !strings.Contains(result, required) {
			t.Errorf("FormatFull() output missing '%s'", required)
		}
	}
}

func TestFormatTimezone(t *testing.T) {
	tests := []struct {
		name     string
		loc      *time.Location
		contains []string
	}{
		{
			name:     "UTC时区",
			loc:      time.UTC,
			contains: []string{"UTC"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatTimezone(tt.loc)

			for _, required := range tt.contains {
				if !strings.Contains(result, required) {
					t.Errorf("FormatTimezone() output missing '%s', got: %s", required, result)
				}
			}
		})
	}
}
