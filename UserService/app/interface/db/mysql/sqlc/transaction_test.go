package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegisterUser(t *testing.T) {
	store := NewTxStore(testDB)

	user1 := RegisterUserParams{
		Fullname: "test1",
		Phone:    "767637216736",
		Email:    "test4@gmail.com",
		Password: "test",
		Gender:   "male",
	}

	errs := make(chan error)
	results := make(chan RegisterUserResult)

	go func() {
		result, err := store.RegisterUser(context.Background(), user1)
		if err != nil {
			t.Error(err)
		}

		errs <- err
		results <- result

	}()

	err := <-errs
	require.NoError(t, err)
	result := <-results
	require.NotEmpty(t, result)
}

func TestDeleteUser(t *testing.T) {
	store := NewTxStore(testDB)

	err := store.DeleteUser(context.Background(), 1)
	require.NoError(t, err)
}
