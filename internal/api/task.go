package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ggymm/gopkg/uuid"

	"atlas/internal/task"
	"atlas/pkg/data"
	"atlas/pkg/data/model"
)

type TaskApi struct {
	Api
}

func (h *TaskApi) PostExec(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}

	// 记录事件
	now := time.Now()
	data.DB.Create(&model.Event{
		Id:      uuid.NewUUID(),
		Content: "task exec start running",
		Service: "task",
	})

	go func() {
		// 执行任务
		err := task.NewScanner().Start()
		if err != nil {
			h.error(w, http.StatusInternalServerError, err.Error())
			return
		}

		// 记录事件
		data.DB.Create(&model.Event{
			Id:      uuid.NewUUID(),
			Content: fmt.Sprintf("task exec finish running, cost: %v", time.Since(now)),
			Service: "task",
		})
	}()
	h.ok(w, true)
}

func (h *TaskApi) PostClean(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}

	// 记录事件
	now := time.Now()
	data.DB.Create(&model.Event{
		Id:      uuid.NewUUID(),
		Content: "task clean start running",
		Service: "task",
	})

	go func() {
		data.DB.Exec("DROP TABLE IF EXISTS video")
		data.DB.Exec("VACUUM")
		data.DB.Exec(data.InitSQL)

		// 记录事件
		data.DB.Create(&model.Event{
			Id:      uuid.NewUUID(),
			Content: fmt.Sprintf("task clean finish running, cost: %v", time.Since(now)),
			Service: "task",
		})
	}()
	h.ok(w, true)
}

func (h *TaskApi) QueryEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}
	query := &model.Event{
		Service: "task",
	}
	order := "created_at desc"
	events := make([]model.Event, 0)

	// 查询事件
	err := data.DB.Where(query).Order(order).Limit(20).Find(&events).Error
	if err != nil {
		internalServerError(w)
		return
	}
	h.ok(w, events)
}
