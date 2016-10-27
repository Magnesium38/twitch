package twitch

import "strings"

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
