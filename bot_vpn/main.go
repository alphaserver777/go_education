package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
)

// Администраторский чат-ид
const adminChatID = 698226393

// Строка подключения к базе данных
const dbConnStr = "postgres://postgres:K#7sd4Na@localhost:5432/vpn_bot?sslmode=disable"

// Структура для хранения информации о подписке
type Subscription struct {
	UserID   int64
	Username string
	Plan     string
	EndDate  time.Time
}

var subscriptions = make(map[int64]Subscription)

func main() {
	// Подключение к базе данных
	db, err := connectToDatabase()
	if err != nil {
		log.Fatalf("Ошибка соединения с базой данных: %v", err)
	}
	defer db.Close()

	// Создаем таблицы, если их еще нет
	err = createTables(db)
	if err != nil {
		log.Fatalf("Ошибка создания таблиц: %v", err)
	}

	// Инициализация бота
	bot, err := tgbotapi.NewBotAPI("8012686496:AAEsvLO0LSu8ooAyt5AMO93L0CxBjz2RNuM")
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
				handleCommand(bot, db, update.Message)
				continue
			}
			if update.Message.Photo != nil {
				handleScreenshot(bot, db, update.Message)
				continue
			}
		}

		if update.CallbackQuery != nil {
			handleApproval(bot, db, update.CallbackQuery)
		}
	}
}

// Подключение к базе данных
func connectToDatabase() (*sql.DB, error) {
	password := url.QueryEscape("12345678") // Замените на ваш пароль
	connStr := "postgres://postgres:" + password + "@localhost:5432/vpn_bot?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// Создание таблиц
func createTables(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		user_id BIGINT UNIQUE NOT NULL,
		username TEXT,
		plan TEXT,
		end_date TIMESTAMP
	);`
	_, err := db.Exec(query)
	return err
}

func handleCommand(bot *tgbotapi.BotAPI, db *sql.DB, msg *tgbotapi.Message) {
	switch msg.Command() {
	case "start":
		welcomeMessage := "Добро пожаловать! Выберите подписку:\n1 месяц - /subscribe_1\n3 месяца - /subscribe_3\n6 месяцев - /subscribe_6"
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, welcomeMessage))
	case "subscribe_1":
		requestPayment(bot, db, msg, "1 месяц")
	case "subscribe_3":
		requestPayment(bot, db, msg, "3 месяца")
	case "subscribe_6":
		requestPayment(bot, db, msg, "6 месяцев")
	default:
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Неизвестная команда."))
	}
}

func requestPayment(bot *tgbotapi.BotAPI, db *sql.DB, msg *tgbotapi.Message, plan string) {
	message := fmt.Sprintf("Вы выбрали подписку на %s.\n\nПожалуйста, переведите оплату на следующие реквизиты: \n*Карта:* 1234 5678 9012 3456\n\nИ отправьте скриншот об оплате.", plan)
	bot.Send(tgbotapi.NewMessage(msg.Chat.ID, message))
	subscriptions[msg.Chat.ID] = Subscription{
		UserID:   msg.Chat.ID,
		Username: msg.From.UserName,
		Plan:     plan,
	}
	saveUserToDatabase(db, msg.Chat.ID, msg.From.UserName, plan)
}

func saveUserToDatabase(db *sql.DB, userID int64, username, plan string) {
	result, err := db.Exec(`
		INSERT INTO users (user_id, username, plan)  
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id) DO UPDATE 
		SET username = EXCLUDED.username, plan = EXCLUDED.plan`,
		userID, username, plan,
	)
	if err != nil {
		log.Printf("Ошибка сохранения пользователя в базу данных: %v", err)
	} else {
		rowsAffected, _ := result.RowsAffected()
		log.Printf("Пользователь %d (%s) успешно сохранен с подпиской: %s. Изменено строк: %d", userID, username, plan, rowsAffected)
	}
}

func handleScreenshot(bot *tgbotapi.BotAPI, db *sql.DB, msg *tgbotapi.Message) {
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

func handleApproval(bot *tgbotapi.BotAPI, db *sql.DB, callback *tgbotapi.CallbackQuery) {
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

		updateSubscriptionEndDate(db, userID, endDate)
	} else if action[:7] == "decline" {
		userID := parseUserID(action)
		response = "Подписка отклонена."
		bot.Send(tgbotapi.NewMessage(userID, "Ваш запрос на подписку был отклонен."))
	}

	bot.Send(tgbotapi.NewMessage(adminChatID, response))

	callbackResponse := tgbotapi.NewCallback(callback.ID, "Действие выполнено.")
	bot.Request(callbackResponse)
}

func updateSubscriptionEndDate(db *sql.DB, userID int64, endDate time.Time) {
	_, err := db.Exec(`UPDATE users SET end_date = $1 WHERE user_id = $2`, endDate, userID)
	if err != nil {
		log.Printf("Ошибка обновления даты окончания подписки: %v", err)
	}
}

func parseUserID(action string) int64 {
	var userID int64
	fmt.Sscanf(action, "approve_%d", &userID)
	return userID
}

func generateKey() string {
	return "KEY-1234-5678" // Здесь можно использовать генератор уникальных ключей
}
