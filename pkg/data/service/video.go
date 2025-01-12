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

func QueryVideos(page *Page) (int64, []*model.Video, error) {
	var (
		total   int64
		records []*model.Video
	)

	// 查询总数
	err := data.DB.Model(&model.Video{}).Count(&total).Error
	if err != nil {
		return 0, nil, err
	}

	if page == nil {
		// 查询列表
		err = data.DB.Find(&records).Error
		if err != nil {
			return 0, nil, err
		}
	} else {
		// 查询列表
		err = data.DB.Limit(page.GetSize()).Offset(page.GetOffset()).Find(&records).Error
		if err != nil {
			return 0, nil, err
		}
	}
	return total, records, nil
}

func SearchVideos(page *Page, terms []string) (int64, []*model.Video, error) {
	return 0, nil, nil
}
