package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Sayuranga759/TaskHaven-Backend/pkg/custom"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils/constant"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"
)

func SetupFiber(idleTimeout time.Duration) *fiber.App {
	// log-----
	return fiber.New(fiber.Config{
		//Prefork:               false,
		IdleTimeout:           idleTimeout,
		DisableStartupMessage: false,
		ErrorHandler:          ErrHandler,
	})
}

func ErrHandler(ctx *fiber.Ctx, _ error) error {
	// In case the SendFile fails
	errRes := custom.BuildInternalServerErrResult(constant.UnexpectedErrorCode, constant.UnexpectedPanicErrorOccurred, constant.Empty)
	return ctx.Status(http.StatusInternalServerError).JSON(errRes.ErrorList)
}

func Shutdown(app *fiber.App) error {
	// utils.Logger.Info("Shutting down Fiber...")
	err := app.Shutdown()
	// utils.Logger.Info("Shutdown complete")
	return err
}

func GetRequestID(ctx *fiber.Ctx) string {
	return ctx.Locals(requestid.ConfigDefault.ContextKey).(string)
}

func GetHeaderFromRequest(commonLogFields []zap.Field, ctx *fiber.Ctx, headerKey string) string {
	headerValue := ctx.Get(headerKey)
	if headerValue == constant.Empty {
		utils.Logger.Debug(fmt.Sprintf(constant.EmptyHeaderDetails, headerKey), commonLogFields...)
	}

	return headerValue
}

func QueryParser[T any](ctx *fiber.Ctx, commonLogFields []zap.Field) (*T, *custom.ErrorResult) {
	utils.Logger.Debug(fmt.Sprintf(constant.TraceMsgFuncStart, constant.QueryParser), commonLogFields...)
	defer utils.Logger.Debug(fmt.Sprintf(constant.TraceMsgFuncEnd, constant.QueryParser), commonLogFields...)

	var request T
	err := ctx.QueryParser(&request)
	if err != nil {
		utils.Logger.Error(constant.ErrOccurredWhenParsingReqQueryMsg, append(commonLogFields, zap.Error(err))...)
		errRes := custom.BuildBadReqErrResult(constant.ErrOccurredWhenParsingReqQueryCode,
			constant.ErrOccurredWhenParsingReqQueryMsg, constant.Empty)
		return nil, &errRes
	}

	return &request, nil
}