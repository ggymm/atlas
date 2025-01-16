package main

import (
	"github.com/ggymm/webview"

	"atlas/internal/api"
	"atlas/pkg/app"
	"atlas/pkg/data"
	"atlas/pkg/log"
)

func init() {
	app.Init()
	log.Init()
	data.Init()
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
	w.SetSize(1200, 800, webview.HintMin)
	w.SetSize(1200, 800, webview.HintNone)

	_ = w.Bind("openPath", func() {
	})
	w.Navigate("http://localhost:5173/")
	w.Run()
}
