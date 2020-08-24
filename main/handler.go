package main

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

type Command interface {
	handle(*discordgo.Session, *discordgo.Message) error
}

type Help struct{}

type EmojiFromText struct{}

type ParseError struct {
	content string
}

func (Help) handle(session *discordgo.Session, message *discordgo.Message) (err error) {
	messageEmbed := discordgo.MessageEmbed{
		Color:  0xF9A9BF,
		Type:   discordgo.EmbedTypeRich,
		Title:  "アカネチャンのコマンド",
		Fields: HelpMessageEmbeds,
	}
	_, err = session.ChannelMessageSendEmbed(message.ChannelID, &messageEmbed)
	return
}

func (EmojiFromText) handle(session *discordgo.Session, message *discordgo.Message) (err error) {
	err = errors.New("Unimplemented")
	return
}

func (parseError ParseError) handle(session *discordgo.Session, message *discordgo.Message) (err error) {
	reply := parseError.content
	_, err = session.ChannelMessageSend(message.ChannelID, reply)
	return
}
