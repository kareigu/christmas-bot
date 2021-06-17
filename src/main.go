package main

import (
	"cmd_handler"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	handler cmd_handler.CommandHandler
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
	prefix := os.Getenv("PREFIX")
	christmasImg := os.Getenv("CHRISTMAS_IMG")

	if bot_token == "" || prefix == "" || christmasImg == "" {
		log.Fatal("Problem with .env file, values missing")
	}

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

	handler = cmd_handler.CommandHandler{
		Session:      client,
		Prefix:       prefix,
		Christmasimg: christmasImg,
	}

	log.Println("Running...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	client.Close()
}

func onMessageHandler(_ *discordgo.Session, msg *discordgo.MessageCreate) {
	cmd_handler.RunCmd(&handler, msg)
}
