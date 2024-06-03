package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Sayuranga759/TaskHaven-Backend/pkg/custom"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils/constant"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type customString string

func NewString(s string) *customString {
	str := customString(s)
	return &str
}

func (s *customString) SplitAndDecode(commonLogFields []zap.Field, sep string) ([]string, error) {
	encodedStr := string(*s)

	// Decode the base64 encoded credentials
	decodedCreds, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		logFields := append(commonLogFields, zap.String(constant.Encoded, encodedStr), zap.Error(err))
		Logger.Error(constant.ErrorOccurredWhenDecodeStringMsg, logFields...)

		return nil, err
	}

	// Split the credentials into clientID and clientSecret
	list := strings.Split(string(decodedCreds), sep)

	return list, nil
}

func (s *customString) DecodeString(commonLogFields []zap.Field) (string, error) {
	encodedStr := string(*s)

	// Decode the base64 encoded credentials
	decodedString, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		logFields := append(commonLogFields, zap.String(constant.Encoded, encodedStr), zap.Error(err))
		Logger.Error(constant.ErrorOccurredWhenDecodeStringMsg, logFields...)

		return constant.Empty, err
	}

	return string(decodedString), nil
}

func (s *customString) ConvertStringToUint(commonLogFields []zap.Field) (*uint, *custom.ErrorResult) {
	str := string(*s)

	userID, errCon := strconv.ParseUint(str, 10, 64)
	if errCon != nil {
		logFields := append(commonLogFields, zap.Error(errCon))
		Logger.Error(fmt.Sprintf(constant.ErrorOccurredWhen, constant.ErrStringToUintParseMsg), logFields...)

		errResult := custom.BuildBadReqErrResult(constant.ErrStringToUintParseCode, constant.ErrStringToUintParseMsg, errCon.Error())

		return nil, &errResult
	}

	id := uint(userID)

	return &id, nil
}

func GetCurrentTime(commonLogFields []zap.Field) time.Time {
	Logger.Debug(TraceMsgFuncStart(constant.GetCurrentTimeMethod), commonLogFields...)
	defer Logger.Debug(TraceMsgFuncEnd(constant.GetCurrentTimeMethod), commonLogFields...)

	createdTime := time.Now().UTC()

	return createdTime
}

// JSONUnmarshal is a function that unmarshal a byte array to a struct.
func JSONUnmarshal[T any](commonLogFields []zapcore.Field, data []byte) (T, *custom.ErrorResult) {
	Logger.Debug(fmt.Sprintf(constant.TraceMsgFuncStart, constant.JSONUnmarshalMethod), commonLogFields...)
	defer Logger.Debug(fmt.Sprintf(constant.TraceMsgFuncEnd, constant.JSONUnmarshalMethod), commonLogFields...)

	var v T
	err := json.Unmarshal(data, &v)
	if err != nil {
		Logger.Error(constant.UnexpectedWhenUnmarshalError, append(commonLogFields, zap.Any(constant.ErrorNote, err))...)
		errResult := custom.BuildBadReqErrResult(constant.ErrDataUnmarshalCode, constant.UnexpectedWhenUnmarshalError, err.Error())
		errRes := &errResult
		return v, errRes
	}

	return v, nil
}

func StructCaster[T any](commonLogFields []zapcore.Field, d any) (*T, *custom.ErrorResult) {
	Logger.Debug(fmt.Sprintf(constant.TraceMsgFuncStart, constant.StructCasterMethod), commonLogFields...)
	defer Logger.Debug(fmt.Sprintf(constant.TraceMsgFuncEnd, constant.StructCasterMethod), commonLogFields...)

	data, err := json.Marshal(d)
	if err != nil {
		Logger.Error(constant.UnexpectedWhenMarshalError, append(commonLogFields, zap.Any(constant.ErrorNote, err))...)
		errResult := custom.BuildBadReqErrResult(constant.ErrDataMarshalCode, constant.UnexpectedWhenMarshalError, err.Error())
		return nil, &errResult
	}

	r, errRes := JSONUnmarshal[T](commonLogFields, data)
	if errRes != nil {
		logFields := append(commonLogFields, zap.Any(constant.ErrorNote, errRes))
		Logger.Error(TraceMsgErrorOccurredFrom(constant.JSONUnmarshalMethod), logFields...)
		return nil, errRes
	}

	return &r, nil
}
