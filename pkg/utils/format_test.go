package utils

import (
	"testing"
)

func Test_FormatSize(t *testing.T) {
	t.Logf(FormatSize(1188908547961))
}

func Test_FormatDuration(t *testing.T) {
	t.Logf(FormatDuration(3979312))
}
