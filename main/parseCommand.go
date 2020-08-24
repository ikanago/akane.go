package main

import (
	"errors"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID || len(message.Mentions) == 0 || message.Mentions[0].Username != session.State.User.Username {
		return
	}

	command, err := ParseCommand(message.Content)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, err.Error())
		log.Println(err)
		return
	}

	if err = command.handle(session, message.Message); err != nil {
		session.ChannelMessageSend(message.ChannelID, err.Error())
		log.Println(err)
	}
}

func ParseCommand(input string) (Command, error) {
	arguments := strings.Fields(input)
	if len(arguments) < 2 {
		return nil, errors.New("コマンドを指定してください!")
	}

	command := arguments[1]
	if command == "help" {
		return Help{}, nil
	} else if command == "emoji" {
		return EmojiFromText{}, nil
	}
	return nil, errors.New("そのようなコマンドはありません><")
}
