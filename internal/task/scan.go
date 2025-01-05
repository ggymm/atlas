package task

import (
	"atlas/pkg/data"
	"atlas/pkg/data/model"
	"atlas/pkg/data/service"
	"atlas/pkg/log"
	"atlas/pkg/video"
	"github.com/ggymm/gopkg/conv"
	"github.com/ggymm/gopkg/xxhash"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	tag = "scanner"
)

func (s *Scanner) run() error {
	data.Flush()

	// 扫描目录
	err := filepath.Walk(s.root, func(path string, info fs.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		v := new(model.Video)
		err = s.check(v, path)
		if err != nil {
			log.Error(err).
				Str("file", path).
				Msgf("%s check video error", tag)
			return nil
		}
		v.Name = info.Name()
		v.Size = info.Size()

		// 判断是否存在
		if service.CheckVideo(v) {
			err = s.update(v, path)
			if err != nil {
				log.Error(err).
					Str("file", path).
					Msgf("%s update video error", tag)
			}
			return nil
		} else {
			err = s.create(v, path)
			if err != nil {
				log.Error(err).
					Str("file", path).
					Msgf("%s create video error", tag)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	// 清理数据库
	resp, err := service.SelectVideos(nil)
	if err != nil {
		return err
	}
	for _, v := range resp.Records {
		path := filepath.Join(s.root, v.Path)
		if !Exists(path) {
			err = service.DeleteVideo(v.Id)
			if err != nil {
				log.Error(err).
					Str("file", path).
					Msgf("%s delete video error", tag)
			}
		}
	}
	return nil
}

func (s *Scanner) check(v *model.Video, path string) error {
	h := xxhash.New()

	// 检查格式
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	buf := make([]byte, 512)
	_, err = f.Read(buf)
	if err != nil {
		return err
	}

	mime := http.DetectContentType(buf)
	if !strings.HasPrefix(mime, "video/") {
		return err
	}

	// 计算文件 hash
	_, err = io.Copy(h, f)
	if err != nil {
		return err
	}
	v.Id = h.SumHex()
	return nil
}

func (s *Scanner) create(v *model.Video, path string) error {
	// 相对路径
	rel, _ := filepath.Rel(s.root, path)
	v.Path = rel

	// 视频信息
	vi, err := video.Parse(path)
	if err != nil {
		log.Error(err).
			Str("file", path).
			Msgf("%s parse video error", tag)
		return nil
	}
	v.Format = vi.Format.FormatLongName
	v.Duration = int64(conv.ParseFloat64(vi.Format.Duration))

	cov, err := video.Thumbnail(path)
	if err != nil {
		log.Error(err).
			Str("file", path).
			Msgf("%s thumbnail video error", tag)
		return nil
	}
	v.Cover = cov

	// 保存数据库
	return service.CreateVideo(v)
}

func (s *Scanner) update(v *model.Video, path string) error {
	// 相对路径
	rel, _ := filepath.Rel(s.root, path)
	v.Path = rel
	return service.UpdateVideo(v)
}

func (s *Scanner) Start() error {
	i, err := os.Stat(s.root)
	if err != nil {
		return err
	}

	// 判断是否是目录
	if !i.IsDir() {
		return os.ErrNotExist
	}

	// 判断权限是否符合要求
	if i.Mode().Perm()&os.ModePerm != os.ModePerm {
		return os.ErrPermission
	}
	for {
		err = s.run()
		if err != nil {
			log.Error(err).
				Str("root", s.root).
				Msgf("%s scanner run error", tag)
		}
		time.Sleep(1 * time.Minute)
	}
}
