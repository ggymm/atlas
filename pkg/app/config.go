package app

import (
	"path/filepath"

	"github.com/ggymm/gopkg/ini"
)

var (
	Name   string
	Addr   string
	View   string
	Root   string
	Player string
)

var (
	LogPath string
)

var (
	Ffmpeg  string
	Ffprobe string
	Webview string
)

var (
	Datasource string
)

func Log() string {
	return filepath.Join(root, LogPath, Name+".log")
}

func DatabaseLog() string {
	return filepath.Join(root, LogPath, Name+"-db.log")
}

type Config struct {
	// App 应用配置
	App struct {
		Name   string `ini:"name"`
		Addr   string `ini:"addr"`
		View   string `ini:"view"`
		Root   string `ini:"root"`
		Player string `ini:"player"`
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
		Webview string `ini:"webview"`
	} `ini:"bin"`

	// Database 数据库配置
	Database struct {
		Source string `ini:"source"`
	} `ini:"database"`
}

func InitConfig() {
	var (
		cfg  = new(Config)
		path = filepath.Join(root, "config.ini")
	)

	err := ini.MapTo(cfg, path)
	if err != nil {
		panic(err)
	}

	Name = cfg.App.Name
	Addr = cfg.App.Addr
	View = cfg.App.View
	Root = cfg.App.Root
	Player = cfg.App.Player

	LogPath = cfg.Log.Path

	base := filepath.Join(root, cfg.Bin.Root)
	Ffmpeg = filepath.Join(base, cfg.Bin.Ffmpeg)
	Ffprobe = filepath.Join(base, cfg.Bin.Ffprobe)
	Webview = filepath.Join(base, cfg.Bin.Webview)

	Datasource = "file:" + filepath.Join(root, cfg.Database.Source)
}
