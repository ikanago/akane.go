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

type EmojiFromImage struct {
	Alias string
}

type EmojiFromURL struct {
	Alias string
	URL   string
}

type EmojiDelete struct {
	Alias string
}

func (Help) handle(session *discordgo.Session, message *discordgo.Message) (err error) {
	messageEmbed := discordgo.MessageEmbed{
		Color:  0xF9A9BF,
		Type:   discordgo.EmbedTypeRich,
		Title:  "アカネチャンのコマンド",
		Fields: helpMessageEmbeds,
	}
	_, err = session.ChannelMessageSendEmbed(message.ChannelID, &messageEmbed)
	return
}

func (Ping) handle(session *discordgo.Session, message *discordgo.Message) (err error) {
	reply := "Pong!"
	_, err = session.ChannelMessageSend(message.ChannelID, reply)
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
	_, err = session.ChannelMessageSend(message.ChannelID, reply)
	return
}

func (emojiFromImage EmojiFromImage) handle(session *discordgo.Session, message *discordgo.Message) (err error) {
	if len(message.Attachments) != 1 {
		return errors.New("指定できる画像は1つです")
	}

	encodedImage, err := getImageFromURL(message.Attachments[0].URL)
	if err != nil {
		return
	}

	emoji, err := session.GuildEmojiCreate(message.GuildID, emojiFromImage.Alias, encodedImage, nil)
	log.Printf("Emoji: %v", emoji)
	if err != nil {
		log.Println(err)
		return errors.New("カスタム絵文字の追加に失敗しました")
	}

	reply := fmt.Sprintf("カスタム絵文字 :%s: を追加しました", emojiFromImage.Alias)
	_, err = session.ChannelMessageSend(message.ChannelID, reply)
	return
}

func (emojiFromURL EmojiFromURL) handle(session *discordgo.Session, message *discordgo.Message) (err error) {
	encodedImage, err := getImageFromURL(emojiFromURL.URL)
	if err != nil {
		return
	}

	emoji, err := session.GuildEmojiCreate(message.GuildID, emojiFromURL.Alias, encodedImage, nil)
	log.Printf("Emoji: %v", emoji)
	if err != nil {
		log.Println(err)
		return errors.New("カスタム絵文字の追加に失敗しました")
	}

	reply := fmt.Sprintf("カスタム絵文字 :%s: を追加しました", emojiFromURL.Alias)
	_, err = session.ChannelMessageSend(message.ChannelID, reply)
	return
}

func (emojiDelete EmojiDelete) handle(session *discordgo.Session, message *discordgo.Message) (err error) {
	emojiID, err := fetchEmojiID(session, message.GuildID, emojiDelete.Alias)
	if err != nil {
		return
	}

	err = session.GuildEmojiDelete(message.GuildID, emojiID)
	if err != nil {
		return
	}

	reply := fmt.Sprintf("カスタム絵文字 :%s: を削除しました", emojiDelete.Alias)
	_, err = session.ChannelMessageSend(message.ChannelID, reply)
	return
}

// Search emoji whose alias is `alias` in a certain server.
func fetchEmojiID(session *discordgo.Session, guildID string, alias string) (emojiID string, err error) {
	emojis, err := session.GuildEmojis(guildID)
	if err != nil {
		log.Println(err)
		return "", errors.New("絵文字を取得できませんでした")
	}

	for _, emoji := range emojis {
		if emoji.Name == alias {
			emojiID = emoji.ID
			return
		}
	}
	return "", errors.New("そのようなエイリアスの絵文字はありません")
}
