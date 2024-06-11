package service

import (
	"fmt"
	"runtime/debug"

	"github.com/Sayuranga759/TaskHaven-Backend/app/repository"
	"github.com/Sayuranga759/TaskHaven-Backend/app/routes/dto"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/config"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/custom"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils/constant"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type UserService struct {
	_ struct{}
	ServiceContext ServiceContext
	transaction    *gorm.DB
	tokenService   *TokenService
	userRepo 	   repository.UserRepository
}

func CreateUserSerivce(requestID string) *UserService {
	return &UserService{
		ServiceContext: CreateServiceContext(requestID),
	}
}

func (service UserService) RegisterUser(request dto.UserRegistrationRequest) (response *dto.UserRegistrationResponse, errResult *custom.ErrorResult) {
	commonLogFields := utils.CommonLogField(service.ServiceContext.RequestID)
	utils.Logger.Debug(utils.TraceMsgFuncStart(RegisterUserMethod), commonLogFields...)

	defer func() {
		// Panic handling
		if r := recover(); r != nil {
			utils.Logger.Error(constant.PanicOccurred, utils.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(RegisterUserMethod)
		}

		errResult = handleTransaction(commonLogFields, service.transaction, errResult, RegisterUserMethod)
		if errResult != nil {
			utils.Logger.Error(utils.TraceMsgErrorOccurredWhen(HandleTransactionMethod), utils.TraceCustomError(commonLogFields, *errResult)...)
		}

		utils.Logger.Debug(utils.TraceMsgFuncEnd(RegisterUserMethod), commonLogFields...)
	}()

	service.transaction, errResult = BeginNewTransaction()
	if errResult != nil {
		utils.Logger.Error(utils.TraceMsgErrorOccurredWhen(BeginNewTransactionMethod), utils.TraceCustomError(commonLogFields, *errResult)...)

		return nil, errResult
	}

	hashedPassword, err := utils.HashingPassword(commonLogFields, request.Password, config.GetConfig().HashingCost)
	if err != nil {
		errRes := custom.BuildInternalServerErrResult(constant.ErrorOccurredWhenHashingPasswordCode, constant.ErrorOccurredWhenHashingPasswordMsg, err.Error())

		return nil, &errRes
	}

	user := dto.Users{
		Name:    request.Name,
		Email:   request.Email,
		Password: *hashedPassword,
	}

	service.userRepo = repository.CreateUserRepository(service.ServiceContext.RequestID, service.transaction)

	addedUser, err := service.userRepo.AddUser(&user)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == SQLStateUniqueViolation {
				handledErr := fmt.Errorf(UniqueConstraintViolation, pgErr.Message)
				errRes := custom.BuildBadReqErrResult(constant.ErrEmailUniqueConstraintViolationCode, constant.ErrEmailUniqueConstraintViolationMsg, handledErr.Error())

				return nil, &errRes
			}
		}
		utils.Logger.Error(utils.TraceMsgErrorOccurredWhen(repository.AddUserMethod), utils.TraceError(commonLogFields, err)...)
		errRes := buildDBError(repository.Users, err)

		return nil, errRes
	}

	response, errResult = utils.StructCaster[dto.UserRegistrationResponse](commonLogFields, addedUser)
	if errResult != nil {
		utils.Logger.Error(utils.TraceMsgErrorOccurredFrom(constant.StructCasterMethod), utils.TraceCustomError(commonLogFields, *errResult)...)
		return nil, errResult
	}

	return response, nil
}

func (service UserService) Login(request dto.LoginRequest, ctx *fiber.Ctx) (response *dto.LoginResponse, errResult *custom.ErrorResult) {
	commonLogFields := utils.CommonLogField(service.ServiceContext.RequestID)
	utils.Logger.Debug(utils.TraceMsgFuncStart(LoginMethod), commonLogFields...)

	defer func() {
		// Panic handling
		if r := recover(); r != nil {
			utils.Logger.Error(constant.PanicOccurred, utils.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(LoginMethod)
		}

		utils.Logger.Debug(utils.TraceMsgFuncEnd(LoginMethod), commonLogFields...)
	}()

	service.tokenService = CreateTokenSerivce(service.ServiceContext.RequestID)
	service.userRepo = repository.CreateUserRepository(service.ServiceContext.RequestID, nil)

	user, err := service.userRepo.GetUserByEmail(request.Email)
	if err != nil {
		logFields := utils.TraceError(commonLogFields, err)
		utils.Logger.Error(utils.TraceMsgErrorOccurredFrom(repository.GetUserByEmailMethod), logFields...)
		errRes := custom.BuildBadReqErrResult(constant.ErrInvalidUserCredentialsCode, constant.ErrInvalidUserCredentialsMsg, constant.Empty)

		return nil, &errRes
	}

	isMatched, err := utils.CompareHashingPassword(commonLogFields, request.Password, user.Password)
	if !isMatched {
		logFields := utils.TraceError(commonLogFields, err)
		utils.Logger.Error(utils.TraceMsgErrorOccurredWhen(constant.ErrorOccurredWhenHashCompare), logFields...)
		errRes := custom.BuildBadReqErrResult(constant.ErrInvalidUserCredentialsCode, constant.ErrInvalidUserCredentialsMsg, constant.Empty)

		return nil, &errRes
	}

	accessToken, errRes := service.tokenService.generateToken(*user)
	if errRes != nil {
		logFields := utils.TraceCustomError(commonLogFields, *errRes)
		utils.Logger.Error(utils.TraceMsgErrorOccurredFrom(GenerateTokenMethod), logFields...)

		return nil, errRes
	}

	response = &dto.LoginResponse{
		AccessToken: *accessToken,
		UserID: user.UserID,
	}

	return response, nil
}




