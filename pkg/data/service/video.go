package service

import (
	"atlas/pkg/data"
	"atlas/pkg/data/model"
)

type VideoPageReq struct {
	Page
}

type VideoPageResp struct {
	Total   int64          `json:"total"`
	Records []*model.Video `json:"records"`
}

func FetchVideos(req *VideoPageReq) (*VideoPageResp, error) {
	var (
		total   int64
		records []*model.Video
	)

	// 查询总数
	err := data.DB.Model(&model.Video{}).Count(&total).Error
	if err != nil {
		return nil, err
	}

	// 查询列表
	err = data.DB.Limit(req.GetSize()).Offset(req.GetOffset()).Find(&records).Error
	if err != nil {
		return nil, err
	}

	return &VideoPageResp{Total: total, Records: records}, nil
}
