package view

import (
	"embed"
	"net/http"
	"path"
)

//go:embed static/*
var Static embed.FS

var FormatFS = &formatFS{
	fs: http.FS(Static),
}

type formatFS struct {
	fs http.FileSystem
}

func (p *formatFS) Open(name string) (http.File, error) {
	name = path.Join("static", name)
	return p.fs.Open(name)
}

type FileServer struct {
}

func NewFileServer() *FileServer {
	return &FileServer{}
}

func (s *FileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.FileServer(FormatFS).ServeHTTP(w, r)
}
