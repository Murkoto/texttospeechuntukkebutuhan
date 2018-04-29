import (
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	bot, err := linebot.New("6500e2b36c742d5f3988981a5cf30514", 
		"90JkB4mN6PuwC506j/A+jSPlXWYGoW8ag6c5R9+R5AQRN0QiL46V0uZ+uZD+KBPo91gWoPWBuLV5QfiAPEM6FJkzeWDWO5o1SYZ9Pw4O1UYqF46e/CekYFLuBMg2oQSEYVUc37lIS/SEo1gcWtyZ/AdB04t89/1O/w1cDnyilFU=")

	event, error := bot.ParseRequest(req)
	if err != nil {
		
	}
	if event.Type == linebot.EventTypeMessage {
		bot.ReplyMessage(event.ReplyToken, event.Text).Do()	
	}

}