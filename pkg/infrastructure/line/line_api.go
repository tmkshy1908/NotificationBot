package line

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"
)

type LineConf struct {
	Bot *linebot.Client
}

type LineClient interface {
}

func NewLineClient() (lc LineClient, err error) {
	bot, err := linebot.New(
		"aaaaaa",
		"aaaaaa",
	)
	if err != nil {
		fmt.Println("##NewLineClient err##", err)
	} else {
		fmt.Println("##LineClient OK##")
	}

	lc = &LineConf{Bot: bot}
	return
}
