package utils

import (
	"testing"

	"github.com/google/uuid"
)

func TestNew(t *testing.T) {
	m := make(map[string]bool)
	for x := 1; x < 32; x++ {
		u := NewUUID()
		if m[u] {
			t.Errorf("New returned duplicated UUID %q", u)
		}
		m[u] = true
		_, err := uuid.Parse(u)
		if err != nil {
			t.Errorf("New returned %q which does not decode", u)
			continue
		}
	}
}
