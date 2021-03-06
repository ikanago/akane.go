package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateAlias(t *testing.T) {
	t.Run("Valid alias", func(t *testing.T) {
		alias := "abc_123_1a2b3c"
		expected := "abc_123_1a2b3c"
		actual, err := validateAlias(alias)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.Nil(err)
	})

	t.Run("Invalid alias", func(t *testing.T) {
		alias := "abc!"
		expected := ""
		actual, err := validateAlias(alias)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.NotNil(err)
	})
}

func TestProcessText(t *testing.T) {
	t.Run("No folding", func(t *testing.T) {
		text := "破天荒"
		expected := "破天荒"
		actual, err := processText(text)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.Nil(err)
	})

	t.Run("Folding", func(t *testing.T) {
		text := "意気揚々"
		expected := "意気\n揚々"
		actual, err := processText(text)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.Nil(err)
	})

	t.Run("Too long", func(t *testing.T) {
		text := "寿限無寿限無五劫の擦り切れ"
		expected := ""
		actual, err := processText(text)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.NotNil(err)
	})
}

func TestValidateColor(t *testing.T) {
	t.Run("Valid color code(6 hex)", func(t *testing.T) {
		color := "#12FaBc"
		expected := "12fabc"
		actual, err := validateColor(color)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.Nil(err)
	})

	t.Run("Valid color code without #", func(t *testing.T) {
		color := "##12FaBc"
		expected := "12fabc"
		actual, err := validateColor(color)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.Nil(err)
	})

	t.Run("Valid color code(3 hex)", func(t *testing.T) {
		color := "1aB"
		expected := "11aabb"
		actual, err := validateColor(color)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.Nil(err)
	})

	t.Run("Valid color name", func(t *testing.T) {
		color := "siLveR"
		expected := "c0c0c0"
		actual, err := validateColor(color)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.Nil(err)
	})

	t.Run("Invalid color code", func(t *testing.T) {
		color := "#12gabc"
		expected := ""
		actual, err := validateAlias(color)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.NotNil(err)
	})

	t.Run("Invalid color code length", func(t *testing.T) {
		color := "#1234567"
		expected := ""
		actual, err := validateAlias(color)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.NotNil(err)
	})

	t.Run("Invalid color name", func(t *testing.T) {
		color := "siLveRr"
		expected := ""
		actual, err := validateColor(color)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.NotNil(err)
	})
}

func TestValidateIsTransparent(t *testing.T) {
	t.Run("Valid transparency flag", func(t *testing.T) {
		flag := "tRuE"
		expected := "00"
		actual, err := validateTransparency(flag)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.Nil(err)
	})

	t.Run("Valid transparency flag", func(t *testing.T) {
		flag := "FalSE"
		expected := "ff"
		actual, err := validateTransparency(flag)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.Nil(err)
	})

	t.Run("Invalid transparency flag", func(t *testing.T) {
		flag := "truee"
		expected := ""
		actual, err := validateTransparency(flag)
		assert := assert.New(t)
		assert.Equal(expected, actual)
		assert.NotNil(err)
	})
}
