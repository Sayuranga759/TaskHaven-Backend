package service

import (
	"fmt"
	"time"

	"github.com/Sayuranga759/TaskHaven-Backend/app/routes/dto"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/config"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/custom"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils/constant"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

type TokenService struct {
	_              struct{}
	ServiceContext ServiceContext
}

func CreateTokenSerivce(requestID string) *TokenService {
	return &TokenService{
		ServiceContext: CreateServiceContext(requestID),
	}
}

func (service TokenService) ValidateToken(request dto.ValidateTokenRequest) (response *dto.JWTClaims, errResult *custom.ErrorResult) {
	commonLogFields := utils.CommonLogField(service.ServiceContext.RequestID)
	utils.Logger.Debug(utils.TraceMsgFuncStart(ValidateTokenMethod), commonLogFields...)
	defer utils.Logger.Debug(utils.TraceMsgFuncEnd(ValidateTokenMethod), commonLogFields...)

	if request.Cookie == constant.Empty {
		utils.Logger.Error(constant.ErrCookieNotFoundMsg, commonLogFields...)
		errResult := custom.BuildBadReqErrResult(constant.ErrCookieNotFoundCode, constant.ErrCookieNotFoundMsg, constant.Empty)

		return nil, &errResult
	}

	token, errResult := service.validateTokenSignature(request.Cookie, config.GetConfig().JWTSecret)
	if errResult != nil {
		utils.Logger.Error(constant.ErrInvalidTokenSignatureMsg, commonLogFields...)
		return nil, errResult
	}

	jwtClaims, errResult := service.extractClaimsFromToken(token)
	if errResult != nil {
		utils.Logger.Error(constant.ErrInvalidTokenClaimsMsg, commonLogFields...)
		return nil, errResult
	}

	return jwtClaims, nil
}

func (service TokenService) generateToken(user dto.Users) (accessToken *string, errResult *custom.ErrorResult) {
	commonLogFields := utils.CommonLogField(service.ServiceContext.RequestID)
	utils.Logger.Debug(utils.TraceMsgFuncStart(GenerateTokenMethod), commonLogFields...)
	defer utils.Logger.Debug(utils.TraceMsgFuncEnd(GenerateTokenMethod), commonLogFields...)

	claims := dto.JWTClaims{
		Name:   user.Name,
		Email:  user.Email,
		UserID: user.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(constant.IntOne))),
			Issuer:    constant.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.GetConfig().JWTSecret))
	if err != nil {
		utils.Logger.Error(constant.ErrOccurredWhenSigningJWTTokenMsg, utils.TraceError(commonLogFields, err)...)
		errRes := custom.BuildInternalServerErrResult(constant.ErrOccurredWhenSigningJWTTokenCode, constant.ErrOccurredWhenSigningJWTTokenMsg, err.Error())

		return nil, &errRes
	}

	return &tokenString, nil
}

func (service TokenService) validateTokenSignature(tokenString, secretKey string) (*jwt.Token, *custom.ErrorResult) {
	commonLogFields := utils.CommonLogField(service.ServiceContext.RequestID)
	utils.Logger.Debug(utils.TraceMsgFuncStart(validateTokenSignatureMethod), commonLogFields...)
	defer utils.Logger.Debug(utils.TraceMsgFuncEnd(validateTokenSignatureMethod), commonLogFields...)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			utils.Logger.Error(constant.ErrInvalidTokenSignatureMsg, commonLogFields...)			
			return nil, fmt.Errorf(constant.ErrInvalidTokenSignatureMsg)
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		logFields := append(commonLogFields, zap.Any(constant.ErrorNote, err))
		utils.Logger.Error(constant.ErrInvalidTokenSignatureMsg, logFields...)
		errResult := custom.BuildBadReqErrResult(constant.ErrInvalidTokenSignatureCode, constant.ErrInvalidTokenSignatureMsg, err.Error())

		return nil, &errResult
	}

	return token, nil
}

func (service TokenService) extractClaimsFromToken(token *jwt.Token) (*dto.JWTClaims, *custom.ErrorResult) {
	commonLogFields := []zap.Field{zap.String(constant.TraceMsgReqID, service.ServiceContext.RequestID)}
	utils.Logger.Debug(utils.TraceMsgFuncStart(extractClaimsFromTokenMethod), commonLogFields...)
	defer utils.Logger.Debug(utils.TraceMsgFuncEnd(extractClaimsFromTokenMethod), commonLogFields...)

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		utils.Logger.Error(constant.ErrInvalidTokenMsg, commonLogFields...)
		errResult := custom.BuildBadReqErrResult(constant.ErrInvalidTokenCode, constant.ErrInvalidTokenMsg, constant.Empty)

		return nil, &errResult
	}

	jwtClaims := dto.JWTClaims{
		UserID: uint(claims[UserID].(float64)),
		Name:   claims[Name].(string),
		Email:  claims[Email].(string),
	}

	return &jwtClaims, nil
}
