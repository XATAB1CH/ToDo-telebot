package main

import (
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	b, err := tele.NewBot(tele.Settings{
		Token:  "...",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return
	}

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello world!")
	})

	b.Start()
}
