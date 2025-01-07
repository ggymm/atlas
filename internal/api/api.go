package api

import (
	"github.com/ggymm/gopkg/cors"
	"net/http"

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
	handler.HandleFunc("/api/v1/video/page", s.VideoApi.GetPage)
	handler.HandleFunc("/api/v1/video/cover", s.VideoApi.GetCover)

	// 启动服务
	log.Info().Msgf("[api] start")
	log.Info().Msgf("[api] listen on %s", s.Addr)
	return http.ListenAndServe(s.Addr, cors.Handler(handler))
}
