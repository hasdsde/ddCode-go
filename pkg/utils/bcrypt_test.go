package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratePasswordHash(t *testing.T) {
	passwordHash, err := GeneratePasswordHash("123456")
	assert.Equal(t, nil, err)
	fmt.Println(passwordHash)
}

func TestCheckPasswordHash(t *testing.T) {
	passwordHash, err := GeneratePasswordHash("123456")
	assert.Equal(t, nil, err)
	fmt.Println(passwordHash)
	b := VerifyPassword("123456", passwordHash)
	fmt.Println(b)
}
