package custom

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils/constant"
)

// ErrorResult used to define error result of the response
type ErrorResult struct {
	ErrorList  []ErrorInfo
	StatusCode int  `json:"StatusCode" example:"400" swaggerignore:"true"`
	IsError    bool `json:"IsError" example:"true" swaggerignore:"true"`
}

// ErrorInfo use to define error information of the ErrorResult
type ErrorInfo struct {
	ErrorCode    string `json:"ErrorCode" example:"ER0001"`
	ErrorMessage string `json:"ErrorMessage" example:"Records not found"`
	ErrorDetail  string `json:"ErrorDetail" example:"XYZ data not available in db"`
}

// BilddErrorInfo used to build error information
func BuildErrorInfo(errCode, errMessage, errDetail string) ErrorInfo {
	return ErrorInfo{
		ErrorCode:    errCode,
		ErrorMessage: errMessage,
		ErrorDetail:  errDetail,
	}
}

// BuildErrResultWithSuccessStatus used to build ErrorResult with success code
func BuildErrResultWithSuccessStatus(errCode, errMessage, errDetail string) ErrorResult {
	errList := []ErrorInfo{BuildErrorInfo(errCode, errMessage, errDetail)}

	return ErrorResult{
		ErrorList:  errList,
		IsError:    false,
		StatusCode: http.StatusOK,
	}
}

// BuildBadReqErrResultWithList used to build ErrorResult with ErrorInfo list and bad request code
func BuildBadReqErrResultWithList(errInfo ...ErrorInfo) ErrorResult {
	return ErrorResult{
		ErrorList:  errInfo,
		IsError:    false,
		StatusCode: http.StatusBadRequest,
	}
}

// BuildBadReqErrResult used to build ErrorResult with bad request code
func BuildBadReqErrResult(errCode, errMessage, errDetail string) ErrorResult {
	errList := []ErrorInfo{BuildErrorInfo(errCode, errMessage, errDetail)}

	return ErrorResult{
		ErrorList:  errList,
		IsError:    false,
		StatusCode: http.StatusBadRequest,
	}
}

// BuildNotFoundErrResult used to build ErrorResult with bad request code
func BuildNotFoundErrResult(errCode, errMessage, errDetail string) ErrorResult {
	errList := []ErrorInfo{BuildErrorInfo(errCode, errMessage, errDetail)}

	return ErrorResult{
		ErrorList:  errList,
		IsError:    false,
		StatusCode: http.StatusNotFound,
	}
}

// BuildErrResultWithSuccessStatus used to build ErrorResult with bad request code
func BuildInternalServerErrResult(errCode, errMessage, errDetail string) ErrorResult {
	errList := []ErrorInfo{BuildErrorInfo(errCode, errMessage, errDetail)}

	return ErrorResult{
		ErrorList:  errList,
		IsError:    false,
		StatusCode: http.StatusInternalServerError,
	}
}

// GetErrorMessage use to retun error message from ErrorList
func GetErrorMessage(errorResult *ErrorResult) string {
	var errMessages []string
	for _, err := range errorResult.ErrorList {
		errMessages = append(errMessages, err.ErrorMessage)
	}

	return strings.Join(errMessages, ",")
}

// BuildForbiddenErrResult used to build ErrorResult with forbidden code
func BuildForbiddenErrResult(errCode, errMessage, errDetail string) ErrorResult {
	errList := []ErrorInfo{BuildErrorInfo(errCode, errMessage, errDetail)}

	return ErrorResult{
		ErrorList:  errList,
		IsError:    false,
		StatusCode: http.StatusForbidden,
	}
}

func BuildPanicErrResult(panicMethod string) *ErrorResult {
	errRes := BuildInternalServerErrResult(constant.UnexpectedErrorCode, fmt.Sprintf(constant.UnexpectedErrorMessage, panicMethod), "")
	return &errRes
}
