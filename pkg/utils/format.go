package utils

import (
	"fmt"
)

func FormatSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}

func FormatDuration(s int64) string {
	mins := s / 60
	hours := mins / 60

	if hours > 0 {
		return fmt.Sprintf("%d hrs %d mins", hours, mins%60)
	} else if mins > 0 {
		return fmt.Sprintf("%d mins", mins)
	}
	return fmt.Sprintf("%d secs", s)
}
