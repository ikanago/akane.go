package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

type Command interface {
	handle(*discordgo.Session, *discordgo.Message) error
}

type Help struct{}

type Ping struct {}

type EmojiFromText struct {
	Text          string
	Alias         string
	Color         string
	IsTransparent bool
}

func (Help) handle(session *discordgo.Session, message *discordgo.Message) (err error) {
	messageEmbed := discordgo.MessageEmbed{
		Color:  0xF9A9BF,
		Type:   discordgo.EmbedTypeRich,
		Title:  "アカネチャンのコマンド",
		Fields: HelpMessageEmbeds,
	}
	result, err := session.ChannelMessageSendEmbed(message.ChannelID, &messageEmbed)
	log.Println(result)
	return
}

func (Ping) handle(session *discordgo.Session, message *discordgo.Message) (err error) {
	reply := "Pong!"
	result, err := session.ChannelMessageSend(message.ChannelID, reply)
	log.Println(result)
	return
}

func (emojiFromText EmojiFromText) handle(session *discordgo.Session, message *discordgo.Message) (err error) {
	reply := fmt.Sprint(emojiFromText)
	result, err := session.ChannelMessageSend(message.ChannelID, reply)
	log.Println(result)
	return
}
