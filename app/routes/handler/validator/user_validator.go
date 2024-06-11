package validator

import (
	"github.com/Sayuranga759/TaskHaven-Backend/app/routes/dto"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/custom"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils/constant"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func ValidateUserRegistration(requestID string, ctx *fiber.Ctx) (dto.UserRegistrationRequest, *custom.ErrorResult) {
	commonLogFields := []zap.Field{zap.String(constant.TraceMsgReqID, requestID)}
	utils.Logger.Debug(utils.TraceMsgFuncStart(ValidateUserRegistrationMethod), commonLogFields...)
	defer utils.Logger.Debug(utils.TraceMsgFuncEnd(ValidateUserRegistrationMethod), commonLogFields...)

	var (
		request dto.UserRegistrationRequest
		body   	= string(ctx.Body())
		err    	error
	)

	ctx.Request().SetBody([]byte(body))

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

func ValidateLogin(requestID string, ctx *fiber.Ctx) (dto.LoginRequest, *custom.ErrorResult) {
	commonLogFields := []zap.Field{zap.String(constant.TraceMsgReqID, requestID)}
	utils.Logger.Debug(utils.TraceMsgFuncStart(ValidateLoginMethod), commonLogFields...)
	defer utils.Logger.Debug(utils.TraceMsgFuncEnd(ValidateLoginMethod), commonLogFields...)

	var (
		request dto.LoginRequest
		body   	= string(ctx.Body())
		err    	error
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