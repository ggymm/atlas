package service_test

import (
	"testing"

	"atlas/internal/req"
	"atlas/internal/service"
	"atlas/pkg/app"
	"atlas/pkg/data"
	"atlas/pkg/log"
)

var (
	page = req.BasePage{
		Page: 1,
		Size: 10,
	}
)

func init() {
	app.Init()
	log.Init()
	data.Init()
}

func Test_FetchVideos(t *testing.T) {
	r, err := service.FetchVideos(&req.VideoPageReq{
		BasePage: page,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
