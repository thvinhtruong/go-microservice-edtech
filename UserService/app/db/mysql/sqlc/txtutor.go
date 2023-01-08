package db

import (
	"context"
	"errors"
	"server/UserService/pkg/hasher"
)

func (u *TxStore) RegisterTutor(ctx context.Context, arg RegisterTutorParams) (RegisterTutorResult, error) {
	var result RegisterTutorResult
	err := u.enableTx(ctx, func(q *Queries) error {
		tutor, err := q.CreateTutor(ctx, CreateTutorParams{
			Fullname: arg.Fullname,
			Phone:    arg.Phone,
			Gender:   arg.Gender,
			Age:      arg.Age,
			Topic:    arg.Topic,
			City:     arg.City,
			Country:  arg.Country,
		})
		if err != nil {
			return err
		}

		id, err := tutor.LastInsertId()
		if err != nil {
			return err
		}

		hashed, err := hasher.HashPassword(arg.Password)
		if err != nil {
			return err
		}

		_, err = u.CreateTutorPassword(ctx, CreateTutorPasswordParams{
			TutorID:  int32(id),
			Password: hashed,
		})
		if err != nil {
			return err
		}

		result.ID = int32(id)

		return err
	})

	return result, err

}

func (u *TxStore) LoginTutor(ctx context.Context, arg LoginTutorParams) (LoginTutorResult, error) {
	var result LoginTutorResult
	err := u.enableTx(ctx, func(q *Queries) error {
		tutor, err := q.GetTutorByPhone(ctx, arg.Phone)
		if err != nil {
			return err
		}

		isRightPhone := tutor.Phone == arg.Phone
		if !isRightPhone {
			return errors.New("wrong phone or password")
		}

		tutorPassword, err := q.GetTutorPassword(ctx, tutor.ID)
		if err != nil {
			return err
		}

		isRightPassword, err := hasher.ComparePassword(tutorPassword.Password, arg.Password)
		if err != nil {
			return errors.New("wrong phone or password")
		}

		if !isRightPassword {
			return errors.New("wrong phone or password")
		}

		result.ID = tutor.ID

		return nil

	})

	return result, err
}
