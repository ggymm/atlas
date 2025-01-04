package log

import (
	"io"
	"os"

	"github.com/ggymm/gopkg/rolling"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"

	"atlas/pkg/app"
)

var log zerolog.Logger

func Init() {
	format := "2006-01-02 15:04:05.000"
	writers := io.MultiWriter(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: format,
	})
	writers = io.MultiWriter(writers, &rolling.Logger{
		Filename:   app.Log(),
		MaxSize:    256, // megabytes
		MaxAge:     30,  // days
		MaxBackups: 128, // files
	})

	zerolog.TimeFieldFormat = format
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	log = zerolog.New(writers).With().Caller().Timestamp().Logger()
}

func Info() *zerolog.Event {
	return log.Info()
}

func Debug() *zerolog.Event {
	return log.Debug()
}

func Error(err error) *zerolog.Event {
	return log.Error().Stack().Err(errors.WithStack(err))
}

func Fatal(err error) *zerolog.Event {
	return log.Fatal().Stack().Err(errors.WithStack(err))
}
