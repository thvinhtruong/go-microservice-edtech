package entity

import (
	hasher "SepFirst/UserService/pkg/utils"
)

type Password struct {
	ID          int
	Password    string
	UserId      int
	DateCreated int64
	DateUpdated int64
}

// hash password
func (p *Password) HashPassword() error {
	hashed, err := hasher.HashPassword(p.Password)
	if err != nil {
		panic(err)
	}
	p.Password = hashed

	return nil
}

// compare password
func (p *Password) ComparePassword() bool {
	compare, err := hasher.ComparePassword(p.Password, p.Password)
	if err != nil {
		panic(err)
	}
	return compare
}