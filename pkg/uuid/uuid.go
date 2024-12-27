package uuid

import (
	"github.com/ggymm/gopkg/uuid"
)

// New returns a new UUID（without hyphen）
func New() string {
	return uuid.NewUUID()
}
