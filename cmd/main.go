package main

import (
	"log/slog"
	"os/exec"
	"path/filepath"

	"github.com/ggymm/webview"

	"atlas/internal/api"
	"atlas/pkg/app"
	"atlas/pkg/data"
	"atlas/pkg/utils"
)

var (
	vlc string
)

func init() {
	app.Init()
	data.Init()

	vlc = app.Player
	if len(vlc) == 0 {
		vlc = utils.LookupVLC()
	}
}

func main() {
	// 启动服务
	go func() {
		err := api.NewServer().Start()
		if err != nil {
			panic(err)
		}
	}()

	// 打开 webview 窗口
	w := webview.New(true)
	defer w.Destroy()

	w.SetTitle("ATLAS")
	w.SetSize(1440, 960, webview.HintNone)
	w.SetSize(1200, 800, webview.HintMin)

	_ = w.Bind("playVideo", func(path string) {
		path = filepath.Join(app.Root, path)
		err := exec.Command(vlc, path).Start()
		if err != nil {
			slog.Error("play video error",
				slog.Any("error", err),
				slog.String("path", path),
			)
		}
	})
	w.Navigate(app.View)
	w.Run()
}
