package tgclient

import (
	"fmt"
	"log"
	"os"

	"openlog/olclient"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// config
const UPDATE_TIMEOUT = 30

// commands strings
const HELP = "help"
const GREET = "greet"
const LAST_ERROR = "lasterr"

var commands = map[string]string{
	HELP:       "show this message",
	GREET:      "say hi",
	LAST_ERROR: "give the last error by date on openlog database",
}

func Run() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_KEY"))
	if err != nil {
		log.Panic(err)
	}
	respondToUpdates(bot)
}

func respondToUpdates(bot *tgbotapi.BotAPI) {
	updateConfig := getUpdateConfig()
	for update := range bot.GetUpdatesChan(updateConfig) {
		if update.Message.IsCommand() {
			respondToCommands(bot, update)
		}
	}
}

func getUpdateConfig() tgbotapi.UpdateConfig {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = UPDATE_TIMEOUT
	return updateConfig
}

func respondToCommands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = createCommandResponse(update.Message.Command())

	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func createCommandResponse(command string) string {
	switch command {
	case HELP:
		helpResponse := "Available commands are:\n"
		for command, description := range commands {
			helpResponse += fmt.Sprintf("/%s - %s\n", command, description)
		}
		return helpResponse
	case GREET:
		return "Greetings"
	case LAST_ERROR:
		return olclient.GetLastError()
	default:
		return fmt.Sprintf("I don't know that command, use /%s for list commands", HELP)
	}
}
