package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const aliasPattern = `^[a-zA-Z0-9_]+$`

var aliasRegexp = regexp.MustCompile(aliasPattern)

func validateAlias(alias string) (string, error) {
	if aliasRegexp.MatchString(alias) {
		return alias, nil
	}
	return "", errors.New("エイリアスには英数字とアンダーバーのみ使えます")
}

// If there are more than or equal 4 UTF-8 characters, insert newline.
func processText(text string) (string, error) {
	utf8Len := len([]rune(text))
	if utf8Len > 10 {
		return "", errors.New("テキストが長すぎます")
	}
	if utf8Len < 4 {
		return text, nil
	}

	// 0-indexed position to which '\n' is inserted into after
	// Ceiling `utf8Len / 2`.
	indice := (utf8Len + 1) / 2
	inserted := string([]rune(text)[:indice]) + "\n" + string([]rune(text)[indice:])
	return inserted, nil
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

// Validate input and return transparency value corresponding to user input.
func validateTransparency(input string) (string, error) {
	input = strings.ToLower(input)
	if input == "true" {
		return "00", nil
	} else if input == "false" {
		return "ff", nil
	}
	return "", errors.New("TRANSPにはtrueまたはfalseを指定してください")
}
