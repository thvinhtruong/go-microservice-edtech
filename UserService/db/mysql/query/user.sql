-- name: CreateUser :execresult
INSERT INTO User (fullname, username, phone, email, gender)
VALUES(? , ? , ? , ? , ?);

-- name: CreateUserPassword :execresult
INSERT INTO User_Password (user_id, password) VALUES(?, ?);

-- name: ListUsers :many
SELECT * FROM User WHERE blocked = false
ORDER BY id DESC;

-- name: UpdateUserInfo :execresult
UPDATE User SET fullname = ?, username = ?, phone = ?, email = ?, gender = ?
WHERE id = ? AND blocked = 0;

-- name: UpdateUserPassword :execresult
UPDATE User_Password SET password = ? WHERE user_id = ?;

-- name: DeleteUser :exec
DELETE FROM User WHERE id = ?;

-- name: GetUser :one
SELECT * FROM User WHERE id = ? AND blocked = 0 LIMIT 1;

