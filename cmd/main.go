package main

import (
	"atlas/pkg/app"
	"atlas/pkg/data"
	"atlas/pkg/log"
	"atlas/view"
)

func init() {
	app.Init()
	log.Init()
	data.Init()
}

func main() {
	view.Show()
}
