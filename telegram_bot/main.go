package main

import (
	"fmt"
	"log"
	utils "restaurant-app/telegram-bot/utils"
	"strings"

	tbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tbot.BotAPI

var menu string = `\# hi there`

// var htmlMenu string = `<h1>Menu</h1>`

//     <h3>Starters:</h3>
//     <ol>
//       <li>Hello Sir</li>
//     </ol>
//     <h3>Mains:</h3>
//     <ol>
//       <li>Yessir</li>
//     </ol>
//     <h3>Dessert:</h3>
//     <ol>
//       <li>123</li>
//     </ol>
// `

// https://gitlab.com/Athamaxy/telegram-bot-tutorial/-/blob/main/TutorialBot.go

func handleMessage(msg *tbot.Message) {
	user := msg.From
	text := msg.Text
	chatId := msg.Chat.ID

	if user == nil {
		log.Printf("User doesn't exist for message: %s", text)
		return
	}

	log.Printf("%s wrote %s", user, text)

	if strings.HasPrefix(text, "/") {
		split := strings.Split(text, " ")

		switch split[0] {
		case "/menu":
			tempMessage := tbot.NewMessage(chatId, menu)
			tempMessage.ParseMode = tbot.ModeMarkdownV2
			tempMessage.ReplyToMessageID = msg.MessageID
			_, err := bot.Send(tempMessage)
			if err != nil {
				log.Print("Failed to send msg to telegram API ", err, tempMessage)
				return
			}
			return
		case "/help":
			return
		default:
			tempMessage := tbot.NewMessage(chatId, fmt.Sprintf("'%s' is not a valid command.", split[0]))
			tempMessage.ReplyToMessageID = msg.MessageID
			_, err := bot.Send(tempMessage)
			if err != nil {
				log.Print("Failed to send msg to telegram API ", tempMessage)
				return
			}
			return
		}
	} else if strings.Trim(text, " ") == "help" {

	} else {

		tempMessage := tbot.NewMessage(chatId, fmt.Sprintf("'%s' is not a valid command.", text))
		tempMessage.ReplyToMessageID = msg.MessageID
		_, err := bot.Send(tempMessage)
		if err != nil {
			log.Print("Failed to send msg to telegram API ", tempMessage)
			return
		}
	}
}

func processUpdates(update tbot.Update) {
	switch {
	case update.Message != nil:
		handleMessage(update.Message)
	default:
		log.Println("Unhandled update", update)
	}
}

func receiveUpdates(updates tbot.UpdatesChannel) {
	for update := range updates {
		go processUpdates(update)
	}
}

func main() {
	var err error
	var config utils.Config
	config.GetEnvVariables()

	log.Println(config)

	bot, err = tbot.NewBotAPI(config.TelegramApiToken)
	if err != nil {
		log.Fatal("Failed to initialise the telegram bot")
	}

	u := tbot.NewUpdate(0)
	u.Timeout = config.Timeout

	updates := bot.GetUpdatesChan(u)

	receiveUpdates(updates)
}
