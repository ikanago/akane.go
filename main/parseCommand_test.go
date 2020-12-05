package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseBasicCommands(t *testing.T) {
	t.Run("Command not specified", func(t *testing.T) {
		input := "<@!746704561451827202>"
		actual, err := ParseCommand(input)
		assert := assert.New(t)
		assert.Nil(actual)
		assert.NotNil(err)
	})

	t.Run("help command", func(t *testing.T) {
		input := "<@!746704561451827202> help"
		expected := Help{}
		actual, err := ParseCommand(input)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.Nil(err)
	})

	t.Run("ping command", func(t *testing.T) {
		input := "<@!746704561451827202> ping"
		expected := Ping{}
		actual, err := ParseCommand(input)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.Nil(err)
	})

	t.Run("Goodjob command", func(t *testing.T) {
		input := "<@!746704561451827202> goodjob"
		expected := GoodJob{}
		actual, err := ParseCommand(input)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.Nil(err)
	})

	t.Run("Unknown command", func(t *testing.T) {
		input := "<@!746704561451827202> hoge"
		actual, err := ParseCommand(input)
		assert := assert.New(t)
		assert.Nil(actual)
		assert.NotNil(err)
	})
}

func TestParseEmojiFromText(t *testing.T) {
	t.Run("emoji from text without specifying color and transparency", func(t *testing.T) {
		input := "<@!746704561451827202> emoji    \nhoge123_456   あいうabc"
		expected := EmojiFromText{
			Alias:        "hoge123_456",
			Text:         "あいう\nabc",
			Color:        "000000",
			Transparancy: "ff",
		}
		actual, err := ParseCommand(input)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.Nil(err)
	})

	t.Run("emoji from text without specifying transparency", func(t *testing.T) {
		input := "<@!746704561451827202> emoji hoge123_456 あいうabc ##12Fa4C"
		expected := EmojiFromText{
			Alias:        "hoge123_456",
			Text:         "あいう\nabc",
			Color:        "12fa4c",
			Transparancy: "ff",
		}
		actual, err := ParseCommand(input)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.Nil(err)
	})

	t.Run("emoji from text, transparency", func(t *testing.T) {
		input := "<@!746704561451827202> emoji hoge123_456 あいうabc #12Fa4C false"
		expected := EmojiFromText{
			Alias:        "hoge123_456",
			Text:         "あいう\nabc",
			Color:        "12fa4c",
			Transparancy: "ff",
		}
		actual, err := ParseCommand(input)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.Nil(err)
	})

	t.Run("emoji from text, transparency", func(t *testing.T) {
		input := "<@!746704561451827202> emoji hoge123_456 あいうabc #12Fa4C true"
		expected := EmojiFromText{
			Alias:        "hoge123_456",
			Text:         "あいう\nabc",
			Color:        "12fa4c",
			Transparancy: "00",
		}
		actual, err := ParseCommand(input)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.Nil(err)
	})
}

func TestParseEmojiFromImage(t *testing.T) {
	t.Run("emoji from image", func(t *testing.T) {
		input := "<@!746704561451827202> emoji image hoge123_456 "
		expected := EmojiFromImage{
			Alias: "hoge123_456",
		}
		actual, err := ParseCommand(input)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.Nil(err)
	})
}

func TestParseEmojiFromURL(t *testing.T) {
	t.Run("emoji from image", func(t *testing.T) {
		input := "<@!746704561451827202> emoji url hoge123_456 http://example.com"
		expected := EmojiFromURL{
			Alias: "hoge123_456",
			URL:   "http://example.com",
		}
		actual, err := ParseCommand(input)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.Nil(err)
	})

	t.Run("emoji from image with invalid URL", func(t *testing.T) {
		input := "<@!746704561451827202> emoji url hoge123_456 http:/example.com"
		actual, err := ParseCommand(input)
		assert := assert.New(t)
		assert.Nil(actual)
		assert.NotNil(err)
	})
}

func TestParseEmojiDelete(t *testing.T) {
	t.Run("emoji from image", func(t *testing.T) {
		input := "<@!746704561451827202> emoji delete hoge123_456"
		expected := EmojiDelete{
			Alias: "hoge123_456",
		}
		actual, err := ParseCommand(input)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.Nil(err)
	})
}
