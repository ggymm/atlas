package task

import (
	"io/fs"
	"path/filepath"
	"sync"
	"time"

	"github.com/ggymm/gopkg/conv"
	"github.com/ggymm/gopkg/xxhash"

	"atlas/pkg/app"
	"atlas/pkg/data"
	"atlas/pkg/data/model"
	"atlas/pkg/data/service"
	"atlas/pkg/log"
	"atlas/pkg/video"
)

var (
	state = new(sync.Mutex)
)

func run() {
	state.Lock()
	defer state.Unlock()

	data.Flush()
	root := app.Root
	hash := xxhash.New()

	// 扫描目录
	err := filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		id := check(p, hash)
		if len(id) == 0 {
			return nil
		}
		if service.CheckVideo(id) {
			return nil
		}

		// 基础信息
		v := new(model.Video)
		v.Id = id
		v.Name = d.Name()

		rel, _ := filepath.Rel(root, p)
		v.Path = rel

		// 文件信息
		fi, err := d.Info()
		if err != nil {
			log.Error(err).Str("file", p).Msg("file info error")
			return nil
		}
		v.Size = fi.Size()

		// 视频信息
		vi, err := video.Parse(p)
		if err != nil {
			log.Error(err).Str("file", p).Msg("parse video error")
			return nil
		}
		v.Format = vi.Format.FormatLongName
		v.Duration = int64(conv.ParseFloat64(vi.Format.Duration))

		cov, err := video.Thumbnail(p)
		if err != nil {
			log.Error(err).Str("file", p).Msg("thumbnail video error")
			return nil
		}
		v.Cover = cov

		// 保存数据库
		err = service.CreateVideo(v)
		if err != nil {
			log.Error(err).Msg("create video error")
		}
		return nil
	})
	if err != nil {
		log.Error(err).Str("file", root).Msg("walk dir error")
		return
	}
}

func Start() {
	t := time.NewTicker(2 * time.Minute)
	defer t.Stop()

	time.AfterFunc(0, func() {
		run()
	})

	for {
		select {
		case <-t.C:
			if !state.TryLock() {
				continue
			}

			// 执行扫描任务
			run()
		}
	}
}
