package services

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/noname/adapter"
	"github.com/noname/types"
	"time"
)

type ActivityService struct {
	Update tgbotapi.Update
	Bot    *tgbotapi.BotAPI
	ActivityDAO *adapter.ActivityDAO
}

func NewActivityService(update tgbotapi.Update, bot *tgbotapi.BotAPI, activityDAO *adapter.ActivityDAO) *ActivityService {
	return &ActivityService{update, bot, activityDAO}
}

func (a *ActivityService) Start(activity types.Activity) (bool) {
	//duration := activity.ActivityConfig.Duration
	chunk := activity.ActivityConfig.ChunkSize

	if chunk >= 0 {
		activity.StartDateTime = time.Now()
		activity.Launched = true
		ticker := time.NewTicker(time.Second * chunk)
		go func() {
			for range ticker.C {
				//send a proposal to change the activity
				a.Bot.Send(tgbotapi.NewMessage(
					a.Update.CallbackQuery.Message.Chat.ID, "tick tack"))
			}
		}()
	}
	a.ActivityDAO.Start(activity)
	return activity.Launched
}

func (a *ActivityService) Stop(activity types.Activity) (bool) {
	activity.Launched = false
	activity.StopDateTime = time.Now()
	return activity.Launched
}

//Return activity with default sitting
func (a *ActivityService) GetActivity(activityType string) (types.Activity) {
	var config types.ActivityConfig

	activity := types.Activity{
		ActivityType:   activityType,
		User: types.User{
			Id:       a.Update.CallbackQuery.From.ID,
			UserName: a.Update.CallbackQuery.From.FirstName,
		},
	}

	switch activityType {
	case types.WORK:
		config.Duration = 28800
		config.ChunkSize = 5
		break
	case types.RELAX:
		config.Duration = 420
		break
	case types.EAT:
		config.Duration = 3600
		break
	}

	activity.ActivityConfig = config
	return activity
}
