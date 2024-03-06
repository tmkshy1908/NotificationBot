package line

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/tmkshy1908/NotificationBot/domain"
)

type LineConf struct {
	Bot *linebot.Client
}

type LineClient interface {
	CathEvents(context.Context, *http.Request) (*domain.UserStates, error)
	MsgReply(*domain.UserStates) error
}

func NewLineClient() (lc LineClient, err error) {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("ACCESS_TOKEN"),
	)
	if err != nil {
		fmt.Println("##NewLineClient err##")
		return nil, err
	} else {
		fmt.Println("LineClient Connected.")
	}

	lc = &LineConf{Bot: bot}
	return
}

func (bot *LineConf) CathEvents(ctx context.Context, req *http.Request) (umsg *domain.UserStates, err error) {
	events, err := bot.Bot.ParseRequest(req)
	if err != nil {
		fmt.Println("ParseReq", err)
	}
	umsg = &domain.UserStates{Id: "", Message: ""}
	for _, event := range events {

		if event.Type == linebot.EventTypeMessage {
			userId := event.Source.UserID

			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				msg := message.Text
				umsg = &domain.UserStates{Id: userId, Message: msg}

			case *linebot.StickerMessage:
				msg := "いいスタンプだね"
				umsg = &domain.UserStates{Id: userId, Message: msg}

			case *linebot.ImageMessage:
				msg := "いい写真だね"
				umsg = &domain.UserStates{Id: userId, Message: msg}

			}
		} else {
			fmt.Println("EventTypeが違う")
		}
	}
	return
}

func (bot *LineConf) MsgReply(umsg *domain.UserStates) (err error) {
	userId := umsg.Id
	msg := umsg.Message
	if _, err := bot.Bot.PushMessage(userId, linebot.NewTextMessage(msg)).Do(); err != nil {
		fmt.Println("##MsgReply Pushmessage##")
		return err
	}
	return
}
