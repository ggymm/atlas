package task

import (
	"os"
	"path/filepath"
)

func rel(root, path string) string {
	p, _ := filepath.Rel(root, path)
	return p
}

func exists(path string) bool {
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
