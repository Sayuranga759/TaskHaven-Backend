package middleware

import (
	"github.com/Sayuranga759/TaskHaven-Backend/app/routes/dto"
	"github.com/Sayuranga759/TaskHaven-Backend/app/routes/handler"
	"github.com/Sayuranga759/TaskHaven-Backend/app/service"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/custom"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils/constant"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/web"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/web/responsebuilder"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func TokenValidateMiddleware(ctx *fiber.Ctx) error {
	var requestID = web.GetRequestID(ctx)
	commonLogFields := []zap.Field{zap.String(constant.TraceMsgReqID, requestID)}
	utils.Logger.Info(utils.TraceMsgFuncStart(TokenValidateMiddlewareMethod), commonLogFields...)
	defer utils.Logger.Info(utils.TraceMsgFuncEnd(TokenValidateMiddlewareMethod), commonLogFields...)

	var (
		errRes       *custom.ErrorResult
		response     *dto.JWTClaims
		tokenService = service.CreateTokenSerivce(requestID)
	)

	cookie := ctx.Cookies(constant.CookieName)

	request := dto.ValidateTokenRequest{
		Cookie: cookie,
	}

	response, errRes = tokenService.ValidateToken(request)
	if errRes != nil {
		logFields := append(commonLogFields, zap.Any(constant.ErrorNote, errRes))
		utils.Logger.Error(utils.TraceMsgErrorOccurredFrom(service.ValidateTokenMethod), logFields...)

		statusCode, errRes := handler.HandleError(errRes)

		responseBuilder := responsebuilder.APIResponse{
			Ctx:          	ctx,
			HTTPStatus:   	statusCode,
			ErrorResponse: 	errRes,
			Response:     	nil,
			RequestID:    	requestID,
		}
		responseBuilder.BuildAPIResponse()
		return nil
	}

	ctx.Locals(tokenClaims, response)
	return ctx.Next()
}