package service

import (
	"atlas/internal/req"
	"atlas/internal/resp"
	"atlas/pkg/data"
	"atlas/pkg/data/model"
)

func FetchVideos(req *req.VideoPageReq) (*resp.VideoPageResp, error) {
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

	return &resp.VideoPageResp{Total: total, Records: records}, nil
}
