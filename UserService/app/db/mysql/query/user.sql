-- name: CreateUser :execresult
INSERT INTO User (fullname, phone, gender, blocked, datecreated)
VALUES(?, ? , ?, false, NOW());

-- name: CreateUserPassword :exec
INSERT INTO User_Password (user_id, password) VALUES(?, ?);

-- name: ListUsers :many
SELECT * FROM User WHERE blocked = false
ORDER BY id
LIMIT 1
FOR UPDATE;

-- name: GetUserByPhone :one
SELECT * FROM User WHERE phone = ? AND blocked = 0 LIMIT 1
FOR UPDATE;

-- name: GetUserPassword :one
SELECT * FROM User_Password WHERE user_id = ? LIMIT 1;

-- name: UpdateUserInfo :exec
UPDATE User SET fullname = ?, phone = ?, gender = ?, dateupdated = NOW()
WHERE id = ? AND blocked = 0;

-- name: UpdateUserPassword :exec
UPDATE User_Password SET password = ? WHERE user_id = ?;

-- name: DeleteUser :exec
DELETE FROM User WHERE id = ?;

-- name: GetUser :one
SELECT * FROM User WHERE id = ? AND blocked = 0 LIMIT 1
FOR UPDATE;

-- name: DeleteUserPassword :exec
DELETE FROM User_Password WHERE user_id = ?;