package dbconfig

import (
	"fmt"

	"github.com/Sayuranga759/TaskHaven-Backend/pkg/config"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils/constant"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbCon *gorm.DB

func GetDBConnection() *gorm.DB {
	return dbCon
}

func SetDBConnection(db *gorm.DB) {
	dbCon = db
}

// Create a new connection to the database
func InitDBConnection() error {
	dbConfig := config.GetConfig().DBConfig

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort, dbConfig.DBSSLMode)

	// If ISCloudSql is true, overwrite the DSN.
	if dbConfig.ISCloudSQL {
		dsn = fmt.Sprintf("host=/cloudsql/%s user=%s password=%s dbname=%s sslmode=%s",
			dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBSSLMode)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Logger.Error(constant.DBConnectionOpenError, zap.Error(err))
		return err
	}

	dbCon = db

	return nil
}

func InitDBConWithAutoMigrate(dst ...any) error {
	err := InitDBConnection()
	if err != nil {
		utils.Logger.Error(constant.DBConnectionOpenError, zap.Error(err))
		return err
	}

	err = dbCon.AutoMigrate(dst...)
	if err != nil {
		utils.Logger.Error(constant.DBErrorOccurredWhenAutoMigrate, zap.Error(err))
		return err
	}

	return nil
}
