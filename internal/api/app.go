package api

import (
	"log/slog"
	"net/http"

	"github.com/ggymm/gopkg/cors"

	"atlas/pkg/app"
)

type Server struct {
	Addr     string
	TaskApi  *TaskApi
	VideoApi *VideoApi
}

func NewServer() *Server {
	return &Server{
		Addr:     app.Addr,
		TaskApi:  new(TaskApi),
		VideoApi: new(VideoApi),
	}
}

func (s *Server) Start() error {
	handler := http.NewServeMux()

	handler.HandleFunc("/api/task/exec", s.TaskApi.Exec)

	handler.HandleFunc("/api/video/page", s.VideoApi.GetPage)
	handler.HandleFunc("/api/video/cover/{id}", s.VideoApi.GetCover)

	// 启动服务
	slog.Info("[api] server start addr " + s.Addr)
	return http.ListenAndServe(s.Addr, cors.Handler(handler))
}
