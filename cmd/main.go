package main

import (
	"os/exec"

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

	// 启动浏览器
	err := exec.Command(app.Webview, app.View).Run()
	if err != nil {
		panic(err)
	}
}
