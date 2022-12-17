-- name: CreateTutor :execresult
INSERT INTO Tutor (fullname, phone, gender, validate, adminId, datecreated) VALUES(?, ?, ?, false, 0, NOW());

-- name: CreateTutorPassword :execresult
INSERT INTO Tutor_Password (tutor_id, password) VALUES(?, ?);

-- name: GetTutor :one
SELECT * FROM Tutor WHERE id = ? AND blocked = 0 LIMIT 1;

-- name: UpdateTutorInfo :execresult
UPDATE Tutor SET fullname = ?, phone = ?, gender = ? WHERE id = ? AND blocked = 0;

-- name: DeleteTutor :exec
DELETE FROM Tutor WHERE id = ?;

-- name: UpdateTutorPassword :exec
UPDATE Tutor_Password SET password = ? WHERE tutor_id = ?;