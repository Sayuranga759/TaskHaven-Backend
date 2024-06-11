package utils

import (
	"fmt"

	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils/constant"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/custom"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Logger is the global logger instance.
	Logger *zap.Logger
)

// TraceMsgFuncStart returns a formatted string for tracing function start.
func TraceMsgFuncStart(methodName string) string {
	return fmt.Sprintf(constant.TraceMsgFuncStart, methodName)
}

// TraceMsgFuncEnd returns a formatted string for tracing function end.
func TraceMsgFuncEnd(methodName string) string {
	return fmt.Sprintf(constant.TraceMsgFuncEnd, methodName)
}

// TraceMsgBeforeInvoke returns a formatted string for tracing before invoking a method.
func TraceMsgBeforeInvoke(methodName string) string {
	return fmt.Sprintf(constant.TraceMsgBeforeInvoke, methodName)
}

// TraceMsgAfterInvoke returns a formatted string for tracing after invoking a method.
func TraceMsgAfterInvoke(methodName string) string {
	return fmt.Sprintf(constant.TraceMsgAfterInvoke, methodName)
}

// TraceMsgErrorOccurredWhen returns a formatted string for tracing an error occurred when.
func TraceMsgErrorOccurredWhen(where string) string {
	return fmt.Sprintf(constant.ErrorOccurredWhen, where)
}

// TraceMsgErrorOccurredFrom returns a formatted string for tracing an error occurred from a method.
func TraceMsgErrorOccurredFrom(methodName string) string {
	return fmt.Sprintf(constant.ErrorOccurredFromMethod, methodName)
}

// TraceMsgErrorOccurredWhenSelecting returns a formatted string for tracing an error occurred when selecting.
func TraceMsgErrorOccurredWhenSelecting(what string) string {
	return fmt.Sprintf(constant.ErrorOccurredWhenSelecting, what)
}

// TraceMsgErrorOccurredWhenInserting returns a formatted string for tracing an error occurred when inserting.
func TraceMsgErrorOccurredWhenInserting(what string) string {
	return fmt.Sprintf(constant.ErrorOccurredWhenInserting, what)
}

// TraceMsgErrorOccurredWhenDeleting returns a formatted string for tracing an error occurred when deleting.
func TraceMsgErrorOccurredWhenDeleting(what string) string {
	return fmt.Sprintf(constant.ErrorOccurredWhenDeleting, what)
}

// TraceMsgErrorOccurredWhenUpdating returns a formatted string for tracing an error occurred when updating.
func TraceMsgErrorOccurredWhenUpdating(what string) string {
	return fmt.Sprintf(constant.ErrorOccurredWhenUpdating, what)
}

// TraceMsgBeforeRollback returns a formatted string for tracing before rollback.
func TraceMsgBeforeRollback(methodName string) string {
	return fmt.Sprintf(constant.TraceMsgBeforeRollback, methodName)
}

// TraceMsgAfterRollback returns a formatted string for tracing after rollback.
func TraceMsgAfterRollback(methodName string) string {
	return fmt.Sprintf(constant.TraceMsgAfterRollback, methodName)
}

// TraceMsgBeforeCommit returns a formatted string for tracing before commit.
func TraceMsgBeforeCommit(methodName string) string {
	return fmt.Sprintf(constant.TraceMsgBeforeCommit, methodName)
}

// TraceMsgAfterCommit returns a formatted string for tracing after commit.
func TraceMsgAfterCommit(methodName string) string {
	return fmt.Sprintf(constant.TraceMsgAfterCommit, methodName)
}

// TraceMsgBeforeParse returns a formatted string for tracing before parse.
func TraceMsgBeforeParse(methodName string) string {
	return fmt.Sprintf(constant.TraceMsgBeforeParse, methodName)
}

// TraceRequestType is a function that appends the request type to a slice of zapcore.Field.
func TraceRequestType(request any) string {
	return fmt.Sprintf(constant.TraceRequestType, request)
}

// TraceStack is a function that appends a stack trace to a slice of zapcore.Field.
func TraceStack(commonLogFields []zapcore.Field, stack []byte) []zapcore.Field {
	return append(commonLogFields, zap.String(constant.StackTrace, string(stack)))
}

// TraceLogMethodOutputs is a function that appends method inputs to a slice of zapcore.Field.
func TraceLogMethodOutputs(commonLogFields []zapcore.Field, output, err any) []zapcore.Field {
	return append(commonLogFields, zap.Any(constant.MethodOutput, output), zap.Any(constant.MethodError, err))
}

// TraceLogMethodOutputs is a function that appends method inputs to a slice of zapcore.Field.
func TraceLogMethodOutputsWithErr(commonLogFields []zapcore.Field, output any, err error) []zapcore.Field {
	return append(commonLogFields, zap.Any(constant.MethodOutput, output), zap.Error(err))
}

// CommonLogField is a function that appends the request ID to a slice of zap.Field.
func CommonLogField(requestID string, fields ...zapcore.Field) []zap.Field {
	logFields := []zap.Field{zap.String(constant.TraceMsgReqID, requestID)}
	logFields = append(logFields, fields...)
	return logFields
}

// TraceCustomError is a function that appends a custom error to a slice of zapcore.Field.
func TraceCustomError(commonLogFields []zapcore.Field, err custom.ErrorResult) []zapcore.Field {
	return append(commonLogFields, zap.Any(constant.ErrorNote, err))
}

// TraceError is a function that appends an error to a slice of zapcore.Field.
func TraceError(commonLogFields []zapcore.Field, err error) []zapcore.Field {
	return append(commonLogFields, zap.Error(err))
}
