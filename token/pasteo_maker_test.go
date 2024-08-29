package token

import (
	"fmt"
	"testing"
	"time"

	"github.com/REZ-OAN/simplebank/utils"
	"github.com/aead/chacha20poly1305"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(utils.RandomString(32))
	require.NoError(t, err)

	username := utils.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(utils.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(utils.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidPASETOToken(t *testing.T) {
	// Generate a valid payload
	payload, err := NewPayload(utils.RandomOwner(), time.Minute)
	require.NoError(t, err)

	// Create a PASETO maker with a random key
	maker, err := NewPasetoMaker(utils.RandomString(32))
	require.NoError(t, err)

	// Create a valid PASETO token
	token, err := maker.CreateToken(payload.Username, time.Minute)
	require.NoError(t, err)

	// Tamper the token by modifying its contents to simulate an invalid token
	tamperedToken := token[:len(token)-1] // Remove the last character

	// Attempt to verify the tampered token
	payload, err = maker.VerifyToken(tamperedToken)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}

func TestNewPasetoMakerInvalidKeySize(t *testing.T) {
	// Create a key with an invalid size (too short)
	invalidKey := utils.RandomString(chacha20poly1305.KeySize - 1)

	// Attempt to create a new PASETO maker with the invalid key
	maker, err := NewPasetoMaker(invalidKey)

	// Ensure that the error is not nil and that it returns the expected error message
	require.Nil(t, maker)
	require.Error(t, err)
	require.EqualError(t, err, fmt.Sprintf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize))
}
