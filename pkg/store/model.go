package store

type Video struct {
	Id        int64  `gorm:"column:id;type:integer;not null;comment:ID"`
	Name      string `gorm:"column:name;type:text;comment:名称"`
	Tags      string `gorm:"column:tags;type:text;comment:标签"`
	RelPath   string `gorm:"column:rel_path;type:text;comment:相对路径"`
	Thumbnail []byte `gorm:"column:thumbnail;type:blob;comment:封面缩略图"`
}
