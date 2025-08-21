package order

import (
	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/validation"
	"github.com/xinliangnote/go-gin-api/internal/services/order"
	"net/http"
)

type createRequest struct {
	OrderNo     string `form:"OrderNo" binding:"required"`     // 订单序列号
	OrderFee    int32  `form:"OrderFee" binding:"required"`    // 订单金额
	CreatedUser string `form:"CreatedUser" binding:"required"` // 创建人
}

type createResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// Create 创建订单
// @Summary 创建订单
// @Description 创建订单
// @Tags API.order
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param OrderNo formData string true "订单序列号"
// @Param OrderFee formData int32 true "订单金额"
// @Param CreatedUser formData string true "创建者"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/order/create [post]
func (h *handler) Create() core.HandlerFunc {
	return func(c core.Context) {
		req := new(createRequest)
		res := new(createResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err),
			)
			return
		}

		createData := new(order.CreateOrderData)
		createData.OrderNo = req.OrderNo
		createData.OrderFee = req.OrderFee
		createData.CreatedUser = req.CreatedUser
		createData.UpdatedUser = req.CreatedUser
		createData.Status = 1
		createData.IsDeleted = -1

		id, err := h.orderService.Create(c, createData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.OrderCreateError,
				code.Text(code.OrderCreateError)).WithError(err),
			)
			return
		}

		res.Id = id
		c.Payload(res)
	}
}
