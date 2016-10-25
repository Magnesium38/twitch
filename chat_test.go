package twitch

import (
	"bufio"
	"os"
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
		a, b, err := singleSplit(c.s, c.delim)
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

func TestParseChatMessage(t *testing.T) {
	testCases, err := os.Open("test_parse_messages.txt")
	if err != nil {
		t.Error(err)
	}
	defer testCases.Close()

	scanner := bufio.NewScanner(testCases)
	for scanner.Scan() {
		c := scanner.Text()
		msg, err := ParseChatMessage(c)
		if err != nil {
			t.Error(err)
		}

		output := msg.Build()
		if c != output {
			t.Error("Build does not match input.\nInput:  \"" +
				c + "\";\nOutput: \"" + output + "\"")
		}
	}
}
