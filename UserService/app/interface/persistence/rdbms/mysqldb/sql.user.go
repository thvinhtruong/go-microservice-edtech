package mysqldb

import (
	entity "SepFirst/UserService/app/domain/entities"
	"SepFirst/UserService/pkg/conversion"
	"context"
	"time"
)

const (
	sql_create_user     = "INSERT INTO user (fullname, email, password, username, gender, blocked , datecreated) VALUES (?, ?, ?, ?, ?, ?, ?)"
	sql_create_password = "INSERT INTO password (user_id, password, datecreated) VALUES (?, ?, ?)"
)

func insertUser(q *Querier, ctx context.Context, user entity.User) (int, error) {
	date := conversion.ConvertUnixTimeMySqlTime(time.Now().Unix())

	user_result, err := q.DB.ExecContext(ctx, sql_create_user, user.FullName, user.Email, user.Username, user.Gender, 1, date)
	if err != nil {
		return 0, err
	}

	user_id, err := user_result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(user_id), nil
}

func insertPassword(q *Querier, ctx context.Context, userId int, password string) error {
	date := conversion.ConvertUnixTimeMySqlTime(time.Now().Unix())
	password_obj := entity.Password{
		UserId:   userId,
		Password: password,
	}
	password_obj.HashPassword()
	_, err := q.DB.ExecContext(ctx, sql_create_password, password_obj.UserId, password_obj.Password, date)
	if err != nil {
		return err
	}

	return nil
}
