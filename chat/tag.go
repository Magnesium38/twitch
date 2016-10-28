package chat

// A Tag is a key value pair as given by the Twitch API.
type Tag struct {
	key, value string
}

// Key is the getter for the key.
func (tag *Tag) Key() string {
	return tag.key
}

// Value is the getter for the value.
func (tag *Tag) Value() string {
	return tag.value
}

func (tag *Tag) String() string {
	return tag.key + "=" + tag.value
}

// Tags is a wrapper around a slice of Tag.
type Tags struct {
	tags []Tag
}

// Length returns the number of tags in the internal slice.
func (tags *Tags) Length() int {
	return len(tags.tags)
}

// Add takes a tag and adds it to the internal slice.
func (tags *Tags) Add(tag Tag) {
	tags.tags = append(tags.tags, tag)
}

// Has takes the key and returns the Tag if it exists.
func (tags *Tags) Has(tagName string) (bool, Tag) {
	for _, tag := range tags.tags {
		if tag.key == tagName {
			return true, tag
		}
	}
	return false, Tag{}
}

func (tags *Tags) String() string {
	str := "@"
	for _, tag := range tags.tags {
		str += tag.String() + ";"
	}

	return str[:len(str)-1]
}
