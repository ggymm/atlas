package app_test

import (
	"log/slog"
	"testing"

	"atlas/pkg/app"
)

func Test_InitLogger(t *testing.T) {
	app.Init()

	slog.Info("message", slog.String("key", "value"))
}
