package yn

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/ryanmiville/amigobot/amigobotfakes"
)

func TestGreet(t *testing.T) {
	h := Handler{}
	actual := &discordgo.Message{}
	s := &amigobotfakes.FakeSession{}
	s.ChannelMessageSendStub = func(channelID, content string) (*discordgo.Message, error) {
		actual.Content = content
		return actual, nil
	}
	s.MessageReactionAddStub = func(channelID, messageID, emojiID string) error {
		reaction := &discordgo.MessageReactions{
			Count: 1,
			Me:    true,
			Emoji: &discordgo.Emoji{ID: emojiID},
		}
		actual.Reactions = append(actual.Reactions, reaction)
		return nil
	}
	h.Handle(s, &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content: "?yn Does Cody smell good?",
		},
	})
	if actual.Content != "@everyone Does Cody smell good?" {
		t.Errorf("Expected Content: '@everyone Does Cody smell good?' but received %v", actual.Content)
	}
	if len(actual.Reactions) != 2 {
		t.Errorf("Expected len(actual.Reactions) == 2 but is actually %d", len(actual.Reactions))
	}
}
