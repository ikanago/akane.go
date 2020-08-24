package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseCommand(t *testing.T) {
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

	t.Run("emoji command", func(t *testing.T) {
		input := "<@!746704561451827202> emoji"
		expected := EmojiFromText{}
		actual, err := ParseCommand(input)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.Nil(err)
	})

	t.Run("emoji command", func(t *testing.T) {
		input := "<@!746704561451827202> hoge"
		actual, err := ParseCommand(input)
		assert := assert.New(t)
		assert.Nil(actual)
		assert.NotNil(err)
	})
}
