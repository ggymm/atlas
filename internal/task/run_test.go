package task_test

import (
	"testing"

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

func Test_Start(t *testing.T) {
	task.Start()
}
