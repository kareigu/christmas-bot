package cmd_handler

import (
	"commands"

	"github.com/bwmarrin/discordgo"
)

type CommandHandler struct {
	prefix       string
	christmasimg string
	layout       string
	session      *discordgo.Session
}

func New(
	prefix string,
	christmasimg string,
	layout string,
	session *discordgo.Session) *CommandHandler {
	cmd_handler := CommandHandler{
		prefix:       prefix,
		christmasimg: christmasimg,
		layout:       layout,
		session:      session,
	}

	return &cmd_handler
}

func RunCmd(cmd_handler *CommandHandler, msg *discordgo.MessageCreate) {
	if msg.Author.ID == cmd_handler.session.State.User.ID {
		return
	}

	switch msg.Content {
	case cmd_handler.prefix + "christmas":
		commands.ChristmasCmd(cmd_handler.session, msg, cmd_handler.layout, cmd_handler.christmasimg)
	}
}
