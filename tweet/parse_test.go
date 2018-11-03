package tweet

import (
	"codeassignment/tf/util"
	"testing"
)

func TestParseFromFile(t *testing.T) {
	store := NewFileStore("../res/tweet.txt")
	util.AssertNotNil(t, store, "should create file store")

	tweets, err := store.All()
	util.AssertNoErr(t, err, "should parse tweet file")
	util.AssertNotNil(t, tweets, "should parse tweets from file")
	util.CheckEq(t, 3, len(tweets), "should read correct number of tweets")

	checkTweetEq(t, tweets[0], "Alan", "If you have a procedure with 10 parameters, you probably missed some.")
	checkTweetEq(t, tweets[1], "Ward", "There are only two hard things in Computer Science: cache invalidation, naming things and off-by-1 errors.")
	checkTweetEq(t, tweets[2], "Alan", "Random numbers should not be generated with a method chosen at random.")
}

func TestParseLine(t *testing.T) {
	// empty line should not parse correctly
	_, _, err := parseLine("")
	if err == nil || err != ErrMalformedTweet {
		t.Logf("expected error %s from empty input, but got %s", ErrMalformedTweet, err)
		t.Fail()
	}

	// empty tweet should not parse
	_, _, err = parseLine("Bob> ")
	if err == nil || err != ErrMalformedTweet {
		t.Logf("expected error %s from malformed input, but got %s", ErrMalformedTweet, err)
		t.Fail()
	}

	// invalid separator
	_, _, err = parseLine("Bob< hello")
	if err == nil || err != ErrMalformedTweet {
		t.Logf("expected error %s from malformed input, but got %s", ErrMalformedTweet, err)
		t.Fail()
	}

	user, message, err := parseLine("Bob> the builder ")
	util.AssertNoErr(t, err, "")
	util.CheckEq(t, "Bob", user, "should parse user correctly")
	util.CheckEq(t, "the builder", message, "should parse tweet message correctly")

	// should cut off messages > 140 chars
	const longMsg = "This is just a really, really long tweet which should ultimatley get cut off after one hundred and forty (140) characters, damn not there yet, here is just a bit more nonesense ..."
	user, message, err = parseLine("Bob> " + longMsg)
	util.AssertNoErr(t, err, "")
	util.CheckEq(t, "Bob", user, "should parse use140r correctly")
	util.CheckEq(t, longMsg[:maxTweetLen], message, "should cut off length tweet > 140 chars")
}

func checkTweetEq(t *testing.T, tweet *Tweet, userID, message string) {
	util.CheckEq(t, userID, tweet.User, "should match name")
	util.CheckEq(t, message, tweet.Message, "should match message")
}
