package httpadapters

import "time"

type TransactionBody struct {
	ID            int       `json:"id"`
	AccountID     int       `json:"account_id"`
	OperationType int       `json:"operation_type_id"`
	Amount        float64   `json:"amount"`
	EventDate     time.Time `json:"event_date"`
}

func (tb *TransactionBody) GetID() int {
	return tb.ID
}

func (tb *TransactionBody) GetAccountID() int {
	return tb.AccountID
}

func (tb *TransactionBody) GetOperationType() int {
	return tb.OperationType
}

func (tb *TransactionBody) GetAmount() float64 {
	return tb.Amount
}

func (tb *TransactionBody) GetEventDate() time.Time {
	return tb.EventDate
}

func (tb *TransactionBody) SetID(id int) {
	tb.ID = id
}

func (tb *TransactionBody) SetAccountID(accountId int) {
	tb.AccountID = accountId
}

func (tb *TransactionBody) SetOperationType(operationType int) {
	tb.OperationType = operationType
}

func (tb *TransactionBody) SetAmount(amount float64) {
	tb.Amount = amount
}

func (tb *TransactionBody) SetEventDate(date time.Time) {
	tb.EventDate = date
}
