package main

import (
	"atlas/internal/api"
	"atlas/internal/view"
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

	err := view.NewWebview().Start()
	if err != nil {
		panic(err)
	}
}
