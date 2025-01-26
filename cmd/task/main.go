package main

import (
	"atlas/internal/task"
	"atlas/pkg/app"
	"atlas/pkg/data"
)

func init() {
	app.Init()
	data.Init()
}

func main() {
	err := task.NewScanner().Start()
	if err != nil {
		panic(err)
	}
}
