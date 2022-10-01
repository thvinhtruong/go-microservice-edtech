-- name: CreateTutor :execresult
INSERT INTO Tutor (fullname, username, email, phone, gender) VALUES(? , ? , ? , ? , ? );

-- name: CreateTutorPassword :execresult
INSERT INTO Tutor_Password (tutor_id, password) VALUES(?, ?);

-- name: GetTutor :one
SELECT * FROM Tutor WHERE id = ? AND blocked = 0 LIMIT 1;

-- name: UpdateTutor :execresult
UPDATE Tutor SET fullname = ?, username = ?, email = ?, phone = ?, gender = ? WHERE id = ? AND blocked = 0;

-- name: DeleteTutor :exec
DELETE FROM Tutor WHERE id = ?

