package app

import (
	"log/slog"

	"github.com/ggymm/gopkg/log"
)

func InitLogger() {
	log.Init(Log(), slog.LevelInfo)
}
