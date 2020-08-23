package main

import (
	"strings"
)

func ParseCommand(message string) []string {
	arguments := strings.Fields(message)
	return arguments
}
