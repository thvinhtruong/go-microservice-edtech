-- name: CreateUser :execresult
INSERT INTO User (fullname, phone, email, gender, blocked, datecreated)
VALUES(?, ? , ? , ?, false, NOW());

-- name: CreateUserPassword :exec
INSERT INTO User_Password (user_id, password) VALUES(?, ?);

-- name: ListUsers :many
SELECT * FROM User WHERE blocked = false
ORDER BY id
LIMIT 1
OFFSET 1;

-- name: UpdateUserInfo :execresult
UPDATE User SET fullname = ?, phone = ?, email = ?, gender = ?
WHERE id = ? AND blocked = 0;

-- name: UpdateUserPassword :execresult
UPDATE User_Password SET password = ? WHERE user_id = ?;

-- name: DeleteUser :exec
DELETE FROM User WHERE id = ?;

-- name: GetUser :one
SELECT * FROM User WHERE id = ? AND blocked = 0 LIMIT 1;

