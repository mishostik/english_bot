package usecase

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type UpdateUsecase struct {
	UpdateConfig tgbotapi.UpdateConfig
	bot          *tgbotapi.BotAPI
}

func Init(bot *tgbotapi.BotAPI) *UpdateUsecase {

	updateCfg := tgbotapi.NewUpdate(0)
	updateCfg.Timeout = 5

	return &UpdateUsecase{
		UpdateConfig: updateCfg,
		bot:          bot,
	}
}

func (u *UpdateUsecase) NewUpdates() (*tgbotapi.UpdatesChannel, error) {
	updates, err := u.bot.GetUpdatesChan(u.UpdateConfig)
	if err != nil {
		return nil, err
	}
	if len(updates) == 0 {
		return nil, fmt.Errorf("no updates")
	}
	log.Println(fmt.Sprintf("len updated = {%d}", len(updates)))
	log.Println(fmt.Sprintf("updates {%+v}", updates))
	return &updates, nil
}
