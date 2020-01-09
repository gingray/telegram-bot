package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
	"log"
	"os"
)

const TelegramAPIKey = "telegram_api_key"
const RedisUrl = "redis_url"
const RedisPort = "redis_port"

func main() {
	ReadConfig()
	InitStorage(viper.GetString(RedisUrl), viper.GetInt(RedisPort))

	fmt.Printf("%+v\n", viper.GetString("telegram_api_key"))
	bot, err := tgbotapi.NewBotAPI(viper.GetString(TelegramAPIKey))
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
		f, _ := os.Open(viper.GetString("test_file_path"))
		reader := tgbotapi.FileReader{Name: "dance.gif", Reader: f, Size: -1}
		msg2 := tgbotapi.NewAnimationUpload(update.Message.Chat.ID, reader)

		resp, err := bot.Send(msg2)

		if err != nil {
			log.Fatal(err)
		} else {
			bytes, _ := json.Marshal(resp)
			item := Storage{FileID: resp.Animation.FileID, Response: string(bytes)}
			log.Printf("zzz %s", fmt.Sprint(resp))
			SaveResponse(item)
		}
	}
}
