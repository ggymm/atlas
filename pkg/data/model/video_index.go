package model

type VideoIndex struct {
	Id    string `gorm:"column:id;type:text;not null;comment:主键"`
	Tags  string `gorm:"column:tags;type:text;;not null;comment:标签"`
	Title string `gorm:"column:title;type:text;;not null;comment:标题"`
}
