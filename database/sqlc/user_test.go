package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/REZ-OAN/simplebank/utils"

	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	firstName, lastName := utils.RandomFullName()

	arg := CreateUserParams{
		Username:       utils.RandomUserName(firstName),
		HashedPassword: "secret",
		Email:          utils.RandomEmail(),
		FullName:       fmt.Sprintf("%s %s", firstName, lastName),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)
	return user
}
func TestCreateUser(t *testing.T) {

	// to use the create random user for other testing
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := CreateRandomUser(t)

	guser, err := testQueries.GetUser(context.Background(), user.Username)

	require.NoError(t, err)
	require.NotEmpty(t, guser)

	require.Equal(t, user.Username, guser.Username)
	require.Equal(t, user.Email, guser.Email)
	require.WithinDuration(t, user.CreatedAt, guser.CreatedAt, time.Second)
	require.WithinDuration(t, user.PasswordChangedAt, guser.PasswordChangedAt, time.Second)
	require.Equal(t, user.FullName, guser.FullName)
	require.Equal(t, user.HashedPassword, guser.HashedPassword)

}
