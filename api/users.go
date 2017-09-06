package api

type UserRole struct {
	Name        string
	Description string
}

// User represents an authenticated Kontena user with basic details.
type User struct {
	ID    string
	Email string
	Name  string
	Roles []UserRole
}
