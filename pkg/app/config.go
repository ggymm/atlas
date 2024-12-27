package app

import (
	"path/filepath"

	"github.com/ggymm/gopkg/ini"
)

var (
	Name string
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
)

func Log() string {
	return filepath.Join(LogPath, Name+".log")
}

type Config struct {
	// App 应用配置
	App struct {
		Name string `ini:"name"`
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

	// Server 服务器配置
	Server struct {
		Addr int `ini:"addr"`
	} `ini:"server"`

	Database struct {
		Source string `ini:"source"`
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

	LogPath = cfg.Log.Path

	base := filepath.Join(root, cfg.Bin.Root)
	Ffmpeg = filepath.Join(base, cfg.Bin.Ffmpeg)
	Ffprobe = filepath.Join(base, cfg.Bin.Ffprobe)

	Datasource = "file:" + cfg.Database.Source
}
