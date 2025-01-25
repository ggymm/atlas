package model

type Video struct {
	Id        string `gorm:"column:id;type:text;not null;comment:主键" json:"id"`
	Path      string `gorm:"column:path;type:text;comment:路径" json:"path"`
	Size      int64  `gorm:"column:size;type:integer;comment:大小" json:"size"`
	Tags      string `gorm:"column:tags;type:text;comment:标签" json:"tags"`
	Title     string `gorm:"column:title;type:text;comment:标题" json:"title"`
	Stars     int64  `gorm:"column:stars;type:integer;comment:星级" json:"stars"`
	Cover     []byte `gorm:"column:cover;type:blob;comment:封面" json:"cover"`
	Format    string `gorm:"column:format;type:text;comment:格式" json:"format"`
	Duration  int64  `gorm:"column:duration;type:integer;comment:时长" json:"duration"`
	CreatedAt int64  `gorm:"column:created_at;type:integer;autoCreateTime:milli;comment:创建时间" json:"createdAt"`
	UpdatedAt int64  `gorm:"column:updated_at;type:integer;autoCreateTime:milli;comment:更新时间" json:"updatedAt"`
}

func (*Video) TableName() string {
	return "video"
}
