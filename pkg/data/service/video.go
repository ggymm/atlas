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
