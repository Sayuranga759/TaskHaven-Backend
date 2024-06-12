package validator

import (
	"github.com/Sayuranga759/TaskHaven-Backend/app/routes/dto"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/custom"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils/constant"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func ValidateTask(requestID string, ctx *fiber.Ctx) (dto.ManageTaskRequest, *custom.ErrorResult) {
	commonLogFields := []zap.Field{zap.String(constant.TraceMsgReqID, requestID)}
	utils.Logger.Debug(utils.TraceMsgFuncStart(ValidateTaskMethod), commonLogFields...)
	defer utils.Logger.Debug(utils.TraceMsgFuncEnd(ValidateTaskMethod), commonLogFields...)

	var (
		request dto.ManageTaskRequest
		body   	= string(ctx.Body())
		err   	error
	)

	err = ctx.BodyParser(&request)
	if err != nil {
		utils.Logger.Error(constant.InvalidInputAndPassErr, append(commonLogFields, []zap.Field{zap.String(constant.ErrorRequestBody, body), zap.Error(err)}...)...)
		errorResult := custom.BuildBadReqErrResult(constant.BindingErrorCode, constant.InvalidRequestErrorMessage, err.Error())
		
		return request, &errorResult
	}

	errRes := ValidateRequest(requestID, request)
	if errRes != nil {
		utils.Logger.Error(utils.TraceMsgErrorOccurredFrom(ValidateRequestMethod), append(commonLogFields, zap.Any(constant.ErrorNote, errRes))...)
		return request, errRes
	}

	return request, nil
}