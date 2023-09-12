package interfaces

import (
	"github.com/tmkshy1908/NotificationBot/pkg/infrastructure/line"
)

type CommonRepository struct {
	Bot line.LineClient
}
