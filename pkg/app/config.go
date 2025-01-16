package app

import (
	"path/filepath"

	"github.com/ggymm/gopkg/ini"
)

var (
	Name string
	Addr string
	Root string
)

var (
	LogPath string
)

var (
	Ffmpeg  string
	Ffprobe string
)

var (
	Datasource string

	SimpleDict      string
	SimpleExtension string
)

func Log() string {
	return filepath.Join(LogPath, Name+".log")
}

func DatabaseLog() string {
	return filepath.Join(LogPath, Name+"-db.log")
}

type Config struct {
	// App 应用配置
	App struct {
		Name string `ini:"name"`
		Addr string `ini:"addr"`
		Root string `ini:"root"`
	} `ini:"app"`

	// Log 日志配置
	Log struct {
		Path string `ini:"path"`
	} `ini:"log"`

	// Bin 可执行文件配置
	Bin struct {
		Root    string `ini:"root"`
		Ffmpeg  string `ini:"ffmpeg"`
		Ffprobe string `ini:"ffprobe"`
	} `ini:"bin"`

	// Database 数据库配置
	Database struct {
		Source string `ini:"source"`
		Simple string `ini:"simple"`
	} `ini:"database"`
}

func InitConfig() {
	var (
		cfg  = new(Config)
		root = rootPath()
		path = filepath.Join(root, "config.ini")
	)

	err := ini.MapTo(cfg, path)
	if err != nil {
		panic(err)
	}

	Name = cfg.App.Name
	Addr = cfg.App.Addr
	Root = cfg.App.Root

	LogPath = cfg.Log.Path

	base := filepath.Join(root, cfg.Bin.Root)
	Ffmpeg = filepath.Join(base, cfg.Bin.Ffmpeg)
	Ffprobe = filepath.Join(base, cfg.Bin.Ffprobe)

	Datasource = "file:" + filepath.Join(root, cfg.Database.Source)
	SimpleDict = filepath.Join(root, cfg.Database.Simple, "dict")
	SimpleExtension = filepath.Join(root, cfg.Database.Simple, "simple")
}
