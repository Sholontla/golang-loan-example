package loanbusiness

import "github.com/google/uuid"

type LoanBusiness struct {
	LoanBusinessID               uuid.UUID
	TransactionID                string
	TransactionBusinessCreatedAt string
}
