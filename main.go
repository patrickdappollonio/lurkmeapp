package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gempir/go-twitch-irc"
	"github.com/patrickdappollonio/readfile"
)

func main() {
	// Get username and token from the environment variables
	username, token, err := getLoginInfo()
	if err != nil {
		log.Fatalf("Error getting log-in info: %s", err.Error())
	}

	log.Println("Conneting to Twitch with username:", username)

	// Get the location of the channels file or
	// set a default location if none is set
	channelsFileLocation := env("CHANNELS_FILE", "channels.txt")

	// Get all the channels from the parsed file
	channels, err := readfile.New(channelsFileLocation).Parse()
	if err != nil {
		log.Fatalf("Error reading channels file at %s: %s", channelsFileLocation, err.Error())
	}

	// Create a client to connect to the Twitch IRC server
	client := twitch.NewClient(username, token)
	client.Join(channels...)

	// On connect callback, print a message with the successful connection
	client.OnConnect(func() {
		log.Println("Connected to Twitch server, lurking on", len(channels), "channels")
	})

	// Handler for when a new message is received
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		log.Printf("%s on #%s: %s", message.User.DisplayName, message.Channel, message.Message)
	})

	// Connect to the Twitch server
	if err := client.Connect(); err != nil {
		log.Fatalf("Unable to connect to the Twitch IRC server: %s", err.Error())
	}

	notifyExit := make(chan bool, 1)

	// Close the connection once we're closing the program
	go func() {
		notifyClose := make(chan os.Signal, 1)
		signal.Notify(notifyClose, os.Interrupt, syscall.SIGTERM)
		<-notifyClose

		log.Println("Closing connection to Twitch... Goodbye!")
		client.Disconnect()
		notifyExit <- true
	}()

	<-notifyExit
}
