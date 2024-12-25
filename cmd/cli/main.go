package main

import (
	"io/fs"
	"path/filepath"

	"atlas/pkg/app"
	"atlas/pkg/store"
	"atlas/pkg/video"
)

func init() {
	app.Init()
	store.Init()
}

func main() {
	store.Flush()

	root := "C:/Users/19679/Videos/TG/202410"
	err := filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		bs, err := video.Thumbnail(p)
		if err != nil {
			return nil
		}

		v := new(store.Video)
		v.Name = d.Name()
		v.RelPath = p
		v.Thumbnail = bs
		err = store.DB.Create(v).Error
		if err != nil {
			return nil
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
