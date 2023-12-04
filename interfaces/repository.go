package interfaces

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/tmkshy1908/NotificationBot/domain"
	"github.com/tmkshy1908/NotificationBot/pkg/infrastructure/db"
	"github.com/tmkshy1908/NotificationBot/pkg/infrastructure/line"
)

type CommonRepository struct {
	DB  db.SqlHandler
	Bot line.LineClient
}

const (
	INSERT_USERTIME string = "insert into usertime (id, hour, minute) values ($1,$2,$3)"
	DELETE_USERTIME string = "delete from usertime where id = $1"
)

func (r *CommonRepository) Add(ctx context.Context, usertime *domain.UserTime) (err error) {
	err = r.DB.ExecWithTx(func(*sql.Tx) error {
		_, err = r.DB.Exec(ctx, INSERT_USERTIME, usertime.Id, usertime.Hour, usertime.Minute)
		if err != nil {
			fmt.Println(err, "##Rep_Exec##")
		}
		return err
	})
	return
}

func (r *CommonRepository) Delete(ctx context.Context, usertime *domain.UserTime) (err error) {
	err = r.DB.ExecWithTx(func(*sql.Tx) error {
		_, err = r.DB.Exec(ctx, DELETE_USERTIME, usertime.Id)
		if err != nil {
			fmt.Println(err, "##Rep_Delete##")
		}
		return err
	})
	return
}

var timeMatcher = regexp.MustCompile("([0-9]+)時([0-9]+)分")

func (r *CommonRepository) DivideEvent(ctx context.Context, req *http.Request) (umsg *domain.UserMsg, err error) {
	umsg, err = r.Bot.CathEvents(ctx, req)
	return
}

func (r *CommonRepository) DivideMessage(ctx context.Context, umsg *domain.UserMsg) (usertime *domain.UserTime, b bool, err error) {
	m := timeMatcher.FindStringSubmatch(umsg.Message)
	if len(m) == 3 {
		hour, err := strconv.Atoi(m[1])
		if err != nil {
			// return "何時ですか？"
		}
		minute, err := strconv.Atoi(m[2])
		if err != nil {
			// return "何分ですか？"
		}
		// err = h.store.Set(userId, hour, minute)
		if err != nil {
			// return "時間の設定に失敗しました"
		}
		userId := umsg.Id
		usertime = &domain.UserTime{Id: userId, Hour: hour, Minute: minute}
		msg := fmt.Sprintf("%v時%v分ですね。わかりました。", hour, minute)
		umsg := &domain.UserMsg{Id: userId, Message: msg}
		r.Bot.MsgReply(umsg)

		return usertime, true, err
	} else {
		return nil, false, err
	}
}

func (r *CommonRepository) CallReply(umsg *domain.UserMsg) (err error) {
	if err = r.Bot.MsgReply(umsg); err != nil {
		return err
	}
	return
}

func (r *CommonRepository) Alarm(ctx context.Context, usertime *domain.UserTime) (err error) {
	if usertime != nil {
		for {
			now := time.Now()
			timeValue := time.Date(now.Year(), now.Month(), now.Day(), usertime.Hour, usertime.Minute, 0, 0, now.Location())
			if now.After(timeValue) {
				timeValue = timeValue.Add(24 * time.Hour)
				msg := "時間です"
				umsg := &domain.UserMsg{Id: usertime.Id, Message: msg}
				if err = r.Bot.MsgReply(umsg); err != nil {
					fmt.Println(err)
					return err
				}
				if err = r.Delete(ctx, usertime); err != nil {
					fmt.Println(err)
					return err
				}
			}
			time.Sleep(timeValue.Sub(now))
		}
	}
	return
}
