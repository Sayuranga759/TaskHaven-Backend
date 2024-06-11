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
	AddTask(user *dto.Tasks) (*dto.Tasks, error)
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