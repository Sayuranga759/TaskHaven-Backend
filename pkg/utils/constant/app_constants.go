package constant

const (
	Issuer 				= "TaskHaven"
	DatadogTracingSink 	= "datadog"
)

// special charactes
const (
	Colon              = ":"
	Basepath           = "./"
	Empty              = ""
	Hyphen             = "-"
	Dot                = "."
	Slash              = "/"
	Space              = " "
	EqualsQuestionMark = " = ?"
	Zero               = 0
	True               = true
	False              = false
)

const (
	IntOne   		int = 1
	IntTwo   		int = 2
	IntThree 		int = 3
	IntTwentyFour 	int = 24
)

// file constant
const (
	DocumentHTML  = "document.html"
	StaticHTML    = "static.html"
	Doc           = "doc"
	Static        = "static"
	HTML          = "html"
	DotJSON       = ".json"
	DotGoldenJSON = ".golden.json"
)

// token const
const (
	Access  = "access"
	Refresh = "refresh"
	Code    = "code"
)

// cookie const
const (
	CookieName = "jwt"
	CookiePath = "/"
)

// Logger Message
const (
	TraceMsgBeforeFetching        = "Before Fetching %v"
	TraceMsgAfterFetching         = "After Fetching %v"
	TraceMsgBeforeInsertion       = "Before Creating %v"
	TraceMsgAfterInsertion        = "After Creating %v"
	TraceMsgBeforeUpdate          = "Before Update %v"
	TraceMsgAfterUpdate           = "After Update %v"
	TraceMsgFuncEnd               = "%v End here"
	TraceMsgFuncStart             = "%v Start here"
	TraceMsgRequestInitiated      = "%v request initiated"
	TraceMsgReqID                 = "Request Id"
	TraceMsgReqBody               = "Request Body"
	TraceMsgRequestHeader         = "Request Header"
	TranceMsgResponse             = "Response"
	TraceMsgBeforeInvoke          = "Before Call %v"
	TraceMsgAfterInvoke           = "After Call %v"
	TraceMsgBeforeRollback        = "Before rollback from %s"
	TraceMsgAfterRollback         = "After rollback from %s"
	TraceMsgBeforeCommit          = "Before commit from %s"
	TraceMsgAfterCommit           = "After commit from %s"
	TraceMsgAPIResponse           = "Build API Response"
	TraceMsgResponseDetails       = "Response Details"
	TraceMsgAPISuccess            = "Success Response"
	TraceMsgAPIErrorResponse      = "Error Response"
	TraceMsgErrorDetails          = "Error Details"
	MethodInput                   = "Method Input %v"
	Result                        = "Result"
	DebugNote                     = "Debug workflow"
	ErrorNote                     = "Error Note"
	HTMLPassErr                   = "HTML Template pass Error"
	Response                      = "Response"
	ConvertingStrToUint           = "Converting string to uint"
	EmptyHeaderDetails            = "%s does not exist in the request header"
	InactiveChangePwdRecord       = "Inactive change password record"
	InvalidKeyForChangePwd        = "Invalid key for change password"
	InvalidTokenForChangePwd      = "Invalid token for change password"
	MethodOutput                  = "Method Output"
	MethodError                   = "Method Error"
	TraceMsgRespBody              = "Response Body"
	TraceMsgBeforeParse           = "Trace before parse in %s"
	TraceMsgAfterParse            = "Trace after parse in %s"
	TraceRequestType              = "Request Type %T"
	Request                       = "Request"
)

// utils func
const (
	Dial = "Dial"
)

type HTTPMethod string

const (
	Get, Post, Patch, Delete HTTPMethod = "GET", "POST", "PATCH", "DELETE"
)

// Header fields
const (
	XForwardedForHeader = "X-Forwarded-For"
	UserAgentHeader     = "User-Agent"
	PlatformHeader      = "Sec-Ch-Ua-Platform"
	BrowserHeader       = "Sec-Ch-Ua"
)

// Method names
const (
	GetCurrentTimeMethod = "GetCurrentTime"
	JSONUnmarshalMethod  = "JsonUnmarshal"
	StructCasterMethod   = "StructCaster"
	QueryParser          = "QueryParser"
)

