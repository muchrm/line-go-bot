package main

import (
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
	users := []string{}
	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					userID := event.Source.UserID

					for _, user := range users {
						if userID != user {
							if _, err = bot.PushMessage(user, linebot.NewTextMessage(message.Text)).Do(); err != nil {
								log.Print(err)
							}
						}
					}
					if found, _ := inArray(userID, users); found != true {
						users = append(users, userID)
					}
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ข้อความถูกส่งเรียบร้อยแล้ว")).Do()
				}
			}
		}
	})
	// This is just a sample code.
	// For actually use, you must support HTTPS by using `ListenAndServeTLS`, reverse proxy or etc.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
func inArray(val string, array []string) (ok bool, i int) {
	for i = range array {
		if ok = array[i] == val; ok {
			return
		}
	}
	return
}
