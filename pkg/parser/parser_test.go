package parser

import (
	"testing"
	"time"
)

func TestParseTimestamp(t *testing.T) {
	loc := time.UTC

	tests := []struct {
		name      string
		input     string
		wantType  InputType
		wantUnix  int64
		wantError bool
	}{
		{
			name:      "10位秒级时间戳",
			input:     "1768809600",
			wantType:  SecondTimestamp,
			wantUnix:  1768809600,
			wantError: false,
		},
		{
			name:      "13位毫秒级时间戳",
			input:     "1768809600000",
			wantType:  MillisecondTimestamp,
			wantUnix:  1768809600,
			wantError: false,
		},
		{
			name:      "无效长度的时间戳",
			input:     "12345",
			wantType:  Unknown,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotType, gotTime, err := ParseInput(tt.input, loc)

			if (err != nil) != tt.wantError {
				t.Errorf("ParseInput() error = %v, wantError %v", err, tt.wantError)
				return
			}

			if !tt.wantError {
				if gotType != tt.wantType {
					t.Errorf("ParseInput() gotType = %v, want %v", gotType, tt.wantType)
				}
				if gotTime.Unix() != tt.wantUnix {
					t.Errorf("ParseInput() gotTime.Unix() = %v, want %v", gotTime.Unix(), tt.wantUnix)
				}
			}
		})
	}
}

func TestParseDateString(t *testing.T) {
	loc := time.UTC

	// 预先计算期望的时间戳（使用 UTC 时区）
	expectedTime1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2026-01-19 16:00:00", loc)
	expectedTime2, _ := time.ParseInLocation("2006-01-02", "2026-01-19", loc)

	tests := []struct {
		name      string
		input     string
		wantType  InputType
		wantUnix  int64
		wantError bool
	}{
		{
			name:      "标准日期时间格式（短横线）",
			input:     "2026-01-19 16:00:00",
			wantType:  DateString,
			wantUnix:  expectedTime1.Unix(),
			wantError: false,
		},
		{
			name:      "仅日期格式（短横线）",
			input:     "2026-01-19",
			wantType:  DateString,
			wantUnix:  expectedTime2.Unix(),
			wantError: false,
		},
		{
			name:      "标准日期时间格式（斜线）",
			input:     "2026/01/19 16:00:00",
			wantType:  DateString,
			wantUnix:  expectedTime1.Unix(),
			wantError: false,
		},
		{
			name:      "仅日期格式（斜线）",
			input:     "2026/01/19",
			wantType:  DateString,
			wantUnix:  expectedTime2.Unix(),
			wantError: false,
		},
		{
			name:      "无效格式",
			input:     "invalid-date",
			wantType:  Unknown,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotType, gotTime, err := ParseInput(tt.input, loc)

			if (err != nil) != tt.wantError {
				t.Errorf("ParseInput() error = %v, wantError %v", err, tt.wantError)
				return
			}

			if !tt.wantError {
				if gotType != tt.wantType {
					t.Errorf("ParseInput() gotType = %v, want %v", gotType, tt.wantType)
				}
				if gotTime.Unix() != tt.wantUnix {
					t.Errorf("ParseInput() gotTime.Unix() = %v, want %v", gotTime.Unix(), tt.wantUnix)
				}
			}
		})
	}
}
