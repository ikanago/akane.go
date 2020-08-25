package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// EmojifyText access API to get image of emoji and return base64-encoded image.
func (emojiFromText *EmojiFromText) EmojifyText() (encodedImage string, err error) {
	encodedText := url.QueryEscape(emojiFromText.Text)
	apiURL := fmt.Sprintf("https://emoji-gen.ninja/emoji_download?align=center&back_color=FFFFFF%s&color=%sFF&font=notosans-mono-bold&public_fg=false&size_fixed=false&stretch=true&text=%s", emojiFromText.Transparancy, emojiFromText.Color, encodedText)

	response, err := http.Get(apiURL)
	if err != nil {
		log.Println(err)
		return "", errors.New("絵文字の作成に失敗しました")
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return "", errors.New("絵文字の作成に失敗しました")
	}
	encodedImage = base64.StdEncoding.EncodeToString(body)
	encodedImage = fmt.Sprintf("data:png;base64,%s", encodedImage)
	return
}
