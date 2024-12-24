package app

import (
	"testing"
)

func Test_Init(t *testing.T) {
	Init()

	t.Log(Name)
	t.Log(Ffmpeg)
	t.Log(Ffprobe)
}
