package types

type Money int64

type Currency string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)


type Card struct {
	ID        int
	PAN       string
	Balance   Money
	Currency  Currency
	Color     string
	Name      string
	Active    bool
	MinBalance Money
}

type PaymentSource struct {
	Type     string
	Number   string
	Balance  Money
}

type Category string

const (
	FOOD Category = "FOOD"
	CLOTHES Category = "CLOTHES"
	STUDY Category = "STUDY"
)

type Status string

const (
	StatusOk Status = "Ok"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

type Payment struct {
	ID        int64
	Amount    Money
	Category  Category
	Status    Status
}