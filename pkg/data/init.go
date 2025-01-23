package data

import (
	_ "embed"
	"log/slog"

	"github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"atlas/pkg/app"
)

var DB *gorm.DB

//go:embed init.sql
var InitSQL string

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

	err = db.Exec(InitSQL).Error
	if err != nil {
		panic(err)
		return
	}

	v1, _, _ := sqlite3.Version()
	slog.Info("sqlite", "version", v1)

	// 导出数据库对象
	DB = db
}

func Flush() {
	if DB != nil {
		_ = DB.Exec("VACUUM").Error
	}
}
