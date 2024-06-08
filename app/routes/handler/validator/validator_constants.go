package validator

// Regexs
const (
	alphaNumericRegex                        = "[^A-Za-z0-9]"
	alphaNumericWithHyphenRegex              = "[^A-Za-z0-9-]"
	alphaNumericWithHyphenSpaceRegex         = "[^A-Za-z0-9- ]"
	alphaNumericWithHyphenDotAndAddressRegex = "[^A-Za-z0-9-.@]"
	alphaRegex                               = "[^A-Za-z]"
	phoneNumberWithPlusRegex                 = `^\+([0-9]){11,}$`
	phoneNumberWithoutPlusRegex              = `^[0-9]{10,}$`
	positiveIntegerWithPlusRegex             = `[^+0-9]`
	timestampRegex                           = `^(\d{4})-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])T([01][0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])Z`
)

// Validator keys
const (
	alphaNumeric                = "alphaNumeric"
	alphaNumericWithHyphen      = "alphaNumericWithHyphen"
	alphaNumericWithHyphenSpace = "alphaNumericWithHyphenSpace"
	alpha                       = "alpha"
	required                    = "required"
	requiredWithout             = "required_without"
	requiredWith                = "required_with"
	min                         = "min"
	max                         = "max"
	email                       = "email"
	phone                       = "phone"
	password                    = "password"
	timestamp                   = "timestamp"
	intWithPlus                 = "int_with_plus"
	domain                      = "domain"
	url                         = "url"
	timeonly                    = "timeonly"
	timezone                    = "timezone"
	oneof                       = "oneof"
	oneOfRole                   = "oneOfRole"
)

// Methods
const (
	ValidateUserRegistrationMethod            = "ValidateUserRegistration"
	ValidateLoginMethod					      = "ValidateLogin"
	BuildValidationErrorResponseMethod        = "BuildValidationErrorResponse"
	ValidateRequestMethod                     = "ValidateRequest"
)

// Specific constants for validation
const (
	Plus        = "+"
	JSON        = "json"
	Underscore  = "_"
	EmptyString = ""
)

// constant names
const (
	UserEmail     = "User Email"
	RoleName      = "Role Name"
	refreshToken  = "refresh_token"
	code          = "code"
)
