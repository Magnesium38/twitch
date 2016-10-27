package twitch

import "strconv"

// Roomstate is a set of tags that define the set of the room.
type Roomstate struct {
	language string
	r9k      bool
	subOnly  bool
	slow     int
}

// Update takes a message and updates the Roomstate appropriately if needed.
func (state *Roomstate) Update(msg ChatMessage) {
	if msg.command == "ROOMSTATE" {
		if has, tag := msg.tags.Has("broadcaster-lang"); has {
			state.SetLanguage(tag.value)
		}
		if has, tag := msg.tags.Has("r9k"); has {
			value, err := strconv.Atoi(tag.value)
			if err == nil {
				if value == 1 {
					state.r9k = true
				} else {

					state.r9k = false
				}
			}
		}
		if has, tag := msg.tags.Has("subs-only"); has {
			value, err := strconv.Atoi(tag.value)
			if err == nil {
				if value == 1 {
					state.subOnly = true
				} else {
					state.subOnly = false
				}
			}
		}
		if has, tag := msg.tags.Has("slow"); has {
			value, err := strconv.Atoi(tag.value)
			if err == nil {
				state.slow = value
			}
		}
	}
}

// SetLanguage is a simple setter for Roomstate.
func (state *Roomstate) SetLanguage(language string) {
	state.language = language
}

// Language is the getter for the Roomstate's language.
func (state *Roomstate) Language() string {
	return state.language
}

// SetR9K is a simple setter for Roomstate.
func (state *Roomstate) SetR9K(r9k bool) {
	state.r9k = r9k
}

// R9K is the getter for the Roomstate's R9K status.
func (state *Roomstate) R9K() bool {
	return state.r9k
}

// SetSubOnly is a simple setter for Roomstate.
func (state *Roomstate) SetSubOnly(subOnly bool) {
	state.subOnly = subOnly
}

// SubOnly is the getter for the Roomstate's sub only status.
func (state *Roomstate) SubOnly() bool {
	return state.subOnly
}

// SetSlow is a simple setter for Roomstate.
func (state *Roomstate) SetSlow(slow int) {
	state.slow = slow
}

// Slow is the getter for the Roomstate's slow mode length.
func (state *Roomstate) Slow() int {
	return state.slow
}
