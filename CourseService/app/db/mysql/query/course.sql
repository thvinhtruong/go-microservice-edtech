-- name: CreateCourse :execresult
INSERT INTO Course (title, description, price, blocked, datecreated) VALUES(?, ?, ?, false, NOW());

-- name: CreateCourseTutor :exec
INSERT INTO Course_Tutor (course_id, tutor_id) VALUES(?, ?);

-- name: CreateCourseUser :exec
INSERT INTO Course_User (course_id, user_id) VALUES(?, ?);

-- name: CreateFeedback :exec
INSERT INTO Feedback (course_id, user_id, content, datecreated) VALUES(?, ?, ?, NOW());

-- name: CreateLecture :exec
INSERT INTO Lecture (course_id, title, content, datecreated) VALUES(?, ?, ?, NOW());

-- name: ListCourses :many
SELECT * FROM Course WHERE blocked = false ORDER BY id LIMIT 1 FOR UPDATE;

-- name: GetCourse :one
SELECT * FROM Course WHERE id = ? AND blocked = 0 LIMIT 1 FOR UPDATE;

-- name: GetCourseTutor :one
SELECT course_id FROM Course_Tutor WHERE tutor_id = ? FOR UPDATE;

-- name: GetCourseUser :one
SELECT course_id FROM Course_User WHERE user_id = ? FOR UPDATE;

-- name: GetFeedback :one
SELECT * FROM Feedback WHERE course_id = ? AND user_id = ? FOR UPDATE;

-- name: GetLecture :one
SELECT * FROM Lecture WHERE course_id = ? FOR UPDATE;

-- name: UpdateCourse :exec
UPDATE Course SET title = ?, description = ?, price = ? WHERE id = ?;

-- name: UpdateCourseTutor :exec
UPDATE Course_Tutor SET tutor_id = ? WHERE course_id = ?;

-- name: UpdateFeedback :exec
UPDATE Feedback SET content = ? WHERE course_id = ? AND user_id = ?;

-- name: DeleteCourse :exec
DELETE FROM Course WHERE id = ?;

-- name: DeleteCourseTutor :exec
DELETE FROM Course_Tutor WHERE course_id = ?;

-- name: DeleteAllFeedbacks :exec
DELETE FROM Feedback WHERE course_id = ?;

-- name: DeleteLecture :exec
DELETE FROM Lecture WHERE course_id = ?;

-- name: DeleteCourseUser :exec
DELETE FROM Course_User WHERE course_id = ?;

-- name: DeleteFeedback :exec
DELETE FROM Feedback WHERE course_id = ? AND user_id = ?;

-- name: DeleteUserFromCourse :exec
DELETE FROM Course_User WHERE course_id = ? AND user_id = ?;

