package formatter

import (
	"fmt"
	"time"
)

// FormatFull 格式化完整信息（时区、本地时间、秒级和毫秒级时间戳）
func FormatFull(loc *time.Location) string {
	now := time.Now().In(loc)

	// 获取时区信息
	zoneName, offset := now.Zone()
	offsetHours := offset / 3600
	offsetMins := (offset % 3600) / 60

	return fmt.Sprintf(`Timezone: %s (%s, UTC%+d:%02d)
Local Time: %s
Timestamp (s): %d
Timestamp (ms): %d`,
		loc.String(),
		zoneName,
		offsetHours,
		offsetMins,
		now.Format("2006-01-02 15:04:05"),
		now.Unix(),
		now.UnixMilli(),
	)
}

// FormatTimezone 格式化时区信息
func FormatTimezone(loc *time.Location) string {
	now := time.Now().In(loc)
	zoneName, offset := now.Zone()
	offsetHours := offset / 3600
	offsetMins := (offset % 3600) / 60

	return fmt.Sprintf("%s (%s, UTC%+d:%02d)",
		loc.String(),
		zoneName,
		offsetHours,
		offsetMins,
	)
}
