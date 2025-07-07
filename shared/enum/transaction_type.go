package enum

type TransactionType int

const (
	TransactionTypeDeposit TransactionType = iota + 1
	TransactionTypeWithdraw
	TransactionTypePayment
)

func (t TransactionType) String() string {
	switch t {
	case TransactionTypeDeposit:
		return "deposit"
	case TransactionTypeWithdraw:
		return "withdraw"
	case TransactionTypePayment:
		return "payment"
	default:
		return "unknown"
	}
}