package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

// Command is an interface of parsed commands.
type Command interface {
	handle(*discordgo.Session, *discordgo.Message) error
}

// Help represents parsed results of `help` command.
type Help struct{}

// Ping represents parsed results of `ping` command.
type Ping struct{}

// EmojiFromText represents parsed results of `emoji ALIAS TEXT ...` command.
type EmojiFromText struct {
	Text         string
	Alias        string
	Color        string
	Transparancy string
}

func (Help) handle(session *discordgo.Session, message *discordgo.Message) (err error) {
	messageEmbed := discordgo.MessageEmbed{
		Color:  0xF9A9BF,
		Type:   discordgo.EmbedTypeRich,
		Title:  "アカネチャンのコマンド",
		Fields: helpMessageEmbeds,
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
	encodedImage, err := emojiFromText.EmojifyText()
	if err != nil {
		return
	}

	emoji, err := session.GuildEmojiCreate(message.GuildID, emojiFromText.Alias, encodedImage, nil)
	log.Printf("Emoji: %v", emoji)
	if err != nil {
		log.Println(err)
		return errors.New("カスタム絵文字の追加に失敗しました")
	}

	reply := fmt.Sprintf("カスタム絵文字 :%s: を追加しました", emojiFromText.Alias)
	result, err := session.ChannelMessageSend(message.ChannelID, reply)
	log.Println(result)
	return
}
