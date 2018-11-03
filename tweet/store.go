package tweet

// Store defines the tweet persistence interface
type Store interface {
	All() ([]*Tweet, error)
}
