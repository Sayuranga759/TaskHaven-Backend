package repository

import (
	"github.com/Sayuranga759/TaskHaven-Backend/app/routes/dto"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/config/dbconfig"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils/constant"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TaskRepository interface {
	AddTask(task *dto.Tasks) (*dto.Tasks, error)
	UpdateTask(task *dto.Tasks) (*dto.Tasks, error)
	IsTaskExistforUser(taskID, userID uint) (bool, error)
}

type taskRepository struct {
	_                 struct{}
	repositoryContext RepositoryContext
	db                *gorm.DB
}

func CreateTaskRepository(requestID string, transaction *gorm.DB) TaskRepository {
	return &taskRepository{
		repositoryContext: CreateRepositoryContext(requestID, transaction),
		db:                dbconfig.GetDBConnection(),
	}
}

func (repo *taskRepository) getTransaction() *gorm.DB {
	return repo.repositoryContext.Transaction
}

func (repo *taskRepository) AddTask(task *dto.Tasks) (*dto.Tasks, error) {
	commonLogFields := []zap.Field{zap.String(constant.TraceMsgReqID, repo.repositoryContext.RequestID)}
	utils.Logger.Debug(utils.TraceMsgFuncStart(AddTaskMethod), commonLogFields...)
	defer utils.Logger.Debug(utils.TraceMsgFuncEnd(AddTaskMethod), commonLogFields...)

	if err := repo.getTransaction().Create(task).Error; err != nil {
		logFields := append(commonLogFields, zap.Any(Tasks, task), zap.Error(err))
		utils.Logger.Error(utils.TraceMsgErrorOccurredWhenInserting(Tasks), logFields...)
		return task, err
	}

	return task, nil
}

func (repo *taskRepository) UpdateTask(task *dto.Tasks) (*dto.Tasks, error) {
	commonLogFields := []zap.Field{zap.String(constant.TraceMsgReqID, repo.repositoryContext.RequestID)}
	utils.Logger.Debug(utils.TraceMsgFuncStart(UpdateTaskMethod), commonLogFields...)
	defer utils.Logger.Debug(utils.TraceMsgFuncEnd(UpdateTaskMethod), commonLogFields...)

	if err := repo.getTransaction().Save(task).Error; err != nil {
		logFields := append(commonLogFields, zap.Any(Tasks, task), zap.Error(err))
		utils.Logger.Error(utils.TraceMsgErrorOccurredWhenUpdating(Tasks), logFields...)
		return task, err
	}

	return task, nil
}

func (repo *taskRepository) IsTaskExistforUser(taskID, userID uint) (bool, error) {
	commonLogFields := []zap.Field{zap.String(constant.TraceMsgReqID, repo.repositoryContext.RequestID)}
	utils.Logger.Debug(utils.TraceMsgFuncStart(IsTaskExistforUserMethod), commonLogFields...)
	defer utils.Logger.Debug(utils.TraceMsgFuncEnd(IsTaskExistforUserMethod), commonLogFields...)

	var count int64
	if err := repo.db.Model(&dto.Tasks{}).Where(IfTaskIdAndUserIdEqual, taskID, userID).Count(&count).Error; err != nil {
		logFields := append(commonLogFields, zap.Uint(TaskID, taskID), zap.Uint(UserID, userID), zap.Error(err))
		utils.Logger.Error(utils.TraceMsgErrorOccurredWhenSelecting(Tasks), logFields...)
		return false, err
	}

	return count > 0, nil
}


