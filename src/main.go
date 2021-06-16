package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"commands"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	Prefix       string
	ChristmasImg string
)

const (
	Layout = "2006-Jan-02"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't load env file")
	}

	bot_token := os.Getenv("TOKEN")
	Prefix = os.Getenv("PREFIX")
	ChristmasImg = os.Getenv("CHRISTMAS_IMG")

	client, err := discordgo.New("Bot " + bot_token)
	if err != nil {
		log.Fatal("Couldn't create bot client")
	}

	err = client.Open()
	if err != nil {
		log.Fatal("Couldn't start bot client")
	}

	client.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuilds

	client.AddHandler(onMessageHandler)

	log.Println("Running...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	client.Close()
}

func onMessageHandler(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == session.State.User.ID {
		return
	}

	if msg.Content == Prefix+"christmas" {
		new_msg := commands.ChristmasCmd(Layout, ChristmasImg)
		_, err := session.ChannelMessageSendComplex(msg.ChannelID, &new_msg)
		if err != nil {
			log.Println("Error sending message" + err.Error())
		}
	}
}
