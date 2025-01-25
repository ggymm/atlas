package main

import (
	"fmt"
	"time"

	"atlas/internal/task"
	"atlas/pkg/app"
	"atlas/pkg/data"
)

func init() {
	app.Init()
	data.Init()
}

func main() {
	now := time.Now()
	err := task.NewScanner().Start()
	if err != nil {
		panic(err)
	}
	fmt.Printf("task.NewScanner().Start() cost: %v", time.Since(now))
}
