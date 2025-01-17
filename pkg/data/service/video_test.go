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

type VideoIndex struct {
	Id    string `gorm:"column:id;type:text;not null;comment:主键"`
	Tags  string `gorm:"column:tags;type:text;;not null;comment:标签"`
	Title string `gorm:"column:title;type:text;;not null;comment:标题"`
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
