package main

import (
	"github.com/ggymm/gopkg/conv"
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

	base := "C:/Users/19679/Videos/TG/"
	root := "C:/Users/19679/Videos/TG/202410"
	err := filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		fi, _ := d.Info()

		vi, err := video.Parse(p)
		if err != nil {
			return nil
		}

		bs, err := video.Thumbnail(p)
		if err != nil {
			return nil
		}

		rel, _ := filepath.Rel(base, p)

		v := new(model.Video)
		v.Name = d.Name()
		v.Path = rel
		v.Size = fi.Size()
		v.Format = vi.Format.FormatLongName
		v.Duration = conv.ParseInt64(vi.Format.Duration)
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
