package feed

import (
	"bytes"
	"codeassignment/tf/tweet"
	"codeassignment/tf/user"
	"codeassignment/tf/util"
	"testing"
)

const (
	expectedPrint = `user1
	@user1: tweet 1 message
	@user2: tweet 2 message
user2
	@user1: tweet 1 message
	@user2: tweet 2 message
`
)

func TestPrintFeed(t *testing.T) {
	user1 := user.NewUser("user1")
	user2 := user.NewUser("user2")
	user2.Follow(user1)
	user1.Follow(user2)
	mockUsers := []*user.User{user1, user2}

	tweet1 := tweet.NewTweet(user1.Identifier, "tweet 1 message")
	tweet2 := tweet.NewTweet(user2.Identifier, "tweet 2 message")
	mockTweets := []*tweet.Tweet{tweet1, tweet2}

	feed := NewTweetFeed(&mockUserStore{mockUsers}, &mockTweetStore{mockTweets})

	var b bytes.Buffer
	err := feed.Print(&b)
	util.AssertNoErr(t, err, "feed should print successfully")
	util.CheckEq(t, expectedPrint, b.String(), "feed print should match")
}

type mockUserStore struct {
	users []*user.User
}

func (m *mockUserStore) All() ([]*user.User, error) {
	return m.users, nil
}

type mockTweetStore struct {
	tweets []*tweet.Tweet
}

func (m *mockTweetStore) All() ([]*tweet.Tweet, error) {
	return m.tweets, nil
}
