package usecase

import (
	"context"
	"fmt"
	"net/http"
)

type CommonInteractor struct {
	CommonRepository CommonRepository
}

func (i *CommonInteractor) RootMain(ctx context.Context, req *http.Request) {

	msg, userId := i.CommonRepository.DivideEvent(ctx, req)
	users, b := i.CommonRepository.DivideMessage(ctx, userId, msg)
	if b == true {
		fmt.Println(users)
		err := i.CommonRepository.Add(ctx, users)
		if err != nil {
			fmt.Println(err)
		}
	}
	i.CommonRepository.CallReply(msg, userId)

	go i.CommonRepository.Alarm(ctx, userId, users)
	i.CommonRepository.Delete(ctx, users)

}
