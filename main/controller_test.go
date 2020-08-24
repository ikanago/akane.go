package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateEmojiFromText(t *testing.T) {
	expected := "emoji"
	actual, err := CreateEmojiFromText()
	assert := assert.New(t)
	assert.Equal(expected, actual)
	assert.Nil(err)
}
