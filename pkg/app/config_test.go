package app_test

import (
	"testing"

	"atlas/pkg/app"
)

func Test_InitConfig(t *testing.T) {
	app.Init()

	t.Log(app.Name)
	t.Log(app.Ffmpeg)
	t.Log(app.Ffprobe)
	t.Log(app.Datasource)
}
