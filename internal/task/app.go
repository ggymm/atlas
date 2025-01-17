package task

import (
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"atlas/pkg/app"
	"atlas/pkg/data"
	"atlas/pkg/data/model"
)

type Scanner struct {
	root string
}

func NewScanner() *Scanner {
	return &Scanner{
		root: app.Root,
	}
}

func (s *Scanner) Test() error {
	err := s.walk(s.root)
	if err != nil {
		slog.Error("run error",
			slog.Any("error", err),
			slog.String("task", tag),
			slog.String("root", s.root),
		)
		return err
	}

	videos := make([]*model.Video, 0)
	err = data.DB.Find(&videos).Error
	if err != nil {
		slog.Error("query videos error",
			slog.Any("error", err),
			slog.String("task", tag),
		)
		return err
	}
	for _, v := range videos {
		path := filepath.Join(s.root, v.Path)
		if !exists(path) {
			err = data.DB.Delete(v).Error
			if err != nil {
				slog.Error("delete video error",
					slog.Any("error", err),
					slog.String("task", tag),
				)
				continue
			}
		}
	}
	return nil
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
		slog.Info("[task] scanner start")
		err = s.walk(s.root)
		if err != nil {
			slog.Error("run error",
				slog.Any("error", err),
				slog.String("task", tag),
				slog.String("root", s.root),
			)
		}

		videos := make([]*model.Video, 0)
		err = data.DB.Find(&videos).Error
		if err != nil {
			slog.Error("query videos error",
				slog.Any("error", err),
				slog.String("task", tag),
			)
		}
		for _, v := range videos {
			path := filepath.Join(s.root, v.Path)
			if !exists(path) {
				err = data.DB.Delete(v).Error
				if err != nil {
					slog.Error("delete video error",
						slog.Any("error", err),
						slog.String("task", tag),
					)
					continue
				}
			}
		}
		data.Flush()
		time.Sleep(30 * time.Minute)
	}
}
