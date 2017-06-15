package api

type UserRole struct {
	Name        string
	Description string
}

type User struct {
	ID    string
	Email string
	Name  string
	Roles []UserRole
}
