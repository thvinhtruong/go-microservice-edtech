package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	arg := CreateUserParams{
		Fullname: "Bao Tam",
		Username: "baotam",
		Phone:    "0123456789",
		Email:    "baotam@example.com",
		Gender:   "male",
	}

	account, err := testQueries.CreateUser(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
}
