package log

import (
	"log/slog"

	"github.com/ggymm/gopkg/log"

	"atlas/pkg/app"
)

func Init() {
	log.Init(app.Log(), slog.LevelInfo)
}
