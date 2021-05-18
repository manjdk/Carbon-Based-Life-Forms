package log

import (
	"fmt"

	"github.com/gofrs/uuid"
)

const (
	traceIDPrexix = "traceId:"
)

func NewTraceID() string {
	return fmt.Sprintf("%s%s", traceIDPrexix, uuid.Must(uuid.NewV4()).String())
}
