package main

import (
	"fmt"
	"time"

	"atlas/internal/task"
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
	now := time.Now()
	err := task.NewScanner().Test()
	if err != nil {
		panic(err)
	}
	fmt.Printf("task.NewScanner().Start() cost: %v", time.Since(now))
}
