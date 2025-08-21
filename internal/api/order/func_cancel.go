package order

import (
	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"net/http"
)

type cancelRequest struct {
	Id   int32  `form:"id"`   // 主键ID
	User string `form:"user"` // user
}

type cancelResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// Cancel 取消订单
// @Summary 取消订单
// @Description 取消订单
// @Tags API.order
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id formData int true "id"
// @Param user formData string true "user"
// @Success 200 {object} cancelResponse
// @Failure 400 {object} code.Failure
// @Router /api/order/cancel [patch]
func (h *handler) Cancel() core.HandlerFunc {
	return func(c core.Context) {
		req := new(cancelRequest)
		res := new(cancelResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		err := h.orderService.Cancel(c, req.Id, req.User)

		if err != nil {
			return
		}
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.OrderCancelError,
				code.Text(code.OrderCancelError)).WithError(err),
			)
			return
		}

		res.Id = req.Id
		c.Payload(res)
	}
}
