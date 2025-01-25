package main

import (
	"atlas/internal/api"
	"atlas/pkg/app"
	"atlas/pkg/data"
)

func init() {
	app.Init()
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

	select {}
}
