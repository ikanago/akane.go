package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	if len(message.Mentions) == 0 {
		return
	}

	if message.Mentions[0].Username == session.State.User.Username {
		arguments := ParseCommand(message.Content)
		log.Println(len(arguments))
	}
}
