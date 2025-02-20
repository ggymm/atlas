package api

import (
	"atlas/view"
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

	handler.Handle("/", view.NewFileServer())

	handler.HandleFunc("/api/task/post/exec", s.TaskApi.PostExec)
	handler.HandleFunc("/api/task/post/clean", s.TaskApi.PostClean)
	handler.HandleFunc("/api/task/query/events", s.TaskApi.QueryEvents)

	handler.HandleFunc("/api/video/play/{id}", s.VideoApi.Play)
	handler.HandleFunc("/api/video/cover/{id}", s.VideoApi.Cover)
	handler.HandleFunc("/api/video/query/info", s.VideoApi.QueryInfo)
	handler.HandleFunc("/api/video/query/page", s.VideoApi.QueryPage)
	handler.HandleFunc("/api/video/query/paths", s.VideoApi.QueryPaths)
	handler.HandleFunc("/api/video/update/stars", s.VideoApi.UpdateStars)

	// 启动服务
	slog.Info("api server starting", slog.String("addr", s.Addr))
	return http.ListenAndServe(s.Addr, cors.Handler(handler))
}
