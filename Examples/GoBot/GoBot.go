package main

import (
	"log"
	"os"

	tg "gopkg.in/telegram-bot-api.v4"
)

func main() {

	// Use @BotFather to generate a token
	bot, err := tg.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))

	// Check any errors during bot authentication
	if err != nil {
		log.Fatal(err)
	}

	// Staring up message
	log.Printf("Starting %s", bot.Self.UserName)

	// Create a channel for the updates
	updates, err := bot.GetUpdatesChan(tg.NewUpdate(0))

	// Check errors during channel creation
	if err != nil {
		log.Fatal(err)
	}

	// Listen to the channel updates
	for update := range updates {
		// If the update is a message
		if update.Message != nil {
			// Log the message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			// Send a hello message to the user
			bot.Send(tg.NewMessage(update.Message.Chat.ID, "Hola "+update.Message.From.FirstName))
		}
	}
}
