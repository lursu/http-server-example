package data

import (
	"database/sql"
)

type Users interface {
	UserByID(id string) (*User, error)
	CreateUser(name string) (string, error)
	UpdateUser(id, name string) (*User, error)
}

type User struct {
	ID        string
	Name      string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

// UserByID gets a user from the provided id
func (c *client) UserByID(id string) (*User, error) {
	stmt := `SELECT id, name, created_at, updated_at FROM users
		WHERE id = $1;`

	user := new(User)
	err := c.db.QueryRow(stmt, id).Scan(&user.ID, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

// CreateUser creates a user for the provided values
func (c *client) CreateUser(name string) (string, error) {
	stmt := `INSERT INTO users (name) 
		VALUES ($1)
		RETURNING id;`

	var id string
	err := c.db.QueryRow(stmt, name).Scan(&id)
	return id, err
}

// UpdateUser updates the user with the given id and returns the resulting changes
func (c *client) UpdateUser(id, name string) (*User, error) {
	stmt := `UPDATE users 
		SET name = $2, updated_at = NOW() 
		WHERE id = $1 
		RETURNING name, created_at, updated_at;`

	user := &User{
		ID: id,
	}
	err := c.db.QueryRow(stmt, id, name).Scan(&user.Name, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}
