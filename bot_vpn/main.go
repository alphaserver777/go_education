package main

import (
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Администраторский чат-ид (замените на реальный ID чата с администратором)
const adminChatID = 698226393

// Структура для хранения информации о подписке
type Subscription struct {
	UserID   int64
	Username string
	Plan     string
	EndDate  time.Time
}

var subscriptions = make(map[int64]Subscription)

func main() {
	bot, err := tgbotapi.NewBotAPI("7294995368:AAEbyzYTjhaq3_Mvph-6SctYe1w7GwSLT4Q")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	updConfig := tgbotapi.NewUpdate(0)
	updConfig.Timeout = 60

	updates := bot.GetUpdatesChan(updConfig)

	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				handleCommand(bot, update.Message)
				continue
			}
			if update.Message.Photo != nil {
				handleScreenshot(bot, update.Message)
				continue
			}
		}

		if update.CallbackQuery != nil {
			handleApproval(bot, update.CallbackQuery)
		}
	}
}

func handleCommand(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	switch msg.Command() {
	case "start":
		welcomeMessage := "Добро пожаловать! Выберите подписку:\n1 месяц - /subscribe_1\n3 месяца - /subscribe_3\n6 месяцев - /subscribe_6"
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, welcomeMessage))
	case "subscribe_1":
		requestPayment(bot, msg, "1 месяц")
	case "subscribe_3":
		requestPayment(bot, msg, "3 месяца")
	case "subscribe_6":
		requestPayment(bot, msg, "6 месяцев")
	default:
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Неизвестная команда."))
	}
}

func requestPayment(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, plan string) {
	message := fmt.Sprintf("Вы выбрали подписку на %s.\n\nПожалуйста, переведите оплату на следующие реквизиты: \n*Карта:* 1234 5678 9012 3456\n\nИ отправьте скриншот об оплате.", plan)
	bot.Send(tgbotapi.NewMessage(msg.Chat.ID, message))
	subscriptions[msg.Chat.ID] = Subscription{
		UserID:   msg.Chat.ID,
		Username: msg.From.UserName,
		Plan:     plan,
	}
}

func handleScreenshot(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	subscription, exists := subscriptions[msg.Chat.ID]
	if !exists {
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Вы еще не выбрали подписку."))
		return
	}

	caption := fmt.Sprintf("Пользователь: @%s\nПлан: %s", subscription.Username, subscription.Plan)
	photo := tgbotapi.NewPhoto(adminChatID, tgbotapi.FileID(msg.Photo[len(msg.Photo)-1].FileID))
	photo.Caption = caption

	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Согласовать", fmt.Sprintf("approve_%d", msg.Chat.ID)),
			tgbotapi.NewInlineKeyboardButtonData("Отклонить", fmt.Sprintf("decline_%d", msg.Chat.ID)),
		),
	)
	photo.ReplyMarkup = buttons

	_, err := bot.Send(photo)
	if err != nil {
		log.Printf("Ошибка отправки фото администратору: %v", err)
	}
}

func handleApproval(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	action := callback.Data
	var response string

	if action[:7] == "approve" {
		userID := parseUserID(action)
		subscription := subscriptions[userID]
		response = fmt.Sprintf("Подписка на %s успешно активирована!", subscription.Plan)

		duration := map[string]time.Duration{
			"1 месяц":   30 * 24 * time.Hour,
			"3 месяца":  90 * 24 * time.Hour,
			"6 месяцев": 180 * 24 * time.Hour,
		}
		endDate := time.Now().Add(duration[subscription.Plan])
		subscription.EndDate = endDate
		subscriptions[userID] = subscription

		keyMessage := fmt.Sprintf("Ваша подписка активирована!\nКлюч: %s\nДата окончания: %s", generateKey(), endDate.Format("02.01.2006"))
		bot.Send(tgbotapi.NewMessage(userID, keyMessage))
	} else if action[:7] == "decline" {
		userID := parseUserID(action)
		response = "Подписка отклонена."
		bot.Send(tgbotapi.NewMessage(userID, "Ваш запрос на подписку был отклонен."))
	}

	bot.Send(tgbotapi.NewMessage(adminChatID, response))

	// Отправляем подтверждение обработки кнопки
	callbackResponse := tgbotapi.NewCallback(callback.ID, "Действие выполнено.")
	bot.Request(callbackResponse)
}

func parseUserID(action string) int64 {
	var userID int64
	fmt.Sscanf(action, "approve_%d", &userID)
	return userID
}

func generateKey() string {
	return "KEY-1234-5678" // Здесь можно использовать генератор уникальных ключей
}
