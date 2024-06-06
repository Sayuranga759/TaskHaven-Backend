package handler

import "github.com/Sayuranga759/TaskHaven-Backend/pkg/custom"

func HandleError(errRes *custom.ErrorResult) (statusCode int, errorResult custom.ErrorResult) {
	errorResult = *errRes
	errorResult.IsError = true

	return errRes.StatusCode, errorResult
}
