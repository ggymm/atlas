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

func Test_CheckVideo(t *testing.T) {
	ok := service.CheckVideo("hash")
	t.Logf("ok: %v", ok)
}

func Test_SelectVideos(t *testing.T) {
	r, err := service.SelectVideos(&service.VideoPageReq{
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
