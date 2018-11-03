package user

import (
	"allangray/tf/util"
	"sort"
	"testing"
)

func TestParseFromFile(t *testing.T) {
	store := NewFileStore("../res/user.txt")
	util.AssertNotNil(t, store, "should create file store")

	users, err := store.All()
	util.AssertNoErr(t, err, "should parse user file")
	util.AssertNotNil(t, users, "should parse users from file")
	util.CheckEq(t, 3, len(users), "should parse users correctly")

	// sort users by name for consistency
	sort.Slice(users, func(i, j int) bool {
		return users[i].Name < users[j].Name
	})

	checkUserEq(t, users[0], "Alan", "Martin")
	checkUserEq(t, users[1], "Martin")
	checkUserEq(t, users[2], "Ward", "Martin", "Alan")
}

func TestParseLine(t *testing.T) {
	_, _, err := parseLine("")
	if err == nil || err != ErrMalformedUser {
		t.Logf("expected error %s from empty input, but got %s", ErrMalformedUser, err)
		t.Fail()
	}

	_, _, err = parseLine("Bob follow Alice,")
	if err == nil || err != ErrMalformedUser {
		t.Logf("expected error %s from malformed input, but got %s", ErrMalformedUser, err)
		t.Fail()
	}

	user, follows, err := parseLine("Bob follows Alice,Jimmy,Bacon")
	util.AssertNoErr(t, err, "")
	util.CheckEq(t, "Bob", user, "should parse user correctly")
	util.CheckEq(t, 3, len(follows), "should parse follows correctly")
	if len(follows) == 3 {
		util.CheckEq(t, "Alice", follows[0], "should parse follow correctly")
		util.CheckEq(t, "Jimmy", follows[1], "should parse follow correctly")
		util.CheckEq(t, "Bacon", follows[2], "should parse follow correctly")
	}

	user, follows, err = parseLine("Bob follows ")
	util.AssertNoErr(t, err, "should parse valid line")
	util.CheckEq(t, user, "Bob", "should parse user")
	util.CheckEq(t, 0, len(follows), "should allow no follows")
}

func checkUserEq(t *testing.T, user *User, name string, follows ...string) {
	util.CheckEq(t, name, user.Name, "should match name")
	util.CheckEq(t, len(follows), len(user.Follows.ToSlice()), "should have correct number of follows")

	for _, follow := range follows {
		found := false

		it := user.Follows.Iterator()
		for v := range it.C {
			if v.(*User).Name == follow {
				found = true
			}
		}

		if !found {
			t.Logf("unmatched follower for user %s: %s", user.Name, follow)
			t.Fail()
			return
		}
	}
}
