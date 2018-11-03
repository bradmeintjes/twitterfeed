package main

import (
	"allangray/tf/feed"
	"allangray/tf/tweet"
	"allangray/tf/user"
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
	feed.Print(os.Stdout)
}
