package repository

import "gorm.io/gorm"

// RepositoryContext use to create repository context
type RepositoryContext struct {
	_           struct{}
	RequestID   string
	Transaction *gorm.DB
}

// CreateRepositoryContext create repository context
func CreateRepositoryContext(requestID string, transaction *gorm.DB) RepositoryContext {
	return RepositoryContext{
		RequestID:   requestID,
		Transaction: transaction,
	}
}
