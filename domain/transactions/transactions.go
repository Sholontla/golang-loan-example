package transactions

import "github.com/google/uuid"

type TRansactions struct {
	TransactionUserID    uuid.UUID
	TransactionUser      string
	TransactionCreatedAt string
}

type BankBusinessInfo struct {
	ClabeID uuid.UUID
	Clabe   string
}
