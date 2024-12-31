package service_test

import (
	"testing"

	"atlas/pkg/app"
	"atlas/pkg/data"
	"atlas/pkg/data/service"
	"atlas/pkg/log"
)

func init() {
	app.Init()
	log.Init()
	data.Init()
}

func Test_FetchVideos(t *testing.T) {
	r, err := service.FetchVideos(&service.VideoPageReq{
		Page: service.Page{
			Page: 1,
			Size: 20, // 每页显示数量
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("resp: %+v", r)
}
