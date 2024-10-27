package jwttoken

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("登录信息无效")
	ErrExpiredToken = errors.New("登录信息已过期")
)

type Payload struct {
	Info      map[string]interface{}
	IssuedAt  time.Time
	ExpiredAt time.Time
	TokenId   uuid.UUID
}

func NewPayload(duration time.Duration, info map[string]interface{}) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		Info:      info,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
		TokenId:   tokenID,
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
