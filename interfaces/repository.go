package interfaces

import (
	"net/http"

	"github.com/tmkshy1908/NotificationBot/pkg/infrastructure/line"
)

type CommonRepository struct {
	Bot line.LineClient
}

func (r *CommonRepository) DivideEvent(req *http.Request) (msg string, userId string) {
	msg, userId = r.Bot.CathEvents(req)
	return
}

func (r *CommonRepository) CallReply(msg string, userId string) {
	r.Bot.MsgReply(msg, userId)
}
