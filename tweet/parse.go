package tweet

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	maxTweetLen = 140
)

var (
	ErrMalformedTweet = errors.New("malformed tweet line")
)

// FileStore implments the Store interface using a flat file for persistence
type FileStore struct {
	file  string
	cache []*Tweet
}

// NewFileStore creates a new file store using the given file to source the data
func NewFileStore(file string) *FileStore {
	return &FileStore{
		file: file,
	}
}

// All fetches all of the tweets from the storage file, parsing the file only once,
// and returning a cached value for all future instantiations
func (fs *FileStore) All() ([]*Tweet, error) {
	if fs.cache == nil {
		tweets, err := parseTweetsFromFile(fs.file)
		if err != nil {
			return nil, fmt.Errorf("could not fetch all tweets: %s", err)
		}
		fs.cache = tweets

	}
	return fs.cache, nil

}

// Parse tweets from a flat file
func parseTweetsFromFile(file string) ([]*Tweet, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("could not open tweets file: %s", err)
	}

	tweets := make([]*Tweet, 0, 1024)
	lineNr := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		userID, tweet, err := parseLine(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse tweet file %s:%d: %s", file, lineNr, err)
		}

		tweets = append(tweets, NewTweet(userID, tweet))
	}

	return tweets, nil
}

// parses a tweet from a single line of text
// line format expected: <username>> <message(up to 140 chars)>
func parseLine(line string) (user, tweet string, err error) {
	sepIdx := strings.Index(line, "> ")
	if sepIdx == -1 || sepIdx+2 >= len(line) {
		err = ErrMalformedTweet
		return
	}

	user = strings.TrimSpace(line[:sepIdx])
	tweet = strings.TrimSpace(line[sepIdx+2:])
	// tweet max length is 140 charaters
	if len(tweet) > maxTweetLen {
		tweet = tweet[:maxTweetLen]
	}

	if len(user) == 0 || len(tweet) == 0 {
		err = ErrMalformedTweet
	}
	return
}
