package service

import (
	"context"
	"strconv"

	"go-common/app/job/main/block/model"

	"github.com/pkg/errors"
)

func (s *Service) accountNotify(c context.Context, mid int64) (err error) {
	msg := &model.AccountNotify{UID: mid, Action: "blockUser"}
	if err = s.accountNotifyPub.Send(c, strconv.FormatInt(msg.UID, 10), msg); err != nil {
		err = errors.Errorf("mid(%d) s.accountNotify.Send(%+v) error(%+v)", msg.UID, msg, err)
	}
	return
}
