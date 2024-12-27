package main

import (
	"io/fs"
	"path/filepath"

	"atlas/pkg/app"
	"atlas/pkg/data"
	"atlas/pkg/data/model"
	"atlas/pkg/video"
)

func init() {
	app.Init()
	data.Init()

}

func main() {
	data.Flush()

	root := "C:/Users/19679/Videos/TG/202410"
	err := filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		bs, err := video.Thumbnail(p)
		if err != nil {
			return nil
		}

		v := new(model.Video)
		v.Name = d.Name()
		v.Path = p
		v.Thumbnail = bs
		err = v.Create()
		if err != nil {
			return nil
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
