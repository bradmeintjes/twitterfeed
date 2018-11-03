package main

import (
	"codeassignment/tf/feed"
	"codeassignment/tf/tweet"
	"codeassignment/tf/user"
	"log"
	"os"
)

const (
	usage = "feed <users_file> <tweet_file>"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal(usage)
	}

	userStore := user.NewFileStore(args[0])
	tweetStore := tweet.NewFileStore(args[1])

	feed := feed.NewTweetFeed(userStore, tweetStore)
	if err := feed.Print(os.Stderr); err != nil {
		log.Fatal(err)
	}
}
