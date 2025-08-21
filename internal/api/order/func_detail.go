package order

import (
	"net/http"
	"time"

	"github.com/spf13/cast"
	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/services/order"
)

type detailRequest struct {
	Id string `uri:"id"` // id
}

type detailResponse struct {
	Id          int32     `json:"id"`
	OrderNo     string    `json:"order_no"`
	OrderFee    int32     `json:"order_fee"`
	Status      int32     `json:"status"`
	IsDeleted   int32     `json:"is_deleted"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedUser string    `json:"created_user"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedUser string    `json:"updated_user"`
}

// Detail 订单详情
// @Summary 订单详情
// @Description 订单详情
// @Tags API.order
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} detailResponse
// @Failure 400 {object} code.Failure
// @Router /api/order/{id} [get]
func (h *handler) Detail() core.HandlerFunc {
	return func(c core.Context) {
		req := new(detailRequest)
		res := new(detailResponse)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		searchOneData := new(order.SearchOneData)
		searchOneData.Id = cast.ToInt32(req.Id)

		detail, err := h.orderService.Detail(c, searchOneData)
		if err != nil {
			return
		}
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.OrderDetailError,
				code.Text(code.OrderDetailError)).WithError(err),
			)
			return
		}

		res.Id = detail.Id
		res.OrderNo = detail.OrderNo
		res.OrderFee = detail.OrderFee
		res.Status = detail.Status
		res.IsDeleted = detail.IsDeleted
		res.CreatedAt = detail.CreatedAt
		res.UpdatedAt = detail.UpdatedAt
		res.CreatedUser = detail.CreatedUser
		res.UpdatedUser = detail.UpdatedUser

		c.Payload(res)
	}
}
