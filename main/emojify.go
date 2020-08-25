package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// EmojifyText access API to get image of emoji and return base64-encoded image.
func (emojiFromText *EmojiFromText) EmojifyText() (encodedImage string, err error) {
	encodedText := url.QueryEscape(emojiFromText.Text)
	apiURL := fmt.Sprintf("https://emoji-gen.ninja/emoji_download?align=center&back_color=FFFFFF%s&color=%sFF&font=notosans-mono-bold&public_fg=false&size_fixed=false&stretch=true&text=%s", emojiFromText.Transparancy, emojiFromText.Color, encodedText)
	encodedImage, err = getImageFromURL(apiURL)
	return
}

// Get an image by accessing given URL and return the image encoded in base64.
func getImageFromURL(url string) (encodedImage string, err error) {
	response, err := http.Get(url)
	if err != nil {
		return "", errors.New("画像の作成に失敗しました")
	}
	defer response.Body.Close()

	contentType := response.Header.Get("Content-Type")
	if !strings.Contains(contentType, "image/") {
		return "", errors.New("画像へのURLを指定してください")
	}

	imageByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", errors.New("画像の読み込みに失敗しました")
	}

	// Size of image to convert into emoji must be smaller than 256kB.
	maximumSize := 262144
	if len(imageByte) > maximumSize {
		return "", errors.New("画像のサイズは256kB以下にしてください")
	}

	encodedImage = base64.StdEncoding.EncodeToString(imageByte)
	encodedImage = fmt.Sprintf("data:png;base64,%s", encodedImage)
	return
}
