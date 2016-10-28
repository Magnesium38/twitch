package command

import (
	"strconv"

	"github.com/magnesium38/twitch"
)

// ClearChatMessage is a struct for messages with the command CLEARCHAT.
type ClearChatMessage struct {
	twitch.ChatMessage
}

// BanReason returns the given ban reason for the message.
func (msg *ClearChatMessage) BanReason() string {
	if has, tag := msg.Tags().Has("ban-duration"); has {
		return tag.Value()
	}
	return ""
}

// BanDuration returns the ban duration for the message. -1 is permanent.
func (msg *ClearChatMessage) BanDuration() int {
	if has, tag := msg.Tags().Has("ban-duration"); has {
		value, err := strconv.Atoi(tag.Value())
		if err == nil {
			return value
		}
	}
	return -1
}
