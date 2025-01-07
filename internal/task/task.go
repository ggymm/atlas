package task

import (
	"os"
	"path/filepath"
	"time"

	"atlas/pkg/app"
	"atlas/pkg/data"
	"atlas/pkg/data/model"
	"atlas/pkg/data/service"
	"atlas/pkg/log"
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
		log.Error(err).
			Str("root", s.root).
			Msgf("%s scanner run error", tag)
		return err
	}

	videos := make([]*model.Video, 0)
	err = data.DB.Find(&videos).Error
	if err != nil {
		log.Error(err).
			Msgf("%s query videos error", tag)
		return err
	}
	for _, v := range videos {
		path := filepath.Join(s.root, v.Path)
		if !Exists(path) {
			err = service.DeleteVideo(v.Id)
			if err != nil {
				log.Error(err).
					Str("file", path).
					Msgf("%s delete video error", tag)
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
		log.Info().Msgf("[task] start scanner")
		err = s.walk(s.root)
		if err != nil {
			log.Error(err).
				Str("root", s.root).
				Msgf("%s scanner run error", tag)
		}

		videos := make([]*model.Video, 0)
		err = data.DB.Find(&videos).Error
		if err != nil {
			log.Error(err).
				Msgf("%s query videos error", tag)
		}
		for _, v := range videos {
			path := filepath.Join(s.root, v.Path)
			if !Exists(path) {
				err = service.DeleteVideo(v.Id)
				if err != nil {
					log.Error(err).
						Str("file", path).
						Msgf("%s delete video error", tag)
					continue
				}
			}
		}
		data.Flush()
		time.Sleep(30 * time.Minute)
	}
}
