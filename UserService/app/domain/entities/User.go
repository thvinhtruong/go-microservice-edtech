package entity

type User struct {
	ID          int
	FullName    string
	Email       string
	UserName    string
	Gender      string
	Blocked     bool
	DateCreated int64
	DateUpdated int64
}
