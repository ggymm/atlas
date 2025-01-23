package api

import (
	"fmt"
	"net/http"
	"time"

	"atlas/internal/task"
	"atlas/pkg/data"
)

type TaskApi struct {
	Api
}

func (h *TaskApi) Exec(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}

	now := time.Now()
	err := task.NewScanner().Start()
	if err != nil {
		panic(err)
	}
	fmt.Printf("task.NewScanner().Start() cost: %v", time.Since(now))

	h.ok(w, "ok")
}

func (h *TaskApi) Clean(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}

	// 开始时间
	start := time.Now()

	// 删除表
	data.DB.Exec("DROP TABLE IF EXISTS video")

	// 释放空间
	data.DB.Exec("VACUUM")

	// 重新初始化
	data.DB.Exec(data.InitSQL)

	// 返回结果
	h.ok(w, time.Since(start).String())
}
