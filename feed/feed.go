package feed

import (
	"codeassignment/tf/tweet"
	"codeassignment/tf/user"
	"fmt"
	"io"
	"sort"
)

// Feed combines and presents all of the current tweets
type Feed struct {
	userStore  user.Store
	tweetStore tweet.Store
}

// NewTweetFeed creates a new instance of the tweet feed given the relevant stores
func NewTweetFeed(userStore user.Store, tweetStore tweet.Store) *Feed {
	return &Feed{
		userStore:  userStore,
		tweetStore: tweetStore,
	}
}

// Print will output the current tweet feed for each user (sorted alphabetically
func (f *Feed) Print(w io.Writer) error {
	users, err := f.userStore.All()
	if err != nil {
		return fmt.Errorf("could not print feed: %s", err)
	}

	tweets, err := f.tweetStore.All()
	if err != nil {
		return fmt.Errorf("could not print feed: %s", err)
	}

	// sorts the users alphabetically
	sort.Slice(users, func(i, j int) bool {
		return users[i].Identifier < users[j].Identifier
	})

	for _, user := range users {
		fmt.Fprintln(w, user.Identifier)

		for _, tweet := range tweets {
			if user.IsFollowing(tweet.User) {
				fmt.Fprintf(w, "\t%s\n", tweet)
			}
		}
	}
	return nil
}
