package data

import (
	"database/sql"
	_ "embed"

	"github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"atlas/pkg/app"
)

var DB *gorm.DB

//go:embed init.sql
var initSQL string

func Init() {
	// db, err := gorm.Open(sqlite.Open(app.Datasource), &gorm.Config{
	// 	NamingStrategy: schema.NamingStrategy{
	// 		SingularTable: true,
	// 	},
	// })
	name := "sqlite3_simple"
	sql.Register(name, &sqlite3.SQLiteDriver{
		Extensions: []string{app.SimpleExtension},
	})
	db, err := gorm.Open(sqlite.Dialector{DriverName: name, DSN: app.Datasource}, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
		return
	}

	err = db.Exec(initSQL).Error
	if err != nil {
		panic(err)
		return
	}

	// 执行设置 jieba 词典
	configSQL := "select jieba_dict('" + app.SimpleDict + "')"
	err = db.Exec(configSQL).Error
	if err != nil {
		panic(err)
		return
	}

	// 导出数据库对象
	DB = db
}

func Flush() {
	if DB != nil {
		_ = DB.Exec("VACUUM").Error
	}
}
