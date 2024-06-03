package responsebuilder

import (
	"net/http"

	"github.com/Sayuranga759/TaskHaven-Backend/pkg/custom"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils/constant"
	"github.com/gofiber/fiber/v2"

	"go.uber.org/zap"
)

// APIResponse use to define api response
type APIResponse struct {
	_             struct{}
	Ctx           *fiber.Ctx
	Response      any
	RequestID     string
	HTTPStatus    int
	ErrorResponse custom.ErrorResult
}

func (response *APIResponse) BuildAPIResponse() {
	commonLogFields := []zap.Field{zap.String(constant.TraceMsgReqID, response.RequestID)}
	utils.Logger.Debug(constant.TraceMsgAPIResponse, commonLogFields...)

	if response.ErrorResponse.IsError {
		utils.Logger.Debug(constant.TraceMsgAPIErrorResponse, append(commonLogFields, zap.Any(constant.TraceMsgReqBody, response.ErrorResponse))...)

		if response.HTTPStatus == 0 {
			response.HTTPStatus = http.StatusInternalServerError
		}

		err := response.Ctx.Status(response.HTTPStatus).JSON(response.ErrorResponse.ErrorList)
		if err != nil {
			utils.Logger.Error(constant.TraceMsgAPIErrorResponse, append(commonLogFields, zap.Error(err))...)
		}
	} else {
		utils.Logger.Debug(constant.TraceMsgAPISuccess, commonLogFields...)

		err := response.Ctx.Status(response.HTTPStatus).JSON(response.Response)
		if err != nil {
			utils.Logger.Error(constant.TraceMsgAPIErrorResponse, append(commonLogFields, zap.Error(err))...)
		}
	}
}
