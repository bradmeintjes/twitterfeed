package user

// Store defines the user persistence interface
type Store interface {
	All() ([]*User, error)
}
