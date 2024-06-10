package constant

// error message for log
const (
	PanicOccurred                  = "error or panic occurred with following stacktrace"
	StackTrace                     = "error stacktrace"
	ErrorRequestBody               = "error Request"
	InvalidInputAndPassErr         = "error input provided is invalid & unable to parse"
	ErrorOccurredWhenValidate      = "error occoured when validate struct"
	MissingRequiredField           = "missing Required Field"
	ErrorOccurredFromService       = "missing Occurred from  %s"
	ErrorOccurredFromMethod        = "error Occurred from  Method %s"
	BindingErrorMessage            = "error Occurred when bind the request context"
	ErrorOccurredWhen              = "error Occurred When %s"
	ErrorOccurredWhenSelecting     = "error occurred when selecting %s"
	ErrorOccurredWhenInserting     = "error occurred when inserting %s"
	ErrorOccurredWhenTemplateParse = "error occourred when template parse"
	ErrorOccurredWhenDeleting      = "error occurred when deleting %s"
	ErrorOccurredWhenUpdating      = "error occurred when updating %s"
	ErrorOccurredWhenHashing       = "error occurred when hashing"
	ErrorOccurredWhenHashCompare   = "error occurred when hash compare"
	UnexpectedWhenMarshalError     = "Unexpected Error occurred when Marshal the data"
	UnexpectedPanicErrorOccurred   = "unexpected panic error occurred"
)

const (
	Encoded = "Encoded"
	Decoded = "Decoded"
)

// error codes
const (
	UnexpectedErrorCode            = "0000"
	BindingErrorCode               = "0001"
	MissingRequiredFieldErrorCode  = "0002"
	MissingRequireWithoutFieldCode = "0003"
	MissingRequireWithFieldCode    = "0004"
	MinLengthFieldCode             = "0005"
	MaxLengthFieldCode             = "0006"
	GreaterValueFieldCode          = "0007"
	PatternErrorCode               = "0008"
	TestConnectionFilCode          = "0009"
	// auth error codes
	ErrOccurredWhenGenAccessTokenCode     = "0010"
	ErrOccurredWhenStringSplitToArrayCode = "0011"
	ErrInvalidClientCredentialsCode       = "0012"
	ErrOccurredWhenExtractingToken        = "0013"
	ErrInvalidTokenSignatureCode          = "0014"
	ErrOccurredWhenAccessingToken         = "0015"
	ErrEmptyAuthHeaderCode                = "0016"
	ErrInvalidAuthHeaderCode              = "0017"
	ErrInvalidUserCredentialsCode         = "0019"
	ErrInvalidAuthCodeCode                = "0021"
	ErrCookieNotFoundCode				  = "0025"
	// db error codes
	ErrRecordNotFoundCode   = "0022"
	ErrDatabaseCode         = "0023"
	// email error
	ErrEmailUniqueConstraintViolationCode = "0031"
	ErrBeginTransactionCode = "0036"
	// parse error
	ErrStringToUintParseCode = "0024"
	// decript error codes
	ErrorOccurredWhenDecryptStringCode   = "0029"
	ErrorOccurredWhenHashingPasswordCode = "0030"
	//
	ErrDataMarshalCode                 = "0037"
	ErrDataUnmarshalCode               = "0038"
	ErrInvalidGetUserReqCode           = "0039"
	ErrOccurredWhenParsingReqQueryCode = "0040"
	ErrOccurredWhenSigningJWTTokenCode  = "0041"
	ErrInvalidTokenCode               = "0042"
)

