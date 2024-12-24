package app

import (
	"path/filepath"

	"github.com/ggymm/gopkg/ini"
)

var (
	Name string
)

var (
	Ffmpeg  string
	Ffprobe string
)

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
}

func Init() {
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

	Ffmpeg = filepath.Join(root, cfg.Bin.Root, cfg.Bin.Ffmpeg)
	Ffprobe = filepath.Join(root, cfg.Bin.Root, cfg.Bin.Ffprobe)
}
