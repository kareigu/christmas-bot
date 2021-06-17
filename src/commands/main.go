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
		"christmas": ChristmasSlash,
	}
)
