package service

// Token claims
const (
	UserID 			= "UserID"
	Email 			= "Email"
	Name 			= "Name"
)

const (
	ReadyzServiceMethod = "ReadyzService"
	// user service methods
	RegisterUserMethod = "RegisterUser"
	LoginMethod		   = "Login"
	// token service methods
	GenerateTokenMethod 			= "GenerateToken"
	ValidateTokenMethod 			= "ValidateToken"
	validateTokenSignatureMethod 	= "validateTokenSignature"
	extractClaimsFromTokenMethod 	= "extractClaimsFromToken"
	// task service methods
	CreateTaskMethod 		 = "CreateTask"
	UpdateTaskMethod 		 = "UpdateTask"
	isTaskExistforUserMethod = "isTaskExistforUser"
	// Service bootstrap methods
	HandleTransactionMethod            = "HandleTransaction"
	BeginNewTransactionMethod          = "BeginNewTransaction"
)

// transaction
const (
	TransactionNotExist = "Transaction does not exist"
)

// SQLSTATE for unique constraint violation
const (
	SQLStateUniqueViolation   = "23505"
	UniqueConstraintViolation = "unique constraint violation: %s"
	UniqueViolation           = "unique_violation"
)
