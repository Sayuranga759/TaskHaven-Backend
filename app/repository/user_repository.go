package repository

import (
	"github.com/Sayuranga759/TaskHaven-Backend/app/routes/dto"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/config/dbconfig"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils/constant"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository interface {
	AddUser(user *dto.User) (*dto.User, error)
}

type userRepository struct {
	_ 				  struct{}
	repositoryContext RepositoryContext
	db 			      *gorm.DB
}

func CreateUserRepository(requestID string, transaction *gorm.DB) UserRepository {
	return &userRepository{
		repositoryContext: CreateRepositoryContext(requestID, transaction),
		db: dbconfig.GetDBConnection(),
	}
}

func (repo *userRepository) getTransaction() *gorm.DB {
	return repo.repositoryContext.Transaction
}

func (repo *userRepository) AddUser(user *dto.User) (*dto.User, error) {
	commonLogFields := []zap.Field{zap.String(constant.TraceMsgReqID, repo.repositoryContext.RequestID)}
	utils.Logger.Debug(utils.TraceMsgFuncStart(AddUserMethod), commonLogFields...)
	defer utils.Logger.Debug(utils.TraceMsgFuncEnd(AddUserMethod), commonLogFields...)

	if err := repo.getTransaction().Create(user).Error; err != nil {
		logFields := append(commonLogFields, zap.Any(User, user), zap.Error(err))
		utils.Logger.Error(utils.TraceMsgErrorOccurredWhenInserting(User), logFields...)
		return user, err
	}

	return user, nil
}