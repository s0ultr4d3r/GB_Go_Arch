package notification

import (
	"GB/lesson-2/shop/models"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Notification interface {
	SendOrderCreated(order *models.Order) error
}

type telegramBot struct {
	chatID int64
	tgBot  *tgbotapi.BotAPI
}

func (s *telegramBot) SendOrderCreated(order *models.Order) error {

	text := fmt.Sprintf("new order %d\nemail: %s\nphone: %s", order.ID, order.Email, order.CustomerPhone)

	msg := tgbotapi.NewMessage(s.chatID, text)

	_, err := s.tgBot.Send(msg)
	return err
}

func NewTelegramBot(token string, chatID int64) (Notification, error) {
	tgBot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &telegramBot{
		chatID: chatID,
		tgBot:  tgBot,
	}, nil
}
