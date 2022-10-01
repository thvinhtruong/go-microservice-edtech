package db

import (
	"context"
	"testing"

	"SepFirst/UserService/pkg/random"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	arg := CreateUserParams{
		Fullname: random.RandomName(),
		Phone:    random.RandomPhone(),
		Email:    random.RandomEmail(),
		Gender:   random.RandomGender(),
	}

	account, err := testQueries.CreateUser(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
}
