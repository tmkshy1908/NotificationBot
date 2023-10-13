package usecase

import (
	"context"
	"net/http"
)

type CommonInteractor struct {
	CommonRepository CommonRepository
}

func (i *CommonInteractor) DivideMessage(ctx context.Context, req *http.Request) {

	msg, userId := i.CommonRepository.DivideEvent(ctx, req)
	i.CommonRepository.CallReply(msg, userId)
	go i.TimeAlarm(ctx)

}

func (i *CommonInteractor) TimeAlarm(ctx context.Context) {
	userId := "U01db659616022939affa0bdd806b9479"
	i.CommonRepository.TimeSet(ctx, userId)
}
