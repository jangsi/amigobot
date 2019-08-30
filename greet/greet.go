package greet

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigobot"
)

//Handler handles the ?greet [name] command
type Handler struct{}

//Command is the trigger for the greet message
func (h *Handler) Command() string {
	return "?greet "
}

//Usage how the command works
func (h Handler) Usage() string {
	return "Greet whomever is specified. This command is mostly for example purposes for adding commands to amigo-bot rather than providing any real utility."
}

//Handle greets the person specified
func (h *Handler) Handle(s amigobot.Session, m *discordgo.MessageCreate) {
	toGreet := strings.TrimPrefix(m.Content, h.Command())
	s.ChannelMessageSend(m.ChannelID, "Ho there, "+toGreet+"!")
}
