package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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
		current_time := time.Now()
		year := current_time.Year()
		christmas_time, _ := time.Parse(Layout, fmt.Sprintf("%d-Dec-25", year))
		if current_time.After(christmas_time) {
			year += 1
		}

		christmas_time, _ = time.Parse(Layout, fmt.Sprintf("%d-Dec-24", year))

		duration := christmas_time.Sub(current_time)

		days := int64(duration.Hours()) / 24
		hours := int64(duration.Hours()) - days*24
		time_until := fmt.Sprintf("Time until christmas:\n %d days, %d hours", days, hours)

		image := discordgo.MessageEmbedImage{
			URL: ChristmasImg,
		}

		embed := discordgo.MessageEmbed{
			Title:       "Homopapat",
			Description: time_until,
			Type:        discordgo.EmbedTypeImage,
			Image:       &image,
		}

		new_msg := discordgo.MessageSend{
			Embed: &embed,
		}
		_, err := session.ChannelMessageSendComplex(msg.ChannelID, &new_msg)
		if err != nil {
			log.Println("Error sending message" + err.Error())
		}
	}
}
