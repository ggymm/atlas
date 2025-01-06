package task

import (
	"os"
	"path/filepath"
)

func Rel(root, path string) string {
	p, _ := filepath.Rel(root, path)
	return p
}

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
