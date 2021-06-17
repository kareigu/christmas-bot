package cmd_handler

import (
	"commands"

	"github.com/bwmarrin/discordgo"
)

type CommandHandler struct {
	Prefix       string
	Christmasimg string
	Session      *discordgo.Session
}

func RunCmd(cmd_handler *CommandHandler, msg *discordgo.MessageCreate) {
	if msg.Author.ID == cmd_handler.Session.State.User.ID {
		return
	}

	switch msg.Content {
	case cmd_handler.Prefix + "christmas":
		commands.ChristmasCmd(cmd_handler.Session, msg, cmd_handler.Christmasimg)
	}
}
