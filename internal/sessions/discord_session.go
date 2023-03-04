package sessions

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/edy4c7/go-discord-bot/internal/bots"
)

func NewDiscordSession(token string, bot bots.Bot, listening <-chan any) error {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return err
	}

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		isMention := false
		for _, v := range m.Mentions {
			if v.ID == s.State.User.ID {
				isMention = true
				break
			}
		}

		ch, err := dg.Channel(m.ChannelID)
		if err != nil {
			return
		}

		if ch.Type == discordgo.ChannelTypeDM || isMention {
			// DMまたはメンションで送信されたメッセージにのみ反応
			ctx := context.Background()
			reply, err := bot.Dialogue(ctx, m.Content)
			if err != nil {
				return
			}
			_, err = s.ChannelMessageSendReply(m.ChannelID, reply, m.Reference())
			if err != nil {
				return	
			}
		}
	})

	err = dg.Open()
	if err != nil {
		return err
	}
	defer dg.Close()

	<-listening

	return nil
}
