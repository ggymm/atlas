package task

import (
	"os"
)

func Exists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	}

	if info.IsDir() {
		return false
	}
	return true
}
