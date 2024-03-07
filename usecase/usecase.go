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
	tc := make(chan *domain.UserStates)
	umsg, err := i.CommonRepository.DivideEvent(ctx, req)
	if err != nil {
		return err
	}

	ustates, err := i.CommonRepository.DivideMessage(ctx, umsg)
	if err != nil {
		return err
	}

	if err := i.CommonRepository.Add(ctx, ustates); err != nil {
		return err
	}

	if err = i.CommonRepository.CallReply(ustates); err != nil {
		return err
	}
	// at := <-tc
	tc <- ustates
	close(tc)
	ctx = context.Background()
	go i.CommonRepository.Alarm(ctx, tc)
	return
}
