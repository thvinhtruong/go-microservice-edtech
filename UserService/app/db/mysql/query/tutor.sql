-- name: CreateTutor :execresult
INSERT INTO Tutor (fullname, phone, gender, age, topic, country, city, datecreated) VALUES(?, ?, ?, ?, ?, ?, ?, NOW());

-- name: CreateTutorPassword :execresult
INSERT INTO Tutor_Password (tutor_id, password) VALUES(?, ?);

-- name: GetTutor :one
SELECT * FROM Tutor WHERE id = ? LIMIT 1 FOR UPDATE;

-- name: GetTutorByPhone :one
SELECT * FROM Tutor WHERE phone = ? LIMIT 1 FOR UPDATE;

-- name: GetTutorPassword :one
SELECT * FROM Tutor_Password WHERE tutor_id = ? LIMIT 1;

-- name: UpdateTutorInfo :execresult
UPDATE Tutor SET fullname = ?, phone = ?, gender = ?, topic = ?, age = ?, country = ?, city = ? WHERE id = ?;

-- name: DeleteTutor :exec
DELETE FROM Tutor WHERE id = ?;

-- name: UpdateTutorPassword :exec
UPDATE Tutor_Password SET password = ? WHERE tutor_id = ?;