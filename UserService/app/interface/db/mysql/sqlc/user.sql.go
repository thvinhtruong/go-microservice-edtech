// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :execresult
INSERT INTO User (fullname, phone, email, gender, blocked, datecreated)
VALUES(?, ? , ? , ?, false, NOW())
`

type CreateUserParams struct {
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser,
		arg.Fullname,
		arg.Phone,
		arg.Email,
		arg.Gender,
	)
}

const createUserPassword = `-- name: CreateUserPassword :exec
INSERT INTO User_Password (user_id, password) VALUES(?, ?)
`

type CreateUserPasswordParams struct {
	UserID   int32  `json:"user_id"`
	Password string `json:"password"`
}

func (q *Queries) CreateUserPassword(ctx context.Context, arg CreateUserPasswordParams) error {
	_, err := q.db.ExecContext(ctx, createUserPassword, arg.UserID, arg.Password)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM User WHERE id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, fullname, gender, email, phone, blocked, datecreated, dateupdated FROM User WHERE id = ? AND blocked = 0 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Fullname,
		&i.Gender,
		&i.Email,
		&i.Phone,
		&i.Blocked,
		&i.Datecreated,
		&i.Dateupdated,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, fullname, gender, email, phone, blocked, datecreated, dateupdated FROM User WHERE blocked = false
ORDER BY id
LIMIT 1
OFFSET 1
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Fullname,
			&i.Gender,
			&i.Email,
			&i.Phone,
			&i.Blocked,
			&i.Datecreated,
			&i.Dateupdated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserInfo = `-- name: UpdateUserInfo :execresult
UPDATE User SET fullname = ?, phone = ?, email = ?, gender = ?
WHERE id = ? AND blocked = 0
`

type UpdateUserInfoParams struct {
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	ID       int32  `json:"id"`
}

func (q *Queries) UpdateUserInfo(ctx context.Context, arg UpdateUserInfoParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateUserInfo,
		arg.Fullname,
		arg.Phone,
		arg.Email,
		arg.Gender,
		arg.ID,
	)
}

const updateUserPassword = `-- name: UpdateUserPassword :execresult
UPDATE User_Password SET password = ? WHERE user_id = ?
`

type UpdateUserPasswordParams struct {
	Password string `json:"password"`
	UserID   int32  `json:"user_id"`
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateUserPassword, arg.Password, arg.UserID)
}