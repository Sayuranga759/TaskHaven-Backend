package main

import (
	"github.com/Sayuranga759/TaskHaven-Backend/app/routes"
	"github.com/Sayuranga759/TaskHaven-Backend/app/routes/dto"
	"github.com/Sayuranga759/TaskHaven-Backend/app/routes/handler/validator"

	"github.com/Sayuranga759/TaskHaven-Backend/pkg/config"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/config/appconfig"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/config/dbconfig"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils/constant"
	"go.uber.org/zap"
)

func init() {
	config.InitConfig()

	err := dbconfig.InitDBConWithAutoMigrate(
		&dto.Users{}, &dto.Priorities{}, &dto.Tasks{}, &dto.Tags{}, &dto.TaskTags{},
	)
	if err != nil {
		utils.Logger.Error(constant.DBInitFailError, zap.Error(err))
	}

	validator.InitValidator()
}

func main() {
	appconfig.Start(routes.APIRoutes)
}