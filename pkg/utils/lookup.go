package utils

import (
	"bufio"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/ggymm/gopkg/lnk"
)

func LookupVLC() string {
	// 从环境变量中查找
	vlc, err := exec.LookPath("vlc")
	if err == nil {
		return vlc
	}

	// 查找 安装路径
	for _, p := range []string{
		`C:\Program Files\VideoLAN\VLC\vlc.exe`,
		`C:\Program Files (x86)\VideoLAN\VLC\vlc.exe`,
	} {
		if _, err = os.Stat(p); err == nil {
			return p
		}
	}

	// 查找 开始菜单
	for _, p := range []string{
		filepath.Join(os.Getenv("APPDATA"), `Microsoft\Windows\Start Menu\Programs`),
		filepath.Join(os.Getenv("ProgramData"), `Microsoft\Windows\Start Menu\Programs`),
	} {
		ret := ""
		err = filepath.WalkDir(p, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() || filepath.Ext(path) != ".lnk" {
				return nil
			}

			f, err := os.Open(path)
			if err != nil {
				return nil
			}
			defer func() {
				_ = f.Close()
			}()

			link, err := lnk.ParsePath(bufio.NewReader(f))
			if filepath.Base(link) == "vlc.exe" {
				ret = link
				return fs.SkipAll
			}
			return nil
		})
		if len(ret) > 0 {
			return ret
		}
		continue
	}
	return ""
}
