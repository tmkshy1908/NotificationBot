package interfaces

import (
	"context"
	"net/http"
	"time"

	"github.com/tmkshy1908/NotificationBot/pkg/infrastructure/line"
)

type CommonRepository struct {
	Bot line.LineClient
}

func (r *CommonRepository) DivideEvent(ctx context.Context, req *http.Request) (msg string, userId string) {
	msg, userId = r.Bot.CathEvents(ctx, req)
	return
}

func (r *CommonRepository) CallReply(msg string, userId string) {
	r.Bot.MsgReply(msg, userId)
}

func (r *CommonRepository) TimeSet(ctx context.Context, userId string) {
	for {
		now := time.Now()
		timeValue := time.Date(now.Year(), now.Month(), now.Day(), 06, 13, 0, 0, now.Location())
		if now.After(timeValue) {
			timeValue = timeValue.Add(24 * time.Hour)
			msg := "時間です"
			r.Bot.MsgReply(msg, userId)
		}
		time.Sleep(timeValue.Sub(now))
	}
}
