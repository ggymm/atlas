module atlas

go 1.23

replace github.com/ggymm/gopkg => ../gopkg

require (
	github.com/ggymm/gopkg v1.2.6
	github.com/ggymm/webview v1.0.0
	github.com/mattn/go-sqlite3 v1.14.24
	github.com/pkg/errors v0.9.1
	gorm.io/driver/sqlite v1.5.7
	gorm.io/gorm v1.25.12
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.21.0 // indirect
)
