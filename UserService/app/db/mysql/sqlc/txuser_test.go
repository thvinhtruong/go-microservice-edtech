package db

import (
	"context"
	"fmt"
	"testing"

	"server/UserService/pkg/random"

	"github.com/stretchr/testify/require"
)

func TestRegisterUser(t *testing.T) {
	testDB, err := CreateTestDB()
	if err != nil {
		t.Error(err)
	}

	userUseCase := NewRepository(testDB)
	user1 := RegisterUserParams{
		Fullname: "test1",
		Phone:    random.RandomPhone(),
		Password: "test",
		Gender:   "male",
	}

	errs := make(chan error)
	results := make(chan RegisterUserResult)

	go func() {
		result, err := userUseCase.RegisterUser(context.Background(), user1)
		if err != nil {
			t.Error(err)
		}

		errs <- err
		results <- result

	}()

	err = <-errs
	require.NoError(t, err)
	result := <-results
	require.NotEmpty(t, result)

	fmt.Println(result)
}

func TestFilterTutors(t *testing.T) {
	testDB, err := CreateTestDB()
	if err != nil {
		t.Error(err)
	}

	userUseCase := NewRepository(testDB)
	got, err := userUseCase.FilterTutors(context.Background(), FindTutorsMatchParams{
		Gender:  "female",
		Topic:   "Math",
		Country: "Vietnam",
		City:    "Paris",
		Age:     100,
	})

	if err != nil {
		t.Error(err)
	}

	want := []int32{64}

	require.Equal(t, want, got)

}
