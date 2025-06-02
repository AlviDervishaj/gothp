package db

// Example model: User represents a user in the database.
type User struct {
	ID    int
	Name  string
	Email string
}

// You can add methods for CRUD operations here, e.g.:
// func GetUser(id int) (*User, error) { ... }
