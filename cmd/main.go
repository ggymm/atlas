package main

import (
	"net/http"
	_ "net/http/pprof"

	"atlas/internal/view"
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
	go func() {
		err := http.ListenAndServe("localhost:6060", nil)
		if err != nil {
			panic(err)
		}
	}()
	view.Show()
}
