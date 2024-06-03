package handler

import (
	"net/http"

	"github.com/Sayuranga759/TaskHaven-Backend/app/service"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/custom"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/web"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/web/responsebuilder"
	"github.com/gofiber/fiber/v2"
)


func Lives(ctx *fiber.Ctx) (err error) {
	var (
		requestID = web.GetRequestID(ctx)
	)

	responseBuilder := responsebuilder.APIResponse{
		Ctx:        ctx,
		HTTPStatus: fiber.StatusOK,
		Response:   true,
		RequestID:  requestID,
	}
	responseBuilder.BuildAPIResponse()

	return nil
}


func Readyz(ctx *fiber.Ctx) (err error) {
	var (
		errorResult custom.ErrorResult
		statusCode  = http.StatusOK
		requestID   = web.GetRequestID(ctx)
	)

	service := service.CreateHealthService(requestID)

	isReady, errRes := service.ReadyzService()
	if errRes != nil {
		statusCode = errRes.StatusCode
		errRes.IsError = true
		errorResult = *errRes
	}

	responseBuilder := responsebuilder.APIResponse{
		Ctx:           ctx,
		HTTPStatus:    statusCode,
		ErrorResponse: errorResult,
		Response:      isReady,
		RequestID:     requestID,
	}
	responseBuilder.BuildAPIResponse()

	return nil
}