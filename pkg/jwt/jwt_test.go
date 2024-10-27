package jwttoken

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker("util.RandomString(32)util.RandomString(32)util.RandomString(32)", time.Second*60)
	require.NoError(t, err)

	username := "util.RandomOwner()"
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	info := map[string]interface{}{
		"userId":   111,
		"userName": username,
		"role": []string{
			"aaaa", "bbbb", "cccc",
		},
	}
	token, payload, err := maker.CreateToken(info)

	payload, err = maker.VerifyToken(token)
	fmt.Println(token, err)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJWTMaker("util.RandomString(32)util.RandomString(32)util.RandomString(32)", time.Second*1)
	require.NoError(t, err)
	info := map[string]interface{}{
		"userId":   111,
		"userName": "username",
	}
	token, payload, err := maker.CreateToken(info)
	fmt.Println(payload, err)
	time.Sleep(2 * time.Second)

	payload, err = maker.VerifyToken(token)
	fmt.Println(err)
}

//func TestInvalidJWTTokenAlgNone(t *testing.T) {
//	payload, err := NewPayload(0, "util.RandomOwner()", "", time.Minute)
//	require.NoError(t, err)
//
//	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
//	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
//	require.NoError(t, err)
//
//	maker, err := NewJWTMaker("util.RandomString(32)util.RandomString(32)util.RandomString(32)")
//	require.NoError(t, err)
//
//	payload, err = maker.VerifyToken(token)
//	require.Error(t, err)
//	require.EqualError(t, err, ErrInvalidToken.Error())
//	require.Nil(t, payload)
//}
