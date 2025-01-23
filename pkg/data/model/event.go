package model

type Event struct {
	Id        string `gorm:"column:id;type:text;not null;comment:主键" json:"id"`
	Content   string `gorm:"column:content;type:text;comment:日志内容" json:"content"`
	Service   string `gorm:"column:service;type:text;comment:日志业务" json:"service"`
	CreatedAt int64  `gorm:"column:created_at;type:integer;autoCreateTime:milli;comment:日志时间" json:"CreatedAt"`
}

func (*Event) TableName() string {
	return "event"
}
