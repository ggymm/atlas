package data

import (
	"context"
	"log/slog"
	"strings"
	"time"

	"github.com/ggymm/gopkg/log"
	"github.com/pkg/errors"
	"gorm.io/gorm/logger"

	"atlas/pkg/app"
)

type CustomLog struct {
	log      *slog.Logger
	LogLevel logger.LogLevel
}

func NewCustomLog() *CustomLog {
	return &CustomLog{
		log:      log.New(app.DatabaseLog(), slog.LevelInfo),
		LogLevel: logger.Info,
	}
}

func (l *CustomLog) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *CustomLog) Info(_ context.Context, msg string, data ...interface{}) {
	l.log.Info(msg, data...)
}

func (l *CustomLog) Warn(_ context.Context, msg string, data ...interface{}) {
	l.log.Warn(msg, data...)
}

func (l *CustomLog) Error(_ context.Context, msg string, data ...interface{}) {
	l.log.Error(msg, data...)
}

// Trace print sql message
func (l *CustomLog) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}
	cost := time.Since(begin).Milliseconds()

	s, r := fc()
	s = strings.Replace(s, "\"", "'", -1)

	if err != nil {
		if l.LogLevel >= logger.Error && !errors.Is(err, logger.ErrRecordNotFound) {
			l.log.Error("SQLTrace",
				slog.Any("error", err),
				slog.String("sql", s),
				slog.Int64("cost", cost),
				slog.Int64("rowsAffected", r),
			)
		}
	} else {
		if l.LogLevel >= logger.Info {
			l.log.Info("SQLTrace",
				slog.String("sql", s),
				slog.Int64("cost", cost),
				slog.Int64("rowsAffected", r),
			)
		}
	}
}
