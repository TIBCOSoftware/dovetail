package transaction

type TransactionService interface {
	ResolveTransactionInput(txnInputsMetadata []TxnInputAttribute) (map[string]interface{}, error)
	GetInitiatorCertAttribute(attr string) (value string, found bool, err error)
	GetTransactionName() string
	GetTransactionInitiator() (string, error)
	TransactionSecuritySupported() bool
}
