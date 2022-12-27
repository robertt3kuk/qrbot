package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	godotenv.Load()

	go Ginner()
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TGAPI"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			png, _ := qrcode.Encode(update.Message.Text, qrcode.Medium, 256)
			photoFileBytes := tgbotapi.FileBytes{
				Name:  "qr",
				Bytes: png,
			}
			msg := tgbotapi.NewPhoto(update.Message.Chat.ID, photoFileBytes)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}
}
func Ginner() {
	s := gin.Default()
	fmt.Println("runnings")
	s.Run(":8080")

}
