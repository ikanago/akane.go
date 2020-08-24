package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
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

// Parse message from Discord
func ParseCommand(input string) (Command, error) {
	arguments := strings.Fields(input)
	if len(arguments) < 2 {
		return nil, errors.New("コマンドを指定してください!")
	}

	command := arguments[1]
	if command == "help" {
		return Help{}, nil
	} else if command == "ping" {
		return Ping{}, nil
	} else if command == "emoji" {
		if len(arguments) < 4 {
			return nil, errors.New("エイリアスまたは絵文字にするテキストを指定してください")
		}

		alias, err := validateAlias(arguments[2])
		if err != nil {
			return nil, err
		}

		text := arguments[3]
		var color string
		if len(arguments) >= 5 {
			color, err = validateColor(arguments[4])
			if err != nil {
				return nil, err
			}
		}

		var isTransparent bool
		if len(arguments) >= 6 {
			isTransparent, err = validateIsTransparent(arguments[5])
			if err != nil {
				return nil, err
			}
		}
		return EmojiFromText{
			Text:          text,
			Alias:         alias,
			Color:         color,
			IsTransparent: isTransparent,
		}, nil
	}
	return nil, errors.New("そのようなコマンドはありません><")
}

const aliasPattern = `^[a-zA-Z0-9_]+$`

var aliasRegexp = regexp.MustCompile(aliasPattern)

func validateAlias(alias string) (string, error) {
	if aliasRegexp.MatchString(alias) {
		return alias, nil
	}
	return "", errors.New("エイリアスには英数字とアンダーバーのみ使えます")
}

const colorCodePattern = `^[a-f0-9]+$`

var colorCodeRegexp = regexp.MustCompile(colorCodePattern)

func validateColor(color string) (string, error) {
	color = strings.ToLower(color)
	color = strings.TrimLeft(color, "#")
	if colorCode, isExists := colorTable[color]; isExists {
		// If color name is found at color code table
		color = colorCode
	}

	if !colorCodeRegexp.MatchString(color) {
		return "", errors.New("適切な色名またはカラーコードを入力してください")
	}

	if len(color) == 6 {
		return color, nil
	} else if len(color) == 3 {
		// For example, "2b4" -> "22bb44"
		var builder strings.Builder
		for i := 0; i < 3; i++ {
			fmt.Fprintf(&builder, "%c%c", color[i], color[i])
		}
		return builder.String(), nil
	}
	return "", errors.New("カラーコードは3桁または6桁で入力してください")
}

func validateIsTransparent(input string) (bool, error) {
	input = strings.ToLower(input)
	if input == "true" {
		return true, nil
	} else if input == "false" {
		return false, nil
	}
	return false, errors.New("TRANSPにはtrueまたはfalseを指定してください")
}
