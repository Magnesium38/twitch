package chat

import (
	"strings"

	"github.com/magnesium38/twitch"
)

// InvalidChatMessageError is a custom error for failures at parsing a message.
type InvalidChatMessageError struct {
	msg string
}

func (err InvalidChatMessageError) Error() string {
	return err.msg
}

// A Message is the internal representation of a message used by
// Twitch.tv's IRC based chat service.
type Message struct {
	tags      Tags
	source    string
	command   string
	arguments string
}

// Source returns whoever sent the message.
func (m *Message) Source() string {
	return m.source
}

// Command returns the IRC command associated with the message.
func (m *Message) Command() string {
	return m.command
}

// Arguments returns the arguments for the command.
func (m *Message) Arguments() string {
	return m.arguments
}

// Tags returns the tags that are associated with the message.
func (m *Message) Tags() *Tags {
	return &m.tags
}

// Build takes a message and formats it as an actual message.
func (m *Message) Build() (msg string) {
	msg = ""
	if m.tags.Length() > 0 {
		msg += m.Tags().String() + " "
	}
	msg += ":" + m.source
	msg += " " + m.command
	msg += " " + m.arguments
	return
}

// ParseChatMessage takes a string and parses it into a ChatMessage
func ParseChatMessage(msg string) (chatMsg Message, err error) {
	chatMsg = Message{}
	err = nil

	// Parse the space-delimited portions.
	isFinished := false
	var currentPart string
	for {

		currentPart, msg, err = twitch.SingleSplit(msg, " ")
		if err != nil {
			return
		}

		// Parse tag material, if any.

		switch string(currentPart[0]) {
		case "@":
			// Parse tag material, if any.
			var tags Tags
			var tag, key, value string
			currentPart = currentPart[1:]
			for len(currentPart) > 0 {
				if strings.Index(currentPart, ";") != -1 {
					tag, currentPart, err = twitch.SingleSplit(currentPart, ";")
					if err != nil {
						return
					}
				} else {
					tag, currentPart = currentPart, ""
				}

				key, value, err = twitch.SingleSplit(tag, "=")
				if err != nil {
					return
				}

				tags.Add(Tag{key, value})
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
	currentPart, msg, err = twitch.SingleSplit(msg, " ")
	chatMsg.command = currentPart

	chatMsg.arguments = msg
	return
}
