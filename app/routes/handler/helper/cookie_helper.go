package helper

import (
	"time"

	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils/constant"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func BuildCookie(requestID, tokenString string, ctx *fiber.Ctx) {
	commonLogFields := []zap.Field{zap.String(constant.TraceMsgReqID, requestID)}
	utils.Logger.Debug(utils.TraceMsgFuncStart(BuildCookieMethod), commonLogFields...)
	defer utils.Logger.Debug(utils.TraceMsgFuncEnd(BuildCookieMethod), commonLogFields...)

	cookie := new(fiber.Cookie)
	cookie.Name 	= constant.CookieName
	cookie.Value 	= tokenString
	cookie.Path 	= constant.CookiePath
	cookie.HTTPOnly = true
	cookie.MaxAge 	= int(time.Duration(constant.IntTwentyFour) * time.Hour / time.Second)

	ctx.Cookie(cookie)
}