package twitch

import "strings"

// A ChatMessage is the internal representation of a message used by
// Twitch.tv's IRC based chat service.
type ChatMessage struct {
	tags      []Tag
	source    string
	command   string
	arguments string
}

// Build takes a message and formats it as an actual message.
func (m *ChatMessage) Build() (msg string) {
	msg = ""
	if len(m.tags) > 0 {

		tags := "@"
		for _, tag := range m.tags {
			tags += tag.key + "=" + tag.value + ";"
		}
		msg += tags[:len(tags)-1] + " "
	}
	msg += ":" + m.source
	msg += " " + m.command
	msg += " " + m.arguments
	return
}

// InvalidChatMessageError is a custom error for failures at parsing a message.
type InvalidChatMessageError struct {
	msg string
}

func (e InvalidChatMessageError) Error() string {
	return e.msg
}

func singleSplit(s string, delim string) (a, b string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = InvalidChatMessageError{
				"Invalid Chat Message; Error splitting \"" +
					s + "\" with \"" + delim + "\";",
			}
		}
	}()
	split := strings.SplitN(s, delim, 2)
	a, b, err = split[0], split[1], nil
	return
}

// ParseChatMessage takes a string and parses it into a ChatMessage
func ParseChatMessage(msg string) (chatMsg ChatMessage, err error) {
	chatMsg = ChatMessage{}
	err = nil

	// Parse the space-delimited portions.
	isFinished := false
	var currentPart string
	for {

		currentPart, msg, err = singleSplit(msg, " ")
		if err != nil {
			return
		}

		// Parse tag material, if any.

		switch string(currentPart[0]) {
		case "@":
			// Parse tag material, if any.
			var tags []Tag
			var tag, key, value string
			currentPart = currentPart[1:]
			for len(currentPart) > 0 {
				if strings.Index(currentPart, ";") != -1 {
					tag, currentPart, err = singleSplit(currentPart, ";")
					if err != nil {
						return
					}
				} else {
					tag, currentPart = currentPart, ""
				}

				key, value, err = singleSplit(tag, "=")
				if err != nil {
					return
				}

				tags = append(tags, Tag{key, value})
			}

			chatMsg.tags = tags
		case ":":
			// Parse source.
			chatMsg.source = currentPart[1:]
			isFinished = true
		}

		if isFinished {
			break
		}
	}

	// Next is the command.
	currentPart, msg, err = singleSplit(msg, " ")
	chatMsg.command = currentPart

	chatMsg.arguments = msg
	return
}
