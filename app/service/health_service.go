package service

import (
	"fmt"

	"github.com/Sayuranga759/TaskHaven-Backend/pkg/custom"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils/constant"
	"go.uber.org/zap"
)

type HealthService struct {
	_              struct{}
	serviceContext ServiceContext
}

func CreateHealthService(requestID string) HealthService {
	return HealthService{serviceContext: CreateServiceContext(requestID)}
}

func (service *HealthService) ReadyzService() (isReady bool, errResult *custom.ErrorResult) {
	commonLogFields := []zap.Field{zap.String(constant.TraceMsgReqID, service.serviceContext.RequestID)}
	utils.Logger.Debug(fmt.Sprintf(constant.TraceMsgFuncStart, ReadyzServiceMethod), commonLogFields...)

	defer func() {
		utils.Logger.Debug(fmt.Sprintf(constant.TraceMsgFuncEnd, ReadyzServiceMethod), commonLogFields...)
	}()

	isReady = true

	return isReady, nil
}
