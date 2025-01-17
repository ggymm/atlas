package service

import (
	"atlas/pkg/data"
	"atlas/pkg/data/model"
)

func GetVideo(v *model.Video) error {
	return data.DB.First(v).Error
}

func CheckVideo(v *model.Video) bool {
	err := data.DB.Where("path = ?", v.Path).Limit(1).Find(&v).Error
	return err == nil && len(v.Id) > 0
}

func QueryVideos(page *PageReq) (*PageResp[model.Video], error) {
	var (
		total   int64
		records []*model.Video
	)
	if page == nil {
		page = &PageReq{Page: 1, Size: 20}
	}

	// 查询总数
	err := data.DB.Model(&model.Video{}).Count(&total).Error
	if err != nil {
		return nil, err
	}

	// 查询列表
	err = data.DB.Limit(page.GetSize()).Offset(page.GetOffset()).Find(&records).Error
	if err != nil {
		return nil, err
	}
	return &PageResp[model.Video]{Total: total, Records: records}, nil
}
