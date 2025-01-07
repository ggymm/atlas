package data

import (
	_ "embed"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"atlas/pkg/app"
)

var DB *gorm.DB

//go:embed init.sql
var initSQL string

func Init() {
	db, err := gorm.Open(sqlite.Open(app.Datasource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: NewCustomLog(),
	})
	if err != nil {
		panic(err)
		return
	}
	DB = db

	err = db.Exec(initSQL).Error
	if err != nil {
		panic(err)
		return
	}
}

func Flush() {
	if DB != nil {
		_ = DB.Exec("VACUUM").Error
	}
}
