package validator

import (
	"fmt"

	"github.com/Sayuranga759/TaskHaven-Backend/pkg/custom"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils"
	"github.com/Sayuranga759/TaskHaven-Backend/pkg/utils/constant"
	"github.com/go-playground/locales/en_US"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/en"
	"go.uber.org/zap"
)

var (
	validate         *validator.Validate
	trans            ut.Translator
	generalErrorCode map[string]string
)

// InitValidator used to initiate go playground validator
func InitValidator() {
	validate = validator.New()
	RegisterTagName()
	trans, _ = SetTransLatorForStructError(validate)

	RegisterCustomValidation(validate)
	RegisterCustomTranslation(validate, trans)

	generalErrorCode = BuildGeneralErrorCode()
}

// BuildValidationErrorResponse used to build go playground validator error responses
func BuildValidationErrorResponse(requestID string, validationError error) *custom.ErrorResult {
	commonLogFields := []zap.Field{zap.String(constant.TraceMsgReqID, requestID)}
	utils.Logger.Debug(fmt.Sprintf(constant.TraceMsgFuncStart, BuildValidationErrorResponseMethod), commonLogFields...)

	defer utils.Logger.Info(utils.TraceMsgFuncEnd(BuildValidationErrorResponseMethod), commonLogFields...)

	if validationError != nil {
		errorList := []custom.ErrorInfo{}
		for _, validationErrorsTranslation := range validationError.(validator.ValidationErrors) {
			errorList = append(errorList, custom.BuildErrorInfo(generalErrorCode[validationErrorsTranslation.Tag()], validationErrorsTranslation.Translate(trans), ""))
		}

		err := custom.BuildBadReqErrResultWithList(errorList...)

		return &err
	}

	return nil
}

// SetTransLatorForStructError used to set the translator for the struct error
func SetTransLatorForStructError(validate *validator.Validate) (ut.Translator, error) {
	uni := ut.New(en_US.New())
	translator, _ := uni.GetTranslator("en_US")
	validationErr := translations.RegisterDefaultTranslations(validate, translator)

	return translator, validationErr
}

// ValidateRequest used to struct validation
func ValidateRequest(requestID string, request any) *custom.ErrorResult {
	commonLogFields := []zap.Field{zap.String(constant.TraceMsgReqID, requestID)}
	utils.Logger.Debug(fmt.Sprintf(constant.TraceMsgFuncStart, ValidateRequestMethod), commonLogFields...)
	defer utils.Logger.Debug(fmt.Sprintf(constant.TraceMsgFuncEnd, ValidateRequestMethod), commonLogFields...)

	err := validate.Struct(request)
	if err != nil {
		utils.Logger.Error(constant.ErrorOccurredWhenValidate, append(commonLogFields, zap.Any(constant.ErrorRequestBody, request), zap.Error(err))...)
		return BuildValidationErrorResponse(requestID, err)
	}

	return nil
}

// BuildGeneralErrorCode used to validate general error code
func BuildGeneralErrorCode() map[string]string {
	commonErrorMap := make(map[string]string)
	commonErrorMap[required] = constant.MissingRequiredFieldErrorCode
	commonErrorMap[requiredWithout] = constant.MissingRequireWithoutFieldCode
	commonErrorMap[requiredWith] = constant.MissingRequireWithFieldCode
	commonErrorMap[min] = constant.MinLengthFieldCode
	commonErrorMap[max] = constant.MaxLengthFieldCode
	commonErrorMap[alpha] = constant.PatternErrorCode
	commonErrorMap[alphaNumeric] = constant.PatternErrorCode
	commonErrorMap[email] = constant.PatternErrorCode
	commonErrorMap[password] = constant.PatternErrorCode
	commonErrorMap[timestamp] = constant.PatternErrorCode
	commonErrorMap[alphaNumericWithHyphen] = constant.PatternErrorCode

	return commonErrorMap
}
