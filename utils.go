package twitch

import "strings"

// SingleSplitError is a custom error for failures when splitting once.
type SingleSplitError struct {
	msg string
}

func (err SingleSplitError) Error() string {
	return err.msg
}

// SingleSplit is a helper method to only do a single split.
func SingleSplit(s string, delim string) (a, b string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = SingleSplitError{
				"Error splitting \"" + s + "\" with \"" + delim + "\";",
			}
		}
	}()
	split := strings.SplitN(s, delim, 2)
	a, b, err = split[0], split[1], nil
	return
}
