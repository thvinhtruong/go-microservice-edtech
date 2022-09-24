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
	sql_get_users       = "SELECT * FROM user"
	sql_get_user        = "SELECT * FROM user WHERE id = ?"
	sql_update_password = "UPDATE password SET password = ? WHERE user_id = ?"
	sql_delete_user     = "DELETE FROM user WHERE id = ?"
)

func insertUser(q *Querier, ctx context.Context, user entity.User) (int, error) {
	date := conversion.ConvertUnixTimeMySqlTime(time.Now().Unix())

	user_result, err := q.DB.ExecContext(ctx, sql_create_user,
		user.FullName, user.Email, user.Username, user.Gender, 1, date)
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

func getAllUsers(q *Querier, ctx context.Context) ([]entity.User, error) {
	rows, err := q.DB.QueryContext(ctx, sql_get_users)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var result []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.FullName, &user.Email, &user.Username, &user.Gender, &user.Blocked, &user.DateCreated)
		if err != nil {
			return nil, err
		}

		result = append(result, user)
	}
	return result, nil
}

func getUser(q *Querier, ctx context.Context, id int) (entity.User, error) {
	var user entity.User
	err := q.DB.QueryRowContext(ctx, sql_get_user, id).Scan(&user.FullName, &user.Email,
		&user.Username, &user.Gender)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func updatePassword(q *Querier, ctx context.Context, userId int, newPassword string) error {
	user, err := getUser(q, ctx, userId)
	if err != nil {
		return err
	}

	_, err = q.DB.ExecContext(ctx, sql_update_password, newPassword, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func deleteUser(q *Querier, ctx context.Context, id int) error {
	_, err := q.DB.ExecContext(ctx, sql_delete_user, id)
	if err != nil {
		return err
	}

	return nil
}

func blockUser(q *Querier, ctx context.Context, id int) error {
	stmt := "UPDATE user SET blocked = 1 WHERE id = ?"
	_, err := q.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}
