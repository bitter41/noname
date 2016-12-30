package main

//import "fmt"
import "log"
import "time"
import "github.com/go-telegram-bot-api/telegram-bot-api"

const bot_token  = "317035683:AAFvveEwHzO1wy-Rdqj6jCBjBI4msED2CLc"

func main() {
	bot, err := tgbotapi.NewBotAPI( bot_token )
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
