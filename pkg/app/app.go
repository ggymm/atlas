package app

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	root    = rootPath()
	tempDir = os.TempDir()
)

func rootPath() string {
	exe, err := os.Executable()
	if err != nil {
		return ""
	}
	base := filepath.Base(exe)
	if !strings.HasPrefix(exe, tempDir) &&
		!strings.HasPrefix(base, "___") {
		return filepath.Dir(exe)
	} else {
		var absPath string
		_, filename, _, ok := runtime.Caller(0)
		if ok {
			absPath = filepath.Dir(filename)
			// 需要根据当前文件所处目录，修改相对位置
			absPath = filepath.Join(absPath, "../../")
		}
		return absPath
	}
}

func Init() {
	InitConfig()
	InitLogger()
}
