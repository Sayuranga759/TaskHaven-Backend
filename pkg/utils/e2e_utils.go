package utils

import (
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils/constant"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func HashingPassword(commonLogFields []zap.Field, data string, hashingCost int) (*string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(data), hashingCost)
	if err != nil {
		logFields := append(commonLogFields, zap.Any(constant.ErrorNote, err))
		Logger.Error(constant.ErrorOccurredWhenHashing, logFields...)

		return nil, err
	}

	passwordHashString := string(passwordHash)

	return &passwordHashString, nil
}

func CompareHashingPassword(commonLogFields []zap.Field, requestPwd, hashedPwd string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(requestPwd))
	if err != nil {
		logFields := append(commonLogFields, zap.Any(constant.ErrorNote, err))
		Logger.Error(constant.ErrorOccurredWhenHashCompare, logFields...)

		return false, err
	}

	return true, nil
}

