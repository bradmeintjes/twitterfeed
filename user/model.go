package user

import (
	mapset "github.com/deckarep/golang-set"
)

// User defines user properties, and the users
type User struct {
	Name string

	// Follows is a hash map based set
	Follows mapset.Set
}

// NewUser creates a newuser object with no follows
func NewUser(name string) *User {
	return &User{
		Name:    name,
		Follows: mapset.NewSet(),
	}
}
