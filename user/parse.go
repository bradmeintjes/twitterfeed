package user

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	ErrMalformedUser = errors.New("malformed user line")
)

// FileStore implments the Store interface using a fflat file for persistence
type FileStore struct {
	file  string
	cache []*User
}

// NewFileStore creates a new file store using the given file to source the data
func NewFileStore(file string) *FileStore {
	return &FileStore{
		file: file,
	}
}

// All fetches all of the users from the storage file, parsing the file only once,
// and returning a cached value for all future instantiations
func (fs *FileStore) All() ([]*User, error) {
	if fs.cache == nil {
		users, err := parseUsersFromFile(fs.file)
		if err != nil {
			return nil, fmt.Errorf("could not fetch all users: %s", err)
		}

		fs.cache = make([]*User, 0, len(users))
		for _, user := range users {
			fs.cache = append(fs.cache, user)
		}
	}
	return fs.cache, nil

}

// Parse user graph from given file, by mapping the user identity (user name in this case)
// to a set of user identities which the user follows
func parseUsersFromFile(file string) (map[string]*User, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("could not open users file: %s", err)
	}

	users := make(map[string]*User)
	lineNr := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		u := make([]*User, 0, len(users))
		for _, user := range users {
			u = append(u, user)
		}

		userID, follows, err := parseLine(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse users file %s:%d: %s", file, lineNr, err)
		}

		// ensure that all the followed users exist in the map
		for _, follow := range follows {
			if len(follow) != 0 {
				if _, exists := users[follow]; !exists {
					users[follow] = NewUser(follow)
				}
			}
		}

		user, exists := users[userID]
		if !exists {
			user = NewUser(userID)
		}

		// append all followed users to the user object (union)
		for _, id := range follows {
			user.Follows.Add(users[id])
		}
		users[userID] = user
	}

	return users, nil
}

// parses a user identifier and the user identifiers of the users which the user follows
// line format expected: <username> follows <user1>[, <user2>]
func parseLine(line string) (user string, follows []string, err error) {
	const separator = " follows "

	idx := strings.Index(line, separator)
	if idx == -1 {
		err = ErrMalformedUser
		return
	}

	// it is possible to not follow anybody, in which case we still need to record the users existence
	user = strings.TrimSpace(line[0:idx])

	followIdx := idx + len(separator)
	if followIdx >= len(line) {
		return
	}

	rawFollows := strings.Split(line[followIdx:], ",")
	for _, follow := range rawFollows {
		sanitized := strings.TrimSpace(follow)
		if len(sanitized) > 0 {
			follows = append(follows, sanitized)
		}
	}

	return
}
