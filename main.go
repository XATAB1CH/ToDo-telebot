package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/telebot.v3"
)

func main() {

	bot, err := telebot.NewBot(telebot.Settings{
		Token: "6420354925:AAFs0AwespO52MRnRO8Nwj0SPDEknOY-kog",
		Poller: &telebot.LongPoller{
			Timeout: 10 * time.Second,
		},
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	// Обработка команды /start
	bot.Handle("/start", func(c telebot.Context) error {

		button1 := telebot.InlineButton{
			Unique: "showPhotoButton",
			Text:   "Показать фотографию",
		}

		button2 := telebot.InlineButton{
			Unique: "musicButton",
			Text:   "Включить музыку",
		}

		// Создаем сообщение с кнопкой
		replyMarkup := telebot.ReplyMarkup{
			InlineKeyboard: [][]telebot.InlineButton{{button1}, {button2}},
		}

		// Отправляем сообщение с кнопкой
		_, err := bot.Send(c.Chat(), "Выбери команду:", &telebot.SendOptions{
			ReplyMarkup: &replyMarkup,
		})

		return err
	})

	// Обработка нажатия на кнопку "Показать фотографию"
	bot.Handle(&telebot.InlineButton{Unique: "showPhotoButton"}, func(c telebot.Context) error {
		// Создаем объект фотографии (замени путь на свой файл)
		photo := &telebot.Photo{File: telebot.FromDisk("tulip.jpg")}

		// Отправляем фотографию
		_, err := bot.Send(c.Chat(), photo, &telebot.SendOptions{})
		return err
	})

	// Обработка нажатия на кнопку "Включить музыыку"
	bot.Handle(&telebot.InlineButton{Unique: "musicButton"}, func(c telebot.Context) error {
		audio := &telebot.Audio{
			File:     telebot.FromDisk("spectr.mp3"), // Путь к вашему аудиофайлу
			FileName: "spectr.mp3",
		}

		_, err := bot.Send(c.Chat(), audio, &telebot.SendOptions{})
		return err
	})

	bot.Handle("/age", func(c telebot.Context) error {
		_, err := bot.Send(c.Chat(), "100", &telebot.SendOptions{})
		return err
	})

	bot.Handle("/fuckyou", func(c telebot.Context) error {
		_, err := bot.Send(c.Chat(), "пишёл нахуй", &telebot.SendOptions{})
		return err
	})

	bot.Handle("/currentdatetime", func(c telebot.Context) error {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		response := fmt.Sprintf("Текущая дата и время: %s", currentTime)

		_, err := bot.Send(c.Chat(), response, &telebot.SendOptions{})
		return err
	})

	bot.Start()
}
