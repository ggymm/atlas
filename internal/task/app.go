package task

import (
	"log/slog"
	"os"
	"path/filepath"

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

func (s *Scanner) Start() error {
	st, err := os.Stat(s.root)
	if err != nil {
		return err
	}

	// 判断是否是目录
	if !st.IsDir() {
		return os.ErrNotExist
	}

	// 判断权限是否符合要求
	if st.Mode().Perm()&os.ModePerm != os.ModePerm {
		return os.ErrPermission
	}

	// 执行扫描
	slog.Info("[task] scanner start")
	err = s.walk(s.root)
	if err != nil {
		slog.Error("run error",
			slog.Any("error", err),
			slog.String("task", tag),
			slog.String("root", s.root),
		)
	}

	// 执行清理
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
	return nil
}
