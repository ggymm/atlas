package task

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ggymm/gopkg/conv"
	"github.com/ggymm/gopkg/uuid"

	"atlas/pkg/data/model"
	"atlas/pkg/data/service"
	"atlas/pkg/log"
	"atlas/pkg/video"
)

const (
	tag = "scanner"
)

func (s *Scanner) walk(p string) error {
	fs, err := os.ReadDir(p)
	if err != nil {
		return err
	}
	for _, f := range fs {
		name := f.Name()
		path := filepath.Join(p, name)
		if f.IsDir() {
			err = s.walk(path)
			if err != nil {
				return err
			}
		} else {
			_ = s.parse(f, path)
		}
	}
	return nil
}

func (s *Scanner) check(p string) error {
	f, err := os.Open(p)
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
	return nil
}

func (s *Scanner) parse(f os.DirEntry, p string) error {
	err := s.check(p)
	if err != nil {
		log.Error(err).
			Str("file", p).
			Msgf("%s check video error", tag)
		return err
	}

	// 基础信息
	v := new(model.Video)
	v.Name = f.Name()
	v.Path = Rel(s.root, p)
	v.Star = 0  // 默认未收藏
	v.Tags = "" // 默认无标签
	if service.CheckVideo(v) {
		return nil
	}
	v.Id = uuid.NewUUID()

	// 视频信息
	vi, err := video.Parse(p)
	if err != nil {
		log.Error(err).
			Str("file", p).
			Msgf("%s parse video error", tag)
		return err
	}
	v.Size = conv.ParseInt64(vi.Format.Size)
	v.Format = vi.Format.FormatLongName
	v.Duration = int64(conv.ParseFloat64(vi.Format.Duration))

	cov, err := video.Thumbnail(p)
	if err != nil {
		log.Error(err).
			Str("file", p).
			Msgf("%s thumbnail video error", tag)
		return err
	}
	v.Cover = cov

	// 保存数据库
	err = service.CreateVideo(v)
	if err != nil {
		log.Error(err).
			Str("file", p).
			Msgf("%s create video error", tag)
		return err
	}
	return nil
}
