package twitch

// IrcMessage is a basic representation of an IRC message.
type IrcMessage interface {
	Source() string
	Command() string
	Arguments() string
	Tags() *Tags
	Build() string
}
