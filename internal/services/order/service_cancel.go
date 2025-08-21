package order

import (
	"errors"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/order"

	"gorm.io/gorm"
)

func (s *service) Cancel(ctx core.Context, id int32, user string) (err error) {
	// 先查询 id 是否存在
	_, err = order.NewQueryBuilder().
		WhereIsDeleted(mysql.EqualPredicate, -1).
		WhereId(mysql.EqualPredicate, id).
		First(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return err
	}

	if errors.Is(gorm.ErrRecordNotFound, err) {
		return nil
	}

	data := map[string]interface{}{
		"is_deleted":   1,
		"updated_user": user,
	}

	qb := order.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	return
}
