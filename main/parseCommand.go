package main

import (
	"log"
	"strings"
)

func ParseCommand(message string) []string {
	arguments := strings.Fields(message)
	log.Println(len(arguments))
	return arguments
}
