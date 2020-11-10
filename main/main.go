package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := os.Getenv("DISCORD_TOKEN")
	client, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
		return
	}
	client.AddHandler(OnMessageCreate)

	err = client.Open()
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Bot is running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	client.Close()
}

// OnMessageCreate is called when there is a new message in a guild this bot is belogns to.
// If this bot is mentioned, parse command and do corresponding actions.
func OnMessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID || len(message.Mentions) == 0 || message.Mentions[0].Username != session.State.User.Username {
		return
	}

	command, err := ParseCommand(message.Content)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, err.Error())
		log.Println(err)
		return
	}

	if err = command.handle(session, message.Message); err != nil {
		session.ChannelMessageSend(message.ChannelID, err.Error())
		log.Println(err)
	}
}
