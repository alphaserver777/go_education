package main

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartScheduler(bot *tgbotapi.BotAPI) {
	log.Println("StartScheduler start")        // Добавлен лог
	ticker := time.NewTicker(10 * time.Second) // Запускаем каждые 10 секунд.
	defer ticker.Stop()

	for range ticker.C {
		go func() {
			log.Println("Start daily question task")
			err := askAllUsers(bot)
			if err != nil {
				log.Println("Error during daily question task: ", err)
			}
		}()
	}
}

func askAllUsers(bot *tgbotapi.BotAPI) error {
	log.Println("askAllUsers start") // Добавлено логирование

	// Получение всех пользователей, которые когда-либо взаимодействовали с ботом.
	// В этом примере у нас нет списка пользователей, поэтому мы отправим всем, кто хоть раз написал.
	// Для более сложной логики, нужно сохранять chat_id каждого пользователя в базу данных.
	// Мы будем брать из ответа id чата.

	// Отправляем сообщение только пользователям у которых есть chat_id в базе данных.
	// Пока предполагаем что все кто написал - хочет получать уведомления.

	users := getUserIdsFromDB()
	log.Println("getUserIdsFromDB users:", users) // Добавлено логирование
	if len(users) == 0 {
		log.Println("No users in database")
		return nil
	}

	for _, user := range users {
		msg := tgbotapi.NewMessage(user, "Болела ли у вас сегодня голова?")
		// Создаем клавиатуру с кнопками
		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Да", "yes"),
				tgbotapi.NewInlineKeyboardButtonData("Нет", "no"),
			),
		)
		msg.ReplyMarkup = keyboard
		sentMsg, err := bot.Send(msg)
		if err != nil {
			log.Println("Error sending daily question to user:", user, "error:", err)
		} else {
			// Сохраняем id сообщения, чтобы потом его можно было идентифицировать.
			messagesToSend[user] = sentMsg.MessageID
		}
	}

	return nil
}

func getUserIdsFromDB() []int64 {
	var userIds []int64

	rows, err := DB.Query("SELECT DISTINCT user_id FROM headache_reports")
	if err != nil {
		log.Println("Error during getting users from db: ", err)
		return userIds
	}
	defer rows.Close()

	for rows.Next() {
		var userId int64
		err = rows.Scan(&userId)
		if err != nil {
			log.Println("Error during scan user_id: ", err)
			continue
		}
		userIds = append(userIds, userId)
	}

	return userIds
}
