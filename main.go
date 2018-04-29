package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot linebot.Client

func main() {
	bot, err := linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	http.HandleFunc("/callback", callbackHandler)
	port := "1234"
	// port := "1234"
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request){
	events, err := bot.ParseRequest(r)
	if err != nil {
		
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			bot.ReplyMessage(event.ReplyToken, event.Message.Text).Do()	
		}
	}
}