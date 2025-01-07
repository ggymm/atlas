package data

import (
	"context"
	"github.com/pkg/errors"
	"io"
	"os"
	"strings"
	"time"

	"github.com/ggymm/gopkg/rolling"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gorm.io/gorm/logger"

	"atlas/pkg/app"
)

type CustomLog struct {
	log      zerolog.Logger
	LogLevel logger.LogLevel
}

func NewCustomLog() *CustomLog {
	writers := io.MultiWriter(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.DateTime,
	})
	writers = io.MultiWriter(writers, &rolling.Logger{
		Filename:   app.DatabaseLog(),
		MaxSize:    256, // megabytes
		MaxAge:     30,  // days
		MaxBackups: 128, // files
	})

	zerolog.TimeFieldFormat = time.DateTime
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	return &CustomLog{
		log:      zerolog.New(writers).With().Caller().Timestamp().Logger(),
		LogLevel: logger.Error,
	}
}

func (l *CustomLog) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *CustomLog) Info(_ context.Context, msg string, data ...interface{}) {
	l.log.Info().Msgf(msg, data...)
}

func (l *CustomLog) Warn(_ context.Context, msg string, data ...interface{}) {
	l.log.Warn().Msgf(msg, data...)
}

func (l *CustomLog) Error(_ context.Context, msg string, data ...interface{}) {
	l.log.Error().Msgf(msg, data...)
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
			// 忽略记录不存在的错误
			l.log.Error().Err(errors.WithStack(err)).
				Str("sql", s).
				Int64("cost", cost).
				Int64("rowsAffected", r).Msg("SQLTrace")
		}
	} else {
		if l.LogLevel >= logger.Info {
			l.log.Info().
				Str("sql", s).
				Int64("cost", cost).
				Int64("rowsAffected", r).Msg("SQLTrace")
		}
	}
}
