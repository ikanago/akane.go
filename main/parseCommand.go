package main

import (
	"errors"
	"strings"

	"github.com/asaskevich/govalidator"
)

// ParseCommand parses messages from Discord and returns results as sturct.
// Assumes first word of the input as a mention to this bot.
func ParseCommand(input string) (Command, error) {
	arguments := strings.Fields(input)
	if len(arguments) < 2 {
		return nil, errors.New("コマンドを指定してください")
	}

	command := arguments[1]
	if command == "help" {
		return Help{}, nil
	} else if command == "ping" {
		return Ping{}, nil
	} else if command == "goodjob" {
		return GoodJob{}, nil
	} else if command == "emoji" {
		if len(arguments) < 4 {
			return nil, errors.New("エイリアスまたは絵文字にするテキストを指定してください")
		}

		if arguments[2] == "image" {
			alias, err := validateAlias(arguments[3])
			if err != nil {
				return nil, err
			}
			return EmojiFromImage{Alias: alias}, nil
		}

		if arguments[2] == "url" {
			if len(arguments) < 5 {
				return nil, errors.New("画像のURLを指定してください")
			}

			alias, err := validateAlias(arguments[3])
			if err != nil {
				return nil, err
			}

			url := arguments[4]
			isURL := govalidator.IsURL(url)
			if !isURL {
				return nil, errors.New("URLの書式がおかしいです")
			}
			return EmojiFromURL{Alias: alias, URL: url}, nil
		}

		if arguments[2] == "delete" {
			alias, err := validateAlias(arguments[3])
			if err != nil {
				return nil, err
			}
			return EmojiDelete{Alias: alias}, nil
		}

		alias, err := validateAlias(arguments[2])
		if err != nil {
			return nil, err
		}

		text, err := processText(arguments[3])
		if err != nil {
			return nil, err
		}

		color := "000000"
		if len(arguments) >= 5 {
			color, err = validateColor(arguments[4])
			if err != nil {
				return nil, err
			}
		}

		transparency := "ff"
		if len(arguments) >= 6 {
			transparency, err = validateTransparency(arguments[5])
			if err != nil {
				return nil, err
			}
		}

		return EmojiFromText{
			Text:         text,
			Alias:        alias,
			Color:        color,
			Transparancy: transparency,
		}, nil
	}
	return nil, errors.New("そのようなコマンドはありません><")
}
