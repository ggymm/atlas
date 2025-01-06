package task_test

import (
	"testing"
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

func TestScanner_Start(t *testing.T) {
	now := time.Now()
	err := task.NewScanner().Test()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("task.NewScanner().Start() cost: %v", time.Since(now))
}
