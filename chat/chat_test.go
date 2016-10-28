package chat

import (
	"bufio"
	"os"
	"testing"
)

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
