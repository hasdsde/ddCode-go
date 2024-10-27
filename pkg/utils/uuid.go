package utils

import (
	"sync"

	"github.com/google/uuid"
)

var uuidLock sync.Mutex
var lastUUID uuid.UUID

// NewUUID returns a version 1 UUID.
func NewUUID() string {
	uuidLock.Lock()
	defer uuidLock.Unlock()
	result := uuid.Must(uuid.NewUUID())
	// The UUID package is naive and can generate identical UUIDs if the
	// time interval is quick enough.
	// The UUID uses 100 ns increments so it's short enough to actively
	// wait for a new value.
	for lastUUID == result {
		result = uuid.Must(uuid.NewUUID())
	}
	lastUUID = result
	return result.String()
}
