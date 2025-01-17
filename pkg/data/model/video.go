package model

type Video struct {
	Id        string `gorm:"column:id;type:text;not null;comment:主键"`
	Path      string `gorm:"column:path;type:text;comment:路径"`
	Size      int64  `gorm:"column:size;type:integer;comment:大小"`
	Star      int64  `gorm:"column:star;type:integer;comment:星级"`
	Tags      string `gorm:"column:tags;type:text;comment:标签"`
	Title     string `gorm:"column:title;type:text;comment:标题"`
	Cover     []byte `gorm:"column:cover;type:blob;comment:封面"`
	Format    string `gorm:"column:format;type:text;comment:格式"`
	Duration  int64  `gorm:"column:duration;type:integer;comment:时长"`
	CreatedAt int64  `gorm:"column:created_at;type:integer;autoCreateTime:milli;comment:创建时间"`
	UpdatedAt int64  `gorm:"column:updated_at;type:integer;autoCreateTime:milli;comment:更新时间"`
}

func (*Video) TableName() string {
	return "video"
}
