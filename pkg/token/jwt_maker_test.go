package token

import (
	"mygram-api/pkg/random"
	"strconv"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(random.RandomString(32))
	require.NoError(t, err)

	userId := random.RandomInt(1, 5)
	email := random.RandomEmail()
	age := random.RandomInt(10, 20)
	audience := "access"
	username := random.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(int(userId), email, username, int(age), audience, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, strconv.Itoa(int(userId)), payload.Subject)
	require.WithinDuration(t, issuedAt, payload.IssuedAt.Time, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiresAt.Time, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJWTMaker(random.RandomString(32))
	require.NoError(t, err)

	userId := random.RandomInt(1, 5)
	email := random.RandomEmail()
	age := random.RandomInt(10, 20)
	audience := "access"
	username := random.RandomOwner()
	duration := -time.Minute

	token, err := maker.CreateToken(int(userId), email, username, int(age), audience, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	userId := random.RandomInt(1, 5)
	email := random.RandomEmail()
	age := random.RandomInt(10, 20)
	audience := "access"
	username := random.RandomOwner()
	duration := time.Minute

	payload, err := NewPayload(int(userId), email, username, int(age), audience, duration)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(random.RandomString(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}
