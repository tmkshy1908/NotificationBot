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

func (r *CommonRepository) Add(ctx context.Context, ustates *domain.UserStates) (err error) {
	if ustates.Mode == "dbAdd" {
		err = r.DB.ExecWithTx(func(*sql.Tx) error {
			_, err = r.DB.Exec(ctx, INSERT_USERTIME, ustates.Id, ustates.Hour, ustates.Minute)
			if err != nil {
				fmt.Println(err, "##Rep_Exec##")
			}
			return err
		})
	} else {
		return
	}
	return
}

func (r *CommonRepository) Delete(ctx context.Context, usertime *domain.UserStates) (err error) {
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

func (r *CommonRepository) DivideEvent(ctx context.Context, req *http.Request) (umsg *domain.UserStates, err error) {
	umsg, err = r.Bot.CathEvents(ctx, req)
	return
}

// var userstates = make(map[string]*domain.UserStates)

// var usertime *domain.UserStates

func (r *CommonRepository) DivideMessage(ctx context.Context, umsg *domain.UserStates) (states *domain.UserStates, err error) {
	userId := umsg.Id
	// currentState := userstates[userId]

	switch umsg.Mode {
	case "":
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

			msg := fmt.Sprintf("%v時%v分でよろしいですか？", hour, minute)
			// umsg := &domain.UserMsg{Id: userId, }
			states = &domain.UserStates{Id: userId, Message: msg, Mode: "confirming", Hour: hour, Minute: minute}
			r.Bot.MsgReply(umsg)
			// userstates[userId] = &domain.UserStates{Mode: "confirming"}
		} else {
			states = &domain.UserStates{}
			return
		}
	case "confirming":
		if umsg.Message == "はい" {
			msg := fmt.Sprintf("%v時%v分でセットしました。", umsg.Hour, umsg.Minute)
			umsg := &domain.UserStates{Id: userId, Message: msg, Mode: "dbAdd"}
			r.Bot.MsgReply(umsg)
			// delete(userStates, userId) // 状態をリセット
		} else {
			msg := "アラームをキャンセルしました。"
			umsg := &domain.UserStates{Id: userId, Message: msg}
			r.Bot.MsgReply(umsg)
			// delete(userStates, userId)
		}

	}
	return
}

func (r *CommonRepository) CallReply(umsg *domain.UserStates) (err error) {
	if err = r.Bot.MsgReply(umsg); err != nil {
		return err
	}
	return
}

func (r *CommonRepository) Alarm(ctx context.Context, tc <-chan *domain.UserStates) (err error) {
	usertime := <-tc
	for {
		now := time.Now()
		timeValue := time.Date(now.Year(), now.Month(), now.Day(), usertime.Hour, usertime.Minute, 0, 0, now.Location())
		if now.After(timeValue) {
			timeValue = timeValue.Add(24 * time.Hour)
			msg := "時間です"
			umsg := &domain.UserStates{Id: usertime.Id, Message: msg}
			if err = r.Bot.MsgReply(umsg); err != nil {
				return err
			}
			if err = r.Delete(ctx, usertime); err != nil {
				return err
			}
		}
		time.Sleep(timeValue.Sub(now))
	}
}
