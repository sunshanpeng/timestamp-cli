package parser

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// InputType 定义输入类型
type InputType int

const (
	Unknown InputType = iota
	SecondTimestamp
	MillisecondTimestamp
	DateString
)

// ParseInput 解析输入字符串，识别类型并转换为时间对象
func ParseInput(input string, loc *time.Location) (InputType, time.Time, error) {
	input = strings.TrimSpace(input)

	// 尝试解析为数字（时间戳）
	if isNumeric(input) {
		return parseTimestamp(input, loc)
	}

	// 尝试解析为日期字符串
	return parseDateString(input, loc)
}

// isNumeric 检查字符串是否为纯数字
func isNumeric(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

// parseTimestamp 解析时间戳（支持秒级和毫秒级）
func parseTimestamp(input string, loc *time.Location) (InputType, time.Time, error) {
	timestamp, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return Unknown, time.Time{}, err
	}

	switch len(input) {
	case 10:
		// 秒级时间戳
		return SecondTimestamp, time.Unix(timestamp, 0).In(loc), nil
	case 13:
		// 毫秒级时间戳
		seconds := timestamp / 1000
		nanos := (timestamp % 1000) * 1000000
		return MillisecondTimestamp, time.Unix(seconds, nanos).In(loc), nil
	default:
		return Unknown, time.Time{}, errors.New("invalid timestamp length (expected 10 or 13 digits)")
	}
}

// parseDateString 解析日期字符串（支持多种格式）
func parseDateString(input string, loc *time.Location) (InputType, time.Time, error) {
	formats := []string{
		"2006-01-02 15:04:05",
		"2006-01-02",
		"2006/01/02 15:04:05",
		"2006/01/02",
	}

	for _, format := range formats {
		if t, err := time.ParseInLocation(format, input, loc); err == nil {
			return DateString, t, nil
		}
	}

	return Unknown, time.Time{}, errors.New("unsupported date format (expected: 2006-01-02 15:04:05, 2006-01-02, 2006/01/02 15:04:05, or 2006/01/02)")
}