const (
	InvalidRequestErrorMessage   = "Invalid Request validation Error Occurred"
	UnexpectedErrorMessage       = "Unexpected Error occurred at %s"
	UnexpectedWhenUnmarshalError = "Unexpected Error occurred when Unmarshal the data &s with identifier %s"
	UnexpectedFileCreateError    = "Unexpected Error occurred when Create the %s file"
	TestConnectionFilMessage     = "Error Occurred when TestConnection"
	InvalidDataOrFile            = "Invalid Data or File "
	InvalidJSONTestData          = "Invalid Json test data"
	FileReadError                = "Unable to read file %v"
	HTMLTempPassError            = "HTML Template parse error"
	// auth error message
	ErrOccurredWhenGenAccessTokenMsg     = "error occurred when Generate Access Token"
	ErrOccurredWhenStringSplitToArrayMsg = "error occurred When string split to array"
	ErrorOccurredWhenDecodeStringMsg     = "error occurred When Decode String"
	ErrorOccurredWhenPemBlockDecodeMsg   = "error occurred When decode PEM block containing private key"
	ErrorOccurredWhenParsingKeyMsg       = "error occurred when parsing private key"
	ErrorOccurredWhenDecriptStringMsg    = "error occurred When Decript String"
	ErrorOccurredWhenHashingPasswordMsg  = "error occurred When hashing password"
	ErrorOccurredWhenPassingEmptyKeyMsg  = "error occurred when passing empty private key"
	ErrInvalidClientCredentialsMsg       = "invalid Client Credentials"
	ErrInvalidTokenSignatureMsg          = "invalid Token Signature"
	ErrInvalidTokenMsg				   	 = "invalid token"
	ErrInvalidTokenClaimsMsg             = "invalid token claims"
	ErrOccurredWhenAccessingTokenMsg     = "error Occurred when accessing token"
	ErrEmptyAuthHeaderMsg                = "authorization header is required"
	ErrEmptyAuthKeyHeaderMsg             = "key header is required"
	ErrInvalidAuthHeaderMsg              = "invalid authorization header"
	ErrInParsingTokenMsg                 = "error Occurred when parsing token"
	ErrTokenDoesNotHaveRequiredRoleMsg   = "token does not have required role"
	ErrInvalidUserCredentialsMsg         = "Invalid User Credentials"
	ErrInvalidRoleMsg                    = "token should have (at least one of) the following role(s): %s, but it has following role(s): %s"
	ErrStringToUintParseMsg              = "error occurred when converting string to uint"
	ErrInvalidRefreshTokenMsg            = "invalid refresh token"
	ErrInvalidPasswordChangeMsg          = "invalid password change request"
	ErrInvalidPasswordChangeKeyMsg       = "invalid password change request key"
	ErrPasswordChangeRequestExpiredMsg   = "password change request expired"
	ErrOccurredWhenParsingReqQueryMsg    = "error occurred while parsing the request query"
	ErrInvalidAuthCodeMsg                = "invalid authorization code"
	ErrOccurredWhenSigningJWTTokenMsg    = "error occurred when signing jwt token"
	ErrCookieNotFoundMsg				 = "cookie not found"
	// email error message
	ErrEmailUniqueConstraintViolationMsg = "Email already exists"
	// timeonly error message
	TimeFormatErrorMsg     = "Invalid time format. Time should be in the format HH:MM:SS."
	InvalidTypeForTimeOnly = "invalid type %T for TimeOnly"
	WorkTimeErrorMsg       = "Work start time must be less than work end time"
)

// util methods
const (
	SplitAndDecodeMethod = "SplitAndDecode"
)

// db errors
const (
	DBConnectionOpenError          = "Error occurred when opening the database connection"
	DBConnectionCloseError         = "Error occurred when closing the database connection"
	DBInitFailError                = "Failed to initialize database"
	DBConnectionIsNotEstablished   = "database connection is not established"
	DBErrorOccurredWhenAutoMigrate = "error occurred when auto migrate"
)

const (
	ErrRecordNotFoundMsg          = "Record not found"
	ErrDatabaseMsg                = "Database error"
	ErrBeginTransactionMsg        = "Error occurred when begin transaction"
	ErrorOccurredWhenRetrieveData = "error occurred when retrieving data from %s: %s"
	ErrNotFound                   = "%s not found : %s"
)

// auth server errors
const (
	ServerInitFailError = "Failed to initialize TaskHaven server"
)

// jwt errors
const (
	ErrInvalidSigningMethod                = "invalid signing method"
	ErrOccurredWhenGettingJwtSigningMethod = "error occurred when getting jwt signing method"
	ErrOccurredWhenSigningJWTToken         = "error occurred when signing jwt token"
)
