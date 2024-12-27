package model

import (
	"time"

	"atlas/pkg/data"
	"atlas/pkg/uuid"
)

type Video struct {
	Id        string `gorm:"column:id;type:text;not null;comment:ID"`
	Name      string `gorm:"column:name;type:text;comment:名称"`
	Tags      string `gorm:"column:tags;type:text;comment:标签"`
	Path      string `gorm:"column:path;type:text;comment:路径"`
	Size      int64  `gorm:"column:size;type:integer;comment:大小"`
	Format    string `gorm:"column:format;type:text;comment:格式"`
	Duration  int64  `gorm:"column:duration;type:integer;comment:时长"`
	Thumbnail []byte `gorm:"column:thumbnail;type:blob;comment:封面缩略图"`
	CreatedAt int64  `gorm:"column:created_at;type:integer;comment:创建时间"`
	UpdatedAt int64  `gorm:"column:updated_at;type:integer;comment:更新时间"`
}

func (v *Video) Create() error {
	v.Id = uuid.New()
	v.CreatedAt = time.Now().UnixMilli()
	return data.DB.Create(v).Error
}

func (v *Video) Update() error {
	v.UpdatedAt = time.Now().UnixMilli()
	return data.DB.Save(v).Error
}
