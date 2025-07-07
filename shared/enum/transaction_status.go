package enum

type TransactionStatus int

const (
	TransactionStatusPending TransactionStatus = iota + 1
	TransactionStatusSuccess
	TransactionStatusFailed
)

func (s TransactionStatus) String() string {
	switch s {
	case TransactionStatusPending:
		return "pending"
	case TransactionStatusSuccess:
		return "success"
	case TransactionStatusFailed:
		return "failed"
	default:
		return "unknown"
	}
}