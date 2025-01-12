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
	total, records, err := service.QueryVideos(&service.Page{
		Page: 1,
		Size: 20, // 每页显示数量
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("total: %+v", total)
	for _, v := range records {
		t.Logf("record: %+v", v)
	}
}
