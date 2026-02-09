package converter

import (
	"strconv"
	"testing"
	"time"

	"github.com/sunshanpeng/timestamp-cli/pkg/parser"
)

func TestConvert(t *testing.T) {
	loc := time.UTC
	// 使用 UTC 时区创建时间：2026-01-19 16:00:00 UTC
	testTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2026-01-19 16:00:00", loc)

	tests := []struct {
		name      string
		inputType parser.InputType
		time      time.Time
		outputMS  bool
		want      string
	}{
		{
			name:      "秒级时间戳转日期",
			inputType: parser.SecondTimestamp,
			time:      testTime,
			outputMS:  false,
			want:      "2026-01-19 16:00:00",
		},
		{
			name:      "毫秒级时间戳转日期",
			inputType: parser.MillisecondTimestamp,
			time:      testTime,
			outputMS:  false,
			want:      "2026-01-19 16:00:00",
		},
		{
			name:      "日期转秒级时间戳",
			inputType: parser.DateString,
			time:      testTime,
			outputMS:  false,
			want:      strconv.FormatInt(testTime.Unix(), 10),
		},
		{
			name:      "日期转毫秒级时间戳",
			inputType: parser.DateString,
			time:      testTime,
			outputMS:  true,
			want:      strconv.FormatInt(testTime.UnixMilli(), 10),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Convert(tt.inputType, tt.time, tt.outputMS)
			if got != tt.want {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCurrentTimestamp(t *testing.T) {
	loc := time.UTC

	tests := []struct {
		name  string
		useMS bool
	}{
		{
			name:  "获取秒级时间戳",
			useMS: false,
		},
		{
			name:  "获取毫秒级时间戳",
			useMS: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetCurrentTimestamp(loc, tt.useMS)

			// 验证返回的是数字字符串
			if len(got) == 0 {
				t.Errorf("GetCurrentTimestamp() returned empty string")
			}

			// 验证长度
			if tt.useMS && len(got) != 13 {
				t.Errorf("GetCurrentTimestamp() length = %d, want 13 for milliseconds", len(got))
			}
			if !tt.useMS && len(got) != 10 {
				t.Errorf("GetCurrentTimestamp() length = %d, want 10 for seconds", len(got))
			}
		})
	}
}
