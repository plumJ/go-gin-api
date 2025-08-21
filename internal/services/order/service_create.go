package order

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/order"
	"time"
)

type CreateOrderData struct {
	OrderNo     string    // 订单号
	OrderFee    int32     // 订单金额(分)
	Status      int32     // 订单状态 1:已创建  2:已取消
	IsDeleted   int32     // 是否删除 1:是  -1:否
	CreatedAt   time.Time `gorm:"time"` // 创建时间
	UpdatedAt   time.Time `gorm:"time"` // 更新时间
	CreatedUser string    // 创建人
	UpdatedUser string    // 更新人
}

func (s *service) Create(ctx core.Context, orderData *CreateOrderData) (id int32, err error) {
	orderModel := order.NewModel()
	orderModel.OrderNo = orderData.OrderNo
	orderModel.OrderFee = orderData.OrderFee
	orderModel.Status = orderData.Status
	orderModel.IsDeleted = orderData.IsDeleted
	orderModel.CreatedAt = orderData.CreatedAt
	orderModel.UpdatedAt = orderData.UpdatedAt
	orderModel.CreatedUser = orderData.CreatedUser
	orderModel.UpdatedUser = orderData.UpdatedUser

	id, err = orderModel.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
