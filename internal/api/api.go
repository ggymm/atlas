package api

import (
	"atlas/pkg/app"
	"net/http"
)

type Server struct {
	Addr     string
	VideoApi *VideoApi
}

func NewServer() *Server {
	return &Server{
		Addr:     app.Addr,
		VideoApi: new(VideoApi),
	}
}

func (s *Server) Start() error {
	handler := http.NewServeMux()
	handler.HandleFunc("/videos", s.VideoApi.SelectVideos)

	// 启动服务
	return http.ListenAndServe(s.Addr, handler)
}
