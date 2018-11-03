package tweet

import "fmt"

// Tweet defines a single tweet by a user
type Tweet struct {
	User string
	// Message shoud not be more the 140 characters
	Message string
}

// NewTweet creates a new tweet message object
func NewTweet(user, msg string) *Tweet {
	return &Tweet{
		User:    user,
		Message: msg,
	}
}

func (t *Tweet) String() string {
	return fmt.Sprintf("@%s: %s", t.User, t.Message)
}
