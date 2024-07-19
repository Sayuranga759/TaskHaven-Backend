package handler

import (
	"github.com/Sayuranga759/TaskHaven-Backend/app/routes/dto"
	"github.com/Sayuranga759/TaskHaven-Backend/app/routes/handler/validator"
	"github.com/Sayuranga759/TaskHaven-Backend/app/service"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/custom"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils/constant"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/web"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/web/responsebuilder"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func UserRegistrationHandler(ctx *fiber.Ctx) error {
	var requestID = web.GetRequestID(ctx)
	commonLogFields := utils.CommonLogField(requestID)
	utils.Logger.Info(utils.TraceMsgFuncStart(UserRegistrationHandlerMethod), commonLogFields...)

	defer utils.Logger.Info(utils.TraceMsgFuncEnd(UserRegistrationHandlerMethod), commonLogFields...)

	var (
		statusCode 	int
		request 	dto.UserRegistrationRequest
		errorResult *custom.ErrorResult
		errRes 		custom.ErrorResult
		response 	*dto.UserRegistrationResponse
		userService = service.CreateUserSerivce(requestID)
	)

	// validate
	utils.Logger.Debug(utils.TraceMsgBeforeInvoke(validator.ValidateUserRegistrationMethod), commonLogFields...)
	request, errorResult = validator.ValidateUserRegistration(requestID, ctx)
	utils.Logger.Debug(utils.TraceMsgAfterInvoke(validator.ValidateUserRegistrationMethod), commonLogFields...)

	if errorResult == nil {
		response, errorResult = userService.RegisterUser(request)
	}

	if errorResult != nil {
		logFields := append(commonLogFields, zap.Any(constant.ErrorNote, errorResult))
		utils.Logger.Error(utils.TraceMsgErrorOccurredFrom(service.RegisterUserMethod), logFields...)

		statusCode, errRes = HandleError(errorResult)
	}

	// Build the response
	responseBuilder := responsebuilder.APIResponse {
		Ctx: 			ctx,
		HTTPStatus: 	statusCode,
		ErrorResponse: 	errRes,
		Response: 		response,
		RequestID: 		requestID,
	}
	responseBuilder.BuildAPIResponse()

	return nil
}

func UserLoginHandler(ctx *fiber.Ctx) error {
	var requestID = web.GetRequestID(ctx)
	commonLogFields := utils.CommonLogField(requestID)
	utils.Logger.Info(utils.TraceMsgFuncStart(UserLoginHandlerMethod), commonLogFields...)

	defer utils.Logger.Info(utils.TraceMsgFuncEnd(UserLoginHandlerMethod), commonLogFields...)

	var (
		statusCode 	int
		request 	dto.LoginRequest
		errorResult *custom.ErrorResult
		errRes 		custom.ErrorResult
		response 	*dto.LoginResponse
		userService = service.CreateUserSerivce(requestID)
	)

	// validate
	utils.Logger.Debug(utils.TraceMsgBeforeInvoke(validator.ValidateLoginMethod), commonLogFields...)
	request, errorResult = validator.ValidateLogin(requestID, ctx)
	utils.Logger.Debug(utils.TraceMsgAfterInvoke(validator.ValidateLoginMethod), commonLogFields...)

	if errorResult == nil {
		response, errorResult = userService.Login(request, ctx)
	}

	if errorResult != nil {
		logFields := append(commonLogFields, zap.Any(constant.ErrorNote, errorResult))
		utils.Logger.Error(utils.TraceMsgErrorOccurredFrom(service.LoginMethod), logFields...)

		statusCode, errRes = HandleError(errorResult)
	} 

	// Build the response
	responseBuilder := responsebuilder.APIResponse {
		Ctx: 			ctx,
		HTTPStatus: 	statusCode,
		ErrorResponse: 	errRes,
		Response: 		response,
		RequestID: 		requestID,
	}
	responseBuilder.BuildAPIResponse()

	return nil
}