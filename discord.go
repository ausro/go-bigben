package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func sendMessage(time int64, s *discordgo.Session) {
	st := fmt.Sprintf("It is now <t:%d:t>", time)

	word := getRandomWord()

	field := discordgo.MessageEmbedField{
		Name:   word.Name,
		Value:  word.Definition,
		Inline: false,
	}

	embed := discordgo.MessageEmbed{
		Title:       "Bing Bong~~",
		Description: st,
		Color:       1,
		Fields:      []*discordgo.MessageEmbedField{&field},
	}

	send := discordgo.MessageSend{
		Flags:  discordgo.MessageFlags(discordgo.MessageFlagsSuppressNotifications),
		Embeds: []*discordgo.MessageEmbed{&embed},
	}

	s.ChannelMessageSendComplex(Configuration.ChannelId, &send)

	fmt.Printf("Sent random word embed to %s\n", Configuration.ChannelId)
}

func initDiscord(c chan int64) {
	discord, err := discordgo.New("Bot " + Configuration.Token)
	if err != nil {
		fmt.Printf("Failed to create bot: %s\n", err)
		return
	}

	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Bot is ready")
	})

	opErr := discord.Open()
	if opErr != nil {
		log.Fatalf("Cannot open session: %s", err)
	}

	go func() {
		for {
			m := <-c
			sendMessage(m, discord)
		}
	}()
}
