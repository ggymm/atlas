package model

type Video struct {
	Id        string `gorm:"column:id;type:text;not null;comment:ID"`
	Name      string `gorm:"column:name;type:text;comment:名称"`
	Tags      string `gorm:"column:tags;type:text;comment:标签"`
	Path      string `gorm:"column:path;type:text;comment:路径"`
	Size      int64  `gorm:"column:size;type:integer;comment:大小"`
	Cover     []byte `gorm:"column:cover;type:blob;comment:封面"`
	Format    string `gorm:"column:format;type:text;comment:格式"`
	Duration  int64  `gorm:"column:duration;type:integer;comment:时长"`
	CreatedAt int64  `gorm:"column:created_at;type:integer;autoCreateTime:milli;comment:创建时间"`
	UpdatedAt int64  `gorm:"column:updated_at;type:integer;autoCreateTime:milli;comment:更新时间"`
}
