package model

type Index struct {
	Id    string `gorm:"column:id;type:text;not null;comment:主键"`
	Term  string `gorm:"column:term;type:text;;not null;comment:搜索词"`
	Score int64  `gorm:"column:score;type:integer;comment:分数"`
}
