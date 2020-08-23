package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/bwmarrin/discordgo"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

type Credential struct {
	DiscordToken string `json:"DISCORD_TOKEN"`
}

func accessSecretVersion(projectID string, secretID string) (*Credential, error) {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	secretURI := "projects/" + projectID + "/secrets/" + secretID + "/versions/latest"
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: secretURI,
	}

	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return nil, err
	}

	var credential Credential
	if err := json.Unmarshal(result.Payload.Data, &credential); err != nil {
		return nil, err
	}
	return &credential, nil
}

func main() {
	projectID := os.Getenv("PROJECT_ID")
	secretID := os.Getenv("SECRET_ID")
	credential, err := accessSecretVersion(projectID, secretID)
	if err != nil {
		log.Fatal(err)
		return
	}

	client, err := discordgo.New("Bot " + credential.DiscordToken)
	if err != nil {
		log.Fatal(err)
		return
	}
	client.AddHandler(MessageCreate)

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
