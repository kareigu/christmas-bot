package commands

import (
	"github.com/bwmarrin/discordgo"
)

var (
	List = []*discordgo.ApplicationCommand{
		{
			Name:        "christmas",
			Description: "Time until christmas",
		},
	}

	Handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"christmas": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			embed := ChristmasSlash()
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionApplicationCommandResponseData{
					Embeds: []*discordgo.MessageEmbed{
						&embed,
					},
				},
			})
		},
	}
)
