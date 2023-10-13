package line

import (
	"context"
	"fmt"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

type LineConf struct {
	Bot *linebot.Client
}

type LineClient interface {
	CathEvents(context.Context, *http.Request) (string, string)
	MsgReply(string, string)
}

func NewLineClient() (lc LineClient, err error) {
	bot, err := linebot.New(
		"e29ca6d4357cc977c61591eae223aef4",
		"FWbGlGdJpRuj5snIzBYGHIxSWrLJ2usY2mrsdqjsFplBZzg4DpFZxOAu+VZAUkvz7Sr3IAv51KmQJRh1T9z2HZmDWS/1cOfrb2HwtC0+GqDJKoiGiwkTmsL5Ar3S/vWYw2Yn8Vz1YvrEVzay36EJvwdB04t89/1O/w1cDnyilFU=",
	)
	if err != nil {
		fmt.Println("##NewLineClient err##", err)
	} else {
		fmt.Println("##LineClient OK##")
	}

	lc = &LineConf{Bot: bot}
	return
}

func (bot *LineConf) CathEvents(ctx context.Context, req *http.Request) (msg string, userId string) {
	events, err := bot.Bot.ParseRequest(req)
	if err != nil {
		fmt.Println("ParseReq", err)
	}
	for _, event := range events {

		if event.Type == linebot.EventTypeMessage {
			userId = event.Source.UserID

			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				msg = message.Text
				bot.MsgReply(msg, userId)
				// if _, err := bot.Bot.PushMessage(userId, linebot.NewTextMessage(msg)).Do(); err != nil {
				// 	fmt.Println(err, "プッシュエラー")
				// }

			case *linebot.StickerMessage:
				bot.MsgReply("いいスタンプだね", userId)

			case *linebot.ImageMessage:
				bot.MsgReply("いい画像だね", userId)
			}
		} else {
			fmt.Println("EventTypeが違う")
		}
	}
	return
}

func (bot *LineConf) MsgReply(msg string, userId string) {
	if _, err := bot.Bot.PushMessage(userId, linebot.NewTextMessage(msg)).Do(); err != nil {
		fmt.Println(err, "プッシュエラー")
	}
	fmt.Println(userId)
	// replyMessage := linebot.NewTextMessage(msg)
	// bot.Bot.BroadcastMessage(replyMessage).Do()
}
