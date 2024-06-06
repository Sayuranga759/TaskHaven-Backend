package service

import (
	"fmt"
	"runtime/debug"

	"github.com/Sayuranga759/TaskHaven-Backend/pkg/config/dbconfig"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/custom"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils/constant"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func BeginNewTransaction() (transaction *gorm.DB, errResult *custom.ErrorResult) {
	utils.Logger.Debug(utils.TraceMsgFuncStart(BeginNewTransactionMethod))
	defer utils.Logger.Debug(utils.TraceMsgFuncEnd(BeginNewTransactionMethod))

	defer func() {
		// Panic handling
		if r := recover(); r != nil {
			utils.Logger.Error(constant.PanicOccurred, zap.String(constant.StackTrace, string(debug.Stack())))
			errRes := custom.BuildInternalServerErrResult(
				constant.UnexpectedErrorCode,
				fmt.Sprintf(constant.UnexpectedErrorMessage, BeginNewTransactionMethod),
				constant.Empty)

			errResult = &errRes
		}

		utils.Logger.Debug(utils.TraceMsgFuncEnd(BeginNewTransactionMethod))
	}()

	transaction = dbconfig.GetDBConnection().Begin()
	if transaction.Error != nil {
		utils.Logger.Error(constant.ErrBeginTransactionMsg, zap.Error(transaction.Error))
		errRes := custom.BuildInternalServerErrResult(
			constant.ErrBeginTransactionCode,
			constant.ErrBeginTransactionMsg,
			transaction.Error.Error())

		return nil, &errRes
	}

	return transaction, nil
}

func handleTransaction(commonLogFields []zap.Field, transaction *gorm.DB, errRes *custom.ErrorResult, callingMethod string) (errResult *custom.ErrorResult) {
	utils.Logger.Debug(utils.TraceMsgFuncStart(HandleTransactionMethod), commonLogFields...)

	defer func() {
		// Panic handling
		if r := recover(); r != nil {
			utils.Logger.Error(constant.PanicOccurred, append(commonLogFields, []zap.Field{zap.String(constant.StackTrace, string(debug.Stack()))}...)...)
			errRes := custom.BuildInternalServerErrResult(
				constant.UnexpectedErrorCode,
				fmt.Sprintf(constant.UnexpectedErrorMessage, HandleTransactionMethod),
				constant.Empty)

			errResult = &errRes
		}

		utils.Logger.Debug(utils.TraceMsgFuncEnd(HandleTransactionMethod), commonLogFields...)
	}()

	if transaction == nil {
		utils.Logger.Debug(TransactionNotExist, commonLogFields...)
		return errRes
	}

	// rollback
	if errRes != nil {
		utils.Logger.Debug(utils.TraceMsgBeforeRollback(callingMethod), append(commonLogFields, zap.Any(constant.ErrorNote, errRes))...)
		transaction.Rollback()
		utils.Logger.Debug(utils.TraceMsgAfterRollback(callingMethod), commonLogFields...)

		return errRes
	}

	// commit
	utils.Logger.Debug(utils.TraceMsgBeforeCommit(callingMethod), commonLogFields...)
	transaction.Commit()
	utils.Logger.Debug(utils.TraceMsgAfterCommit(callingMethod), commonLogFields...)

	return nil
}

func buildPanicErr(method string) *custom.ErrorResult {
	errRes := custom.BuildInternalServerErrResult(constant.UnexpectedErrorCode,
		fmt.Sprintf(constant.UnexpectedErrorMessage, method),
		constant.Empty)
	return &errRes
}

func buildDBError(method string, err error) *custom.ErrorResult {
	errRes := custom.BuildInternalServerErrResult(constant.ErrDatabaseCode,
		fmt.Sprintf(constant.ErrorOccurredWhenSelecting, method),
		err.Error())
	return &errRes
}
