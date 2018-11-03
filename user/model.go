package user

// User defines user properties, and the users
type User struct {
	Identifier string

	// Follows is a hash map based set
	Follows map[string]*User
}

// NewUser creates a newuser object with no follows
func NewUser(id string) *User {
	return &User{
		Identifier: id,
		Follows:    make(map[string]*User),
	}
}

// IsFollowing return true if the user is following a user with the given identifier, or if the identifier is that of the user
func (u *User) IsFollowing(id string) bool {
	_, exists := u.Follows[id]
	return exists || id == u.Identifier
}

// Follow will add the given user to th eset of user which the current user is following
func (u *User) Follow(usr *User) {
	u.Follows[usr.Identifier] = usr
}

func (u *User) String() string {
	return u.Identifier
}
