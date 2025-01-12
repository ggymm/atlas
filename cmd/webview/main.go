package main

import (
	"fmt"
	"github.com/ggymm/webview"
)

func main() {
	app := webview.New(true)
	defer app.Destroy()

	app.SetTitle("ATLAS")
	app.SetSize(1200, 800, webview.HintNone)
	app.SetSize(1200, 800, webview.HintMin)

	err := app.Bind("quit", func() {
		app.Terminate()
	})
	if err != nil {
		panic(err)
	}
	err = app.Bind("openPath", func() {
		// 打开本地目录
		fmt.Println("openPath")
	})
	if err != nil {
		panic(err)
	}
	app.Navigate("http://localhost:5173/")
	app.Run()
}
