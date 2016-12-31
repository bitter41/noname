package main

//import "fmt"
import "log"
import "time"
import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)


const bot_token  = "317035683:AAFvveEwHzO1wy-Rdqj6jCBjBI4msED2CLc"

type Activity struct {
	StartDateTime 	time.Time
	Type 		string
	Usermane	string
}

type User struct {
	UserID 		int
	Username 	string
}

func (u User) isDbSynced() int {

	return 1
}

func (u User) updateInDB() int {

	return 1
}

func main() {
	Activities := make( map[int]*Activity )

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

		if _, ok := Activities[ update.Message.From.ID ]; ok {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "work already started by " + strconv.Itoa( update.Message.From.ID) )
			bot.Send(msg)
		} else {
			Activities[ update.Message.From.ID ] = &Activity{time.Now(), "work", update.Message.From.UserName }
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "work start for " + strconv.Itoa( update.Message.From.ID))
			bot.Send(msg)
		}

		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//msg.ReplyToMessageID = update.Message.MessageID


	}
}





