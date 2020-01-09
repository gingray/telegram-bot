package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	ReadConfig()
	fmt.Printf("%+v\n", viper.GetString("telegram_api_key"))
	bot, err := tgbotapi.NewBotAPI(viper.GetString("telegram_api_key"))
	if err != nil {
		log.Panic(err)
	}
	fmt.Print("hello")
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		f, _ := os.Open("/Users/gingray/gosource/src/telegram-bot/files/dance.gif")
		reader := tgbotapi.FileReader{Name: "dance.gif", Reader: f, Size: -1}

		msg2 := tgbotapi.NewAnimationUpload(update.Message.Chat.ID, reader)

		resp, err := bot.Send(msg2)
		log.Printf("%s",fmt.Sprint(resp))
		if err != nil {
			log.Fatal(err)
		}
	}
}
