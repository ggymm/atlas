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

func CheckVideo(id string) bool {
	var v model.Video
	err := data.DB.Where("id = ?", id).Limit(1).Find(&v).Error
	return err == nil && len(v.Id) > 0
}

func CreateVideo(v *model.Video) error {
	return data.DB.Create(v).Error
}

func UpdateVideo(v *model.Video) error {
	return data.DB.Save(v).Error
}

func DeleteVideo(id string) error {
	return data.DB.Delete(&model.Video{}, id).Error
}

func SelectVideos(req *VideoPageReq) (*VideoPageResp, error) {
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
