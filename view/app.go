package view

import (
	"os"

	"gioui.org/app"
)

func Show() {
	ui := NewUI()
	go func() {
		w := new(app.Window)
		w.Option(
			app.Title("ATLAS"),
		)
		if err := ui.Run(w); err != nil {
			os.Exit(1)
		}
		os.Exit(0)
	}()
	app.Main()
}
