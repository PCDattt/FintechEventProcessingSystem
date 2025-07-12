package enum

type TransactionStatus int

const (
	TransactionStatusProcessing TransactionStatus = iota + 1
	TransactionStatusSuccess
	TransactionStatusFailed
)

func (s TransactionStatus) String() string {
	switch s {
	case TransactionStatusProcessing:
		return "processing"
	case TransactionStatusSuccess:
		return "success"
	case TransactionStatusFailed:
		return "failed"
	default:
		return "unknown"
	}
}