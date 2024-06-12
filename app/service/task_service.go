package service

import (
	"runtime/debug"

	"github.com/Sayuranga759/TaskHaven-Backend/app/repository"
	"github.com/Sayuranga759/TaskHaven-Backend/app/routes/dto"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/custom"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils/constant"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TaskService struct {
	_              struct{}
	ServiceContext ServiceContext
	transaction    *gorm.DB
	taskRepo       repository.TaskRepository
}

func CreateTaskSerivce(requestID string) *TaskService {
	return &TaskService{
		ServiceContext: CreateServiceContext(requestID),
	}
}

func (service TaskService) CreateTask(request dto.ManageTaskRequest) (response *dto.ManageTaskResponse, errResult *custom.ErrorResult) {
	commonLogFields := utils.CommonLogField(service.ServiceContext.RequestID)
	utils.Logger.Debug(utils.TraceMsgFuncStart(CreateTaskMethod), commonLogFields...)

	defer func() {
		// Panic handling
		if r := recover(); r != nil {
			utils.Logger.Error(constant.PanicOccurred, utils.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(CreateTaskMethod)
		}

		errResult = handleTransaction(commonLogFields, service.transaction, errResult, CreateTaskMethod)
		if errResult != nil {
			utils.Logger.Error(utils.TraceMsgErrorOccurredWhen(HandleTransactionMethod), utils.TraceCustomError(commonLogFields, *errResult)...)
		}

		utils.Logger.Debug(utils.TraceMsgFuncEnd(CreateTaskMethod), commonLogFields...)
	}()

	service.transaction, errResult = BeginNewTransaction()
	if errResult != nil {
		utils.Logger.Error(utils.TraceMsgErrorOccurredWhen(BeginNewTransactionMethod), utils.TraceCustomError(commonLogFields, *errResult)...)

		return nil, errResult
	}

	service.taskRepo = repository.CreateTaskRepository(service.ServiceContext.RequestID, service.transaction)

	task := dto.Tasks{
		UserID: request.UserID,
		PriorityID: request.PriorityID,
		Title: request.Title,
		Description: request.Description,
		Status: request.Status,
		DueDate: request.DueDate,																										
	}

	addedTask, err := service.taskRepo.AddTask(&task)
	if err != nil {
		utils.Logger.Error(utils.TraceMsgErrorOccurredFrom(CreateTaskMethod), append(commonLogFields, zap.Any(constant.ErrorNote, err))...)
		errRes := custom.BuildInternalServerErrResult(constant.ErrDatabaseCode, constant.ErrDatabaseMsg, err.Error())

		return nil, &errRes
	}

	response, errResult = utils.StructCaster[dto.ManageTaskResponse](commonLogFields, addedTask)
	if errResult != nil {
		utils.Logger.Error(utils.TraceMsgErrorOccurredFrom(constant.StructCasterMethod), utils.TraceCustomError(commonLogFields, *errResult)...)
		return nil, errResult
	}

	return response, nil
}

func (service TaskService) UpdateTask(request dto.ManageTaskRequest) (response *dto.ManageTaskResponse, errResult *custom.ErrorResult) {
	commonLogFields := utils.CommonLogField(service.ServiceContext.RequestID)
	utils.Logger.Debug(utils.TraceMsgFuncStart(UpdateTaskMethod), commonLogFields...)

	defer func() {
		// Panic handling
		if r := recover(); r != nil {
			utils.Logger.Error(constant.PanicOccurred, utils.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(UpdateTaskMethod)
		}

		errResult = handleTransaction(commonLogFields, service.transaction, errResult, UpdateTaskMethod)
		if errResult != nil {
			utils.Logger.Error(utils.TraceMsgErrorOccurredWhen(HandleTransactionMethod), utils.TraceCustomError(commonLogFields, *errResult)...)
		}

		utils.Logger.Debug(utils.TraceMsgFuncEnd(UpdateTaskMethod), commonLogFields...)
	}()

	service.transaction, errResult = BeginNewTransaction()
	if errResult != nil {
		utils.Logger.Error(utils.TraceMsgErrorOccurredWhen(BeginNewTransactionMethod), utils.TraceCustomError(commonLogFields, *errResult)...)

		return nil, errResult
	}

	service.taskRepo = repository.CreateTaskRepository(service.ServiceContext.RequestID, service.transaction)

	errRes := service.isTaskExistForUser(request.UserID, request.TaskID)
	if errRes != nil {
		logFields := append(commonLogFields, zap.Any(constant.ErrorNote, errRes))
		utils.Logger.Error(utils.TraceMsgErrorOccurredFrom(UpdateTaskMethod), logFields...)

		return nil, errRes
	}

	task := dto.Tasks{
		UserID: request.UserID,
		PriorityID: request.PriorityID,
		Title: request.Title,
		Description: request.Description,
		Status: request.Status,
		DueDate: request.DueDate,																										
	}

	updatedTask, err := service.taskRepo.UpdateTask(&task)
	if err != nil {
		utils.Logger.Error(utils.TraceMsgErrorOccurredFrom(UpdateTaskMethod), append(commonLogFields, zap.Any(constant.ErrorNote, err))...)
		errRes := custom.BuildInternalServerErrResult(constant.ErrDatabaseCode, constant.ErrDatabaseMsg, err.Error())

		return nil, &errRes
	}

	response, errResult = utils.StructCaster[dto.ManageTaskResponse](commonLogFields, updatedTask)
	if errResult != nil {
		utils.Logger.Error(utils.TraceMsgErrorOccurredFrom(constant.StructCasterMethod), utils.TraceCustomError(commonLogFields, *errResult)...)
		return nil, errResult
	}

	return response, nil
}

func (service TaskService) DeleteTask(request dto.DeleteTaskRequest) (response *dto.ManageTaskResponse, errResult *custom.ErrorResult) {
	commonLogFields := utils.CommonLogField(service.ServiceContext.RequestID)
	utils.Logger.Debug(utils.TraceMsgFuncStart(DeleteTaskMethod), commonLogFields...)

	defer func() {
		// Panic handling
		if r := recover(); r != nil {
			utils.Logger.Error(constant.PanicOccurred, utils.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(DeleteTaskMethod)
		}

		errResult = handleTransaction(commonLogFields, service.transaction, errResult, DeleteTaskMethod)
		if errResult != nil {
			utils.Logger.Error(utils.TraceMsgErrorOccurredWhen(HandleTransactionMethod), utils.TraceCustomError(commonLogFields, *errResult)...)
		}

		utils.Logger.Debug(utils.TraceMsgFuncEnd(DeleteTaskMethod), commonLogFields...)
	}()

	service.transaction, errResult = BeginNewTransaction()
	if errResult != nil {
		utils.Logger.Error(utils.TraceMsgErrorOccurredWhen(BeginNewTransactionMethod), utils.TraceCustomError(commonLogFields, *errResult)...)

		return nil, errResult
	}

	service.taskRepo = repository.CreateTaskRepository(service.ServiceContext.RequestID, service.transaction)

	errRes := service.isTaskExistForUser(request.UserID, request.TaskID)
	if errRes != nil {
		logFields := append(commonLogFields, zap.Any(constant.ErrorNote, errRes))
		utils.Logger.Error(utils.TraceMsgErrorOccurredFrom(DeleteTaskMethod), logFields...)

		return nil, errRes
	}

	deletedTask, err := service.taskRepo.DeleteTask(request.TaskID)
	if err != nil {
		utils.Logger.Error(utils.TraceMsgErrorOccurredFrom(DeleteTaskMethod), append(commonLogFields, zap.Any(constant.ErrorNote, err))...)
		errRes := custom.BuildInternalServerErrResult(constant.ErrDatabaseCode, constant.ErrDatabaseMsg, err.Error())

		return nil, &errRes
	}

	response, errResult = utils.StructCaster[dto.ManageTaskResponse](commonLogFields, deletedTask)
	if errResult != nil {
		utils.Logger.Error(utils.TraceMsgErrorOccurredFrom(constant.StructCasterMethod), utils.TraceCustomError(commonLogFields, *errResult)...)
		return nil, errResult
	}

	return nil, nil
}

func (service TaskService) isTaskExistForUser(userID, taskID uint) (errResult *custom.ErrorResult) {
	commonLogFields := utils.CommonLogField(service.ServiceContext.RequestID)
	utils.Logger.Debug(utils.TraceMsgFuncStart(isTaskExistforUserMethod), commonLogFields...)

	defer utils.Logger.Debug(utils.TraceMsgFuncEnd(isTaskExistforUserMethod), commonLogFields...)

	service.taskRepo = repository.CreateTaskRepository(service.ServiceContext.RequestID, nil)

	isExist, err := service.taskRepo.IsTaskExistforUser(taskID, userID)
	if err != nil {
		utils.Logger.Error(utils.TraceMsgErrorOccurredFrom(isTaskExistforUserMethod), append(commonLogFields, zap.Any(constant.ErrorNote, err))...)
		errRes := custom.BuildInternalServerErrResult(constant.ErrDatabaseCode, constant.ErrDatabaseMsg, err.Error())

		return &errRes
	}

	if !isExist {
		utils.Logger.Error(utils.TraceMsgErrorOccurredFrom(isTaskExistforUserMethod), append(commonLogFields, zap.Any(constant.ErrorNote, constant.ErrUserDoNotHaveAccessMsg))...)
		errRes := custom.BuildNotFoundErrResult(constant.ErrUserDoNotHaveAccessCode, constant.ErrUserDoNotHaveAccessMsg, constant.Empty)
		return &errRes
	}

	return nil
}