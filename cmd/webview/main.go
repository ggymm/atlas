package main

import (
	"github.com/ggymm/webview"
)

func main() {
	app := webview.New(true)
	defer app.Destroy()

	app.SetTitle("ATLAS")
	app.SetSize(800, 600, webview.HintNone)
	err := app.Bind("quit", func() {
		app.Terminate()
	})
	app.Navigate("http://localhost:5173/")
	if err != nil {
		panic(err)
	}
	app.Run()
}
