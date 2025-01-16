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
}

func Test_QueryVideos(t *testing.T) {
	resp, err := service.QueryVideos(&service.PageReq{
		Page: 1,
		Size: 20, // 每页显示数量
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("total: %+v", resp.Total)
	for _, v := range resp.Records {
		t.Logf("record: %+v", v)
	}
}
