package dbmodel

type User struct {
	UserID      string `db:"user_id"`
	Name        string `db:"name"`
	Email       string `db:"email"`
	PhoneNumber string `db:"phone_number"`
}
