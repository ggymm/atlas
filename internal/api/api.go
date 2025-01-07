package api

import (
	"net/http"

	"github.com/ggymm/gopkg/cors"

	"atlas/pkg/app"
	"atlas/pkg/log"
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
	handler.HandleFunc("/api/video/page", s.VideoApi.GetPage)
	handler.HandleFunc("/api/video/cover/{id}", s.VideoApi.GetCover)

	// 启动服务
	log.Info().Msgf("[api] start")
	log.Info().Msgf("[api] listen on %s", s.Addr)
	return http.ListenAndServe(s.Addr, cors.Handler(handler))
}
