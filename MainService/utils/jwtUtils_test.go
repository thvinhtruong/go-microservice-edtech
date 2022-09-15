package utils

import (
	"testing"
)

func TestDecodeToken(t *testing.T) {
	j := JwtUtils{config: TestingConfig{}}
	want := InfoInJwt{UserId: 202}
	token, err := j.GenerateToken(want)
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}
	got, err := j.DecodeToken(token)
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}

	if !equalCredentials(*got, want) {
		t.Errorf("Got %v want %v", got, want)
	}
}

func TestGenerateToken(t *testing.T) {
	j := JwtUtils{config: TestingConfig{}}
	want := InfoInJwt{UserId: 1}
	token, err := j.GenerateToken(want)
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}
	got, err := j.DecodeToken(token)
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}

	if !equalCredentials(*got, want) {
		t.Errorf("Got %v want %v", got, want)
	}
}

func equalCredentials(got, want InfoInJwt) bool {
	return got.UserId == want.UserId
}
