package main

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID || len(message.Mentions) == 0 || message.Mentions[0].Username != session.State.User.Username {
		return
	}

	arguments := strings.Fields(message.Content)
	if len(arguments) < 2 {
		return
	}

	log.Print(arguments)
	command := arguments[1]
	if command == "help" {
		helpMessage := getHelpMessage()
		if _, err := session.ChannelMessageSendEmbed(message.ChannelID, helpMessage); err != nil {
			log.Fatal(err)
			return
		}
	} else if command == "emoji" {
		reply, err := CreateEmojiFromText()
		if err != nil {
			log.Fatal(err)
			return
		}

		log.Print(reply)
		if _, err := session.ChannelMessageSend(message.ChannelID, reply); err != nil {
			log.Fatal(err)
			return
		}
	}
}
