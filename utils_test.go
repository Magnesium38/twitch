package twitch

import (
	"strings"
	"testing"
)

func TestSingleSplit(t *testing.T) {
	for _, c := range []struct {
		s, delim string
	}{
		{":tmi.twitch.tv 421 twitch_username WHO :Unknown command", " "},
		// Add more test cases as needed.
	} {
		a, b, err := SingleSplit(c.s, c.delim)
		if err != nil {
			t.Error(err)

			t.Log("case:", c.s)
			t.Log("delim:", c.delim)
			t.Log("a:", a)
			t.Log("b:", b)
			t.Log(strings.SplitN(c.s, c.delim, 2))
		}
	}
}
