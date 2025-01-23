package service

import (
	"atlas/pkg/data"
	"atlas/pkg/data/model"
)

func QueryEvents(page *PageReq) (*PageResp[model.Video], error) {
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
