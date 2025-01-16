package service

import (
	"atlas/pkg/data"
	"atlas/pkg/data/model"
	"gorm.io/gorm"
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

	// 查询总数
	err := data.DB.Model(&model.Video{}).Count(&total).Error
	if err != nil {
		return nil, err
	}

	if page == nil {
		// 查询列表
		err = data.DB.Find(&records).Error
		if err != nil {
			return nil, err
		}
	} else {
		// 查询列表
		err = data.DB.Limit(page.GetSize()).Offset(page.GetOffset()).Find(&records).Error
		if err != nil {
			return nil, err
		}
	}
	return &PageResp[model.Video]{Total: total, Records: records}, nil
}

func CreateVideo(v *model.Video) error {
	return data.DB.Transaction(func(tx *gorm.DB) error {
		// 创建 video 数据
		err := tx.Create(v).Error
		if err != nil {
			return err
		}

		// 创建 video_index 数据
		vi := new(model.VideoIndex)
		vi.Id = v.Id
		vi.Tags = v.Tags
		vi.Title = v.Title
		return tx.Create(vi).Error
	})
}

func UpdateVideo(v *model.Video) error {
	return data.DB.Transaction(func(tx *gorm.DB) error {
		// 更新 video 数据
		err := tx.Save(v).Error
		if err != nil {
			return err
		}

		// 更新 video_index 数据
		vi := new(model.VideoIndex)
		vi.Id = v.Id
		vi.Tags = v.Tags
		vi.Title = v.Title
		return tx.Model(vi).Where("id=?", v.Id).Updates(vi).Error
	})
}

func DeleteVideo(v *model.Video) error {
	return data.DB.Transaction(func(tx *gorm.DB) error {
		// 删除 video 数据
		err := tx.Delete(v).Error
		if err != nil {
			return err
		}

		// 删除 video_index 数据
		return tx.Where("id = ?", v.Id).Delete(new(model.VideoIndex)).Error
	})
}

func SearchVideos(page *PageReq, terms []string) (*PageResp[model.Video], error) {
	return nil, nil
}
