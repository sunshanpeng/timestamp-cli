package converter

import (
	"strconv"
	"time"

	"github.com/sunshanpeng/timestamp-cli/pkg/parser"
)

// Convert 根据输入类型进行转换
func Convert(inputType parser.InputType, t time.Time, outputMS bool) string {
	switch inputType {
	case parser.SecondTimestamp, parser.MillisecondTimestamp:
		// 时间戳转日期字符串
		return t.Format("2006-01-02 15:04:05")
	case parser.DateString:
		// 日期字符串转时间戳
		if outputMS {
			return strconv.FormatInt(t.UnixMilli(), 10)
		}
		return strconv.FormatInt(t.Unix(), 10)
	default:
		return ""
	}
}

// GetCurrentTimestamp 获取当前时间戳
func GetCurrentTimestamp(loc *time.Location, useMS bool) string {
	now := time.Now().In(loc)
	if useMS {
		return strconv.FormatInt(now.UnixMilli(), 10)
	}
	return strconv.FormatInt(now.Unix(), 10)
}
