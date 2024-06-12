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

func CreateTaskHandler(ctx *fiber.Ctx) error {
	var requestID = web.GetRequestID(ctx)
	commonLogFields := []zap.Field{zap.String(constant.TraceMsgReqID, requestID)}
	utils.Logger.Info(utils.TraceMsgFuncStart(CreateTaskHandlerMethod), commonLogFields...)

	defer utils.Logger.Info(utils.TraceMsgFuncEnd(CreateTaskHandlerMethod), commonLogFields...)

	var (
		statusCode  int
		errorResult *custom.ErrorResult
		errRes      custom.ErrorResult
		request	 	dto.CreateTaskRequest
		response    *dto.ManageTaskResponse
		taskService = service.CreateTaskSerivce(requestID)
	)

	// validate
	utils.Logger.Debug(utils.TraceMsgBeforeInvoke(validator.ValidateCreateTaskMethod), commonLogFields...)
	request, errorResult = validator.ValidateCreateTask(requestID, ctx)
	utils.Logger.Debug(utils.TraceMsgAfterInvoke(validator.ValidateCreateTaskMethod), commonLogFields...)

	request.UserID = ctx.Locals(TokenClaims).(*dto.JWTClaims).UserID

	if errorResult == nil {
		response, errorResult = taskService.CreateTask(request)
	}

	if errorResult != nil {
		logFields := append(commonLogFields, zap.Any(constant.ErrorNote, errorResult))
		utils.Logger.Error(utils.TraceMsgErrorOccurredFrom(service.CreateTaskMethod), logFields...)

		statusCode, errRes = HandleError(errorResult)
	}

	// Build the response
	responseBuilder := responsebuilder.APIResponse{
		Ctx:          	ctx,
		HTTPStatus:   	statusCode,
		ErrorResponse: 	errRes,
		Response:     	response,
		RequestID:    	requestID,
	}
	responseBuilder.BuildAPIResponse()

	return nil
}

func UpdateTaskHandler(ctx *fiber.Ctx) error {
	var requestID = web.GetRequestID(ctx)
	commonLogFields := []zap.Field{zap.String(constant.TraceMsgReqID, requestID)}
	utils.Logger.Info(utils.TraceMsgFuncStart(UpdateTaskHandlerMethod), commonLogFields...)

	defer utils.Logger.Info(utils.TraceMsgFuncEnd(UpdateTaskHandlerMethod), commonLogFields...)

	var (
		statusCode  int
		errorResult *custom.ErrorResult
		errRes      custom.ErrorResult
		request	 	dto.UpdateTaskRequest
		response    *dto.ManageTaskResponse
		taskService = service.CreateTaskSerivce(requestID)
	)

	// validate
	utils.Logger.Debug(utils.TraceMsgBeforeInvoke(validator.ValidateUpdateTaskMethod), commonLogFields...)
	request, errorResult = validator.ValidateUpdateTask(requestID, ctx)
	utils.Logger.Debug(utils.TraceMsgAfterInvoke(validator.ValidateUpdateTaskMethod), commonLogFields...)

	request.UserID = ctx.Locals(TokenClaims).(*dto.JWTClaims).UserID

	if errorResult == nil {
		response, errorResult = taskService.UpdateTask(request)
	}

	if errorResult != nil {
		logFields := append(commonLogFields, zap.Any(constant.ErrorNote, errorResult))
		utils.Logger.Error(utils.TraceMsgErrorOccurredFrom(service.UpdateTaskMethod), logFields...)

		statusCode, errRes = HandleError(errorResult)
	}

	// Build the response
	responseBuilder := responsebuilder.APIResponse{
		Ctx:          	ctx,
		HTTPStatus:   	statusCode,
		ErrorResponse: 	errRes,
		Response:     	response,
		RequestID:    	requestID,
	}
	responseBuilder.BuildAPIResponse()

	return nil
}

func DeleteTaskHandler(ctx *fiber.Ctx) error {
	var requestID = web.GetRequestID(ctx)
	commonLogFields := []zap.Field{zap.String(constant.TraceMsgReqID, requestID)}
	utils.Logger.Info(utils.TraceMsgFuncStart(DeleteTaskHandlerMethod), commonLogFields...)

	defer utils.Logger.Info(utils.TraceMsgFuncEnd(DeleteTaskHandlerMethod), commonLogFields...)

	var (
		statusCode  int
		errorResult *custom.ErrorResult
		errRes      custom.ErrorResult
		request	 	dto.DeleteTaskRequest
		response    *dto.ManageTaskResponse
		taskService = service.CreateTaskSerivce(requestID)
	)

	// validate
	utils.Logger.Debug(utils.TraceMsgBeforeInvoke(validator.ValidateDeleteTaskMethod), commonLogFields...)
	request, errorResult = validator.ValidateDeleteTask(requestID, ctx)
	utils.Logger.Debug(utils.TraceMsgAfterInvoke(validator.ValidateDeleteTaskMethod), commonLogFields...)

	request.UserID = ctx.Locals(TokenClaims).(*dto.JWTClaims).UserID

	if errorResult == nil {
		response, errorResult = taskService.DeleteTask(request)
	}

	if errorResult != nil {
		logFields := append(commonLogFields, zap.Any(constant.ErrorNote, errorResult))
		utils.Logger.Error(utils.TraceMsgErrorOccurredFrom(service.DeleteTaskMethod), logFields...)

		statusCode, errRes = HandleError(errorResult)
	}

	// Build the response
	responseBuilder := responsebuilder.APIResponse{
		Ctx:          	ctx,
		HTTPStatus:   	statusCode,
		ErrorResponse: 	errRes,
		Response:     	response,
		RequestID:    	requestID,
	}
	responseBuilder.BuildAPIResponse()

	return nil
}

func GetTasksByUserIDHandler(ctx *fiber.Ctx) error {
	var requestID = web.GetRequestID(ctx)
	commonLogFields := []zap.Field{zap.String(constant.TraceMsgReqID, requestID)}
	utils.Logger.Info(utils.TraceMsgFuncStart(GetTasksByUserIDHandlerMethod), commonLogFields...)

	defer utils.Logger.Info(utils.TraceMsgFuncEnd(GetTasksByUserIDHandlerMethod), commonLogFields...)

	var (
		statusCode  int
		errorResult *custom.ErrorResult
		errRes      custom.ErrorResult
		response    dto.UserTasksResponse
		taskService = service.CreateTaskSerivce(requestID)
	)

	userID := ctx.Locals(TokenClaims).(*dto.JWTClaims).UserID

	response, errorResult = taskService.GetTaskList(userID)

	if errorResult != nil {
		logFields := append(commonLogFields, zap.Any(constant.ErrorNote, errorResult))
		utils.Logger.Error(utils.TraceMsgErrorOccurredFrom(service.GetTaskListMethod), logFields...)

		statusCode, errRes = HandleError(errorResult)
	}

	// Build the response
	responseBuilder := responsebuilder.APIResponse{
		Ctx:          	ctx,
		HTTPStatus:   	statusCode,
		ErrorResponse: 	errRes,
		Response:     	response,
		RequestID:    	requestID,
	}
	responseBuilder.BuildAPIResponse()

	return nil
}