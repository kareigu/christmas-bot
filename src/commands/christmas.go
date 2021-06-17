package commands

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func christmasCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	const layout = "2006-Jan-02"

	current_time := time.Now()
	year := current_time.Year()
	christmas_time, _ := time.Parse(layout, fmt.Sprintf("%d-Dec-25", year))
	if current_time.After(christmas_time) {
		year += 1
	}

	christmas_time, _ = time.Parse(layout, fmt.Sprintf("%d-Dec-24", year))

	duration := christmas_time.Sub(current_time)

	log.Println(duration)

	days := int64(duration.Hours()) / 24
	hours := int64(duration.Hours()) - days*24
	minutes := int64(duration.Minutes()) - days*24*60 - hours*60
	seconds := int64(duration.Seconds()) - days*24*60*60 - hours*60*60 - minutes*60
	milliseconds := int64(duration.Milliseconds()) - days*24*60*60*1000 - hours*60*60*1000 - minutes*60*1000 - seconds*1000
	microseconds := int64(duration.Microseconds()) - days*24*60*60*1000*1000 - hours*60*60*1000*1000 - minutes*60*1000*1000 - seconds*1000*1000 - milliseconds*1000
	nanoseconds := int64(duration.Nanoseconds()) - days*24*60*60*1000*1000*1000 - hours*60*60*1000*1000*1000 - minutes*60*1000*1000*1000 - seconds*1000*1000*1000 - milliseconds*1000*1000 - microseconds*1000

	time_until := fmt.Sprintf(`Time until christmas: 
		%d days, %d hours
		%d minutes %d seconds
		%d milliseconds
		%d microseconds
		%d nanoseconds`,
		days, hours,
		minutes, seconds,
		milliseconds,
		microseconds,
		nanoseconds)

	image := discordgo.MessageEmbedImage{
		URL: "https://mxrr.dev/files/lmao%20-%20Copy.gif",
	}

	embed := discordgo.MessageEmbed{
		Title:       "Homopapat",
		Description: time_until,
		Type:        discordgo.EmbedTypeImage,
		Image:       &image,
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Embeds: []*discordgo.MessageEmbed{
				&embed,
			},
		},
	})
}
