package main

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/ggymm/gopkg/conv"
	"github.com/ggymm/gopkg/uuid"
	"github.com/ggymm/gopkg/xxhash"

	"atlas/pkg/app"
	"atlas/pkg/data"
	"atlas/pkg/data/model"
	"atlas/pkg/data/service"
	"atlas/pkg/log"
	"atlas/pkg/video"
)

func init() {
	app.Init()
	log.Init()
	data.Init()
}

func main() {
	data.Flush()
	root := app.Root

	h := xxhash.New()
	// 遍历目录
	err := filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		rel, _ := filepath.Rel(root, p)
		if service.CheckVideo(rel) {
			return nil
		}

		var (
			f  *os.File
			fi fs.FileInfo
		)
		defer func() {
			_ = f.Close()
		}()
		f, err = os.Open(p)
		if err != nil {
			return nil
		}

		// 计算 hash
		h.Reset()
		_, err = io.Copy(h, f)
		if err != nil {
			log.Error(err).Str("file", p).Msg("file hash")
			return nil
		}

		v := new(model.Video)
		v.Id = uuid.NewUUID()
		v.Name = d.Name()
		v.Tags = ""
		v.Path = rel

		fi, err = d.Info()
		if err != nil {
			log.Error(err).Str("file", p).Msg("file info")
			return nil
		}
		v.Size = fi.Size()

		// 视频相关信息
		var (
			vi  *video.Info
			cov []byte
		)
		vi, err = video.Parse(p)
		if err != nil {
			log.Error(err).Str("file", p).Msg("parse video")
			return nil
		}

		cov, err = video.Thumbnail(p)
		if err != nil {
			log.Error(err).Str("file", p).Msg("thumbnail video")
			return nil
		}

		v.Cover = cov
		v.Format = vi.Format.FormatLongName
		v.Duration = int64(conv.ParseFloat64(vi.Format.Duration))

		err = service.CreateVideo(v)
		if err != nil {
			log.Error(err).Msg("create video")
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
