package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var messagesToSend = make(map[int64]int) // id сообщения, чтобы потом можно было идентифицировать какой вопрос.

func main() {
	// Загружаем переменные окружения из .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	databaseURL := os.Getenv("DATABASE_URL")

	log.Println("Начало инициализации базы данных")
	// Инициализация базы данных
	initDB(databaseURL)

	// Инициализация бота
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Error initializing bot: %v", err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s\n", bot.Self.UserName)

	// Запуск планировщика
	go StartScheduler(bot)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Обрабатываем обновления
	for update := range updates {
		if update.CallbackQuery != nil {
			handleCallback(bot, update)
			continue
		}

		if update.Message == nil {
			continue
		}

		handleMessage(bot, update)
	}
}

func handleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log.Println("handleMessage start")
	userID := update.Message.From.ID
	text := update.Message.Text
	log.Printf("Message received: %s", text)

	if text == "/start" {
		log.Println("start command handled")
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я бот для отслеживания головной боли. Я буду спрашивать у тебя каждый день: Болела ли у тебя голова?")
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("Error sending welcome message: ", err)
		}
		// Сохранение user_id в базу данных
		err = SaveReport(userID, "first_start_message") // Сохранение user_id при первом старте
		if err != nil {
			log.Println("Error saving report in handleMessage: ", err)
		}

	} else {
		log.Printf("Message from %d: %s\n", userID, text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я вас понял. Спасибо за ответ.")
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("Error sending answer for message: ", err)
		}
	}
	log.Println("handleMessage end")
}

func handleCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	userID := update.CallbackQuery.From.ID
	answer := update.CallbackQuery.Data
	messageId := update.CallbackQuery.Message.MessageID

	if sentMessageId, ok := messagesToSend[userID]; ok && sentMessageId == messageId {
		err := SaveReport(userID, answer)
		if err != nil {
			log.Println("Error saving report: ", err)
			callbackConfig := tgbotapi.NewCallback(update.CallbackQuery.ID, "Произошла ошибка при сохранении ответа.")
			if _, err := bot.Request(callbackConfig); err != nil {
				log.Println("Error answering callback:", err)
			}
		} else {
			callbackConfig := tgbotapi.NewCallback(update.CallbackQuery.ID, fmt.Sprintf("Ответ '%s' записан", answer))
			if _, err := bot.Request(callbackConfig); err != nil {
				log.Println("Error answering callback:", err)
			}
		}
		// Убираем клавиатуру
		msg := tgbotapi.NewEditMessageReplyMarkup(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
			tgbotapi.NewInlineKeyboardMarkup(),
		)
		if _, err = bot.Send(msg); err != nil {
			log.Println("Error remove keyboard: ", err)
		}
		delete(messagesToSend, userID)
	} else {
		callbackConfig := tgbotapi.NewCallback(update.CallbackQuery.ID, "Это сообщение устарело")
		if _, err := bot.Request(callbackConfig); err != nil {
			log.Println("Error answering callback:", err)
		}
	}
}

// Запускаем простой сервер для проверки работы.
// Чтобы не было проблем с Heroku, и можно было развернуть приложение
func init() {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Telegram bot is running")
		})
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}()
}
