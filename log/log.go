package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func ErrorZ(traceID string, err error) *zerolog.Event {
	return log.Error().
		Str("traceId", traceID).
		Err(err)
}

func InfoZ(traceID string) *zerolog.Event {
	return log.Info().
		Str("traceId", traceID)
}

func FatalZ(err error) *zerolog.Event {
	return log.Fatal().
		Err(err)
}
