package creational

import (
	"errors"
	"fmt"
)

// PaymentMethod ...
type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	//Cash ...
	Cash = 1
	//DebitCard ...
	DebitCard = 2
)

//GetPaymentMethod ...
func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(CreditCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized\n", m))
	}
}

//CashPM ...
type CashPM struct{}

//DebitCardPM ...
type DebitCardPM struct{}

//CreditCardPM ...
type CreditCardPM struct{}

//Pay ...
func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

//Pay ...
func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using DEBIT card\n", amount)
}

//Pay ...
func (c *CreditCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using CREDIT card\n", amount)
}
