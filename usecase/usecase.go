package usecase

import (
	"net/http"
)

type CommonInteractor struct {
	CommonRepository CommonRepository
}

func (i *CommonInteractor) DivideMessage(req *http.Request) {

	msg, userId := i.CommonRepository.DivideEvent(req)
	i.CommonRepository.CallReply(msg, userId)

}
