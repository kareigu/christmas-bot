package main

import (
	"cmd_handler"
	"commands"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	handler cmd_handler.CommandHandler
	client  *discordgo.Session
	GuildID string
)

const (
	Layout = "2006-Jan-02"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't load env file")
	}

	bot_token := os.Getenv("TOKEN")
	prefix := os.Getenv("PREFIX")
	christmasImg := os.Getenv("CHRISTMAS_IMG")
	GuildID = os.Getenv("GUILD_ID")

	if bot_token == "" || prefix == "" || christmasImg == "" {
		log.Fatal("Problem with .env file, values missing")
	}

	client, err = discordgo.New("Bot " + bot_token)
	if err != nil {
		log.Fatal("Couldn't create bot client")
	}

	handler = cmd_handler.CommandHandler{
		Session:      client,
		Prefix:       prefix,
		Christmasimg: christmasImg,
	}
}

func init() {
	client.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commands.Handlers[i.Data.Name]; ok {
			h(s, i)
		}
	})
}

func main() {
	client.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuilds

	client.AddHandler(onMessageHandler)

	err := client.Open()
	if err != nil {
		log.Fatal("Couldn't start bot client")
	}

	for _, v := range commands.List {
		_, err = client.ApplicationCommandCreate(client.State.User.ID, GuildID, v)
		if err != nil {
			log.Panicf("Couldn't create command %v : %v", v.Name, err)
		}
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
