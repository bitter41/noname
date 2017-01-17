package main

//import "fmt"
import "log"
import "time"
import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)
import (
	"./types"
	"fmt"
	"regexp"
	"github.com/noname/services"
)

const bot_token = "317035683:AAFvveEwHzO1wy-Rdqj6jCBjBI4msED2CLc"

func main() {
	//Activities := make(map[int]*types.Activity)

	bot, err := tgbotapi.NewBotAPI(bot_token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	activityService := services.ActivityService{
		Bot: bot,
	}

	for update := range updates {
		if update.CallbackQuery != nil {

			activityService.Update = update

			switch update.CallbackQuery.Data {
			case types.WORK:
				act := activityService.GetActivity(types.WORK)
				activityService.Start(act)
			case types.RELAX:
			case types.EAT:
			}
		}

		if update.Message == nil {
			continue
		}

		user := types.User{
			Id:       update.Message.From.ID,
			UserName: update.Message.From.FirstName,
		}
		chatId := update.Message.Chat.ID
		text := update.Message.Text

		log.Printf("[%s] %d %d %s", user.UserName, user.Id, chatId, text)



		command := regexp.MustCompile("/[a-z]+").FindString(update.Message.Text)

		switch command {
		case "/start":
			startButton := tgbotapi.NewInlineKeyboardButtonData(types.WORK, types.WORK)
			relaxButton := tgbotapi.NewInlineKeyboardButtonData(types.RELAX, types.RELAX)
			eatButton := tgbotapi.NewInlineKeyboardButtonData(types.EAT, types.EAT)

			inlineKeyboard := []tgbotapi.InlineKeyboardButton {
				startButton,
				relaxButton,
				eatButton,
			}
			inlineKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup(inlineKeyboard)

			response := tgbotapi.NewMessage(chatId, "Please select what activity you want to start.")
			response.BaseChat.ReplyMarkup = inlineKeyboardMarkup
			response.ParseMode = tgbotapi.ModeMarkdown

			bot.Send(response)
		case "/help":
			msg := tgbotapi.NewMessage(chatId, "To start activity print '/start'")
			bot.Send(msg)
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)


		var reply string
		msg := tgbotapi.NewMessage(chatId, "")
		if update.Message.NewChatMember != nil {
			reply = fmt.Sprintf("Hi @%s! Welcome to debug bot. Keep calm and eat a cookie (:",
				update.Message.NewChatMember.FirstName)
		}

		if reply != "" {
			msg.Text = reply
			bot.Send(msg)
		}

	}
}
