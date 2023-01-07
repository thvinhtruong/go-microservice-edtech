package db

import (
	"context"
	"testing"

	"server/UserService/pkg/random"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) int64 {
	arg := CreateUserParams{
		Phone:    random.RandomPhone(),
		Fullname: random.RandomName(),
		Gender:   random.RandomGender(),
	}

	record, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)

	ids, err := record.LastInsertId()

	require.NoError(t, err)
	require.NotEmpty(t, record)
	require.NotNil(t, ids)
	require.NotEqual(t, ids, -1)
	require.NotEqual(t, ids, 0)

	return ids

}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	arg := CreateUserParams{
		Fullname: random.RandomName(),
		Phone:    random.RandomPhone(),
		Gender:   random.RandomGender(),
	}

	account, err := testQueries.CreateUser(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
}
