package usecase

import (
	"context"
	"net/http"

	"github.com/tmkshy1908/NotificationBot/domain"
)

type CommonInteractor struct {
	CommonRepository CommonRepository
}

func (i *CommonInteractor) RootMain(ctx context.Context, req *http.Request) (err error) {
	tc := make(chan *domain.UserTime)
	umsg, err := i.CommonRepository.DivideEvent(ctx, req)
	if err != nil {
		return err
	}
	usertime, b, err := i.CommonRepository.DivideMessage(ctx, umsg, tc)
	if err != nil {
		return err
	}
	if b == true {
		if err := i.CommonRepository.Add(ctx, usertime); err != nil {
			return err
		}
	}
	if err = i.CommonRepository.CallReply(umsg); err != nil {
		return err
	}
	at := <-tc
	ctx = context.Background()
	go i.CommonRepository.Alarm(ctx, at)
	return
}
