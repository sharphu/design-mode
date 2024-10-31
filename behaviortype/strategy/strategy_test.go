package strategy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestPaymentContext 测试 PaymentContext 上下文是否使用不同策略
func TestPaymentContext(t *testing.T) {
	// 设置为 CreditCardPayment 并支付
	ccContext := NewPaymentContext(NewCreditCardPayment("Jane Doe", "4321-8765-2109-6543"))
	expected := "Paid 150.00 using Credit Card (Name: Jane Doe, CardNo: 4321-8765-2109-6543)"
	assert.Equal(t, expected, ccContext.Pay(150.0), "PaymentContext.Pay() with CreditCardPayment should return correct string")

	// 设置为 PayPalPayment 并支付
	ppContext := NewPaymentContext(NewPayPalPayment("janedoe@example.com"))
	expected = "Paid 300.00 using PayPal (Email: janedoe@example.com)"
	assert.Equal(t, expected, ppContext.Pay(300.0), "PaymentContext.Pay() with PayPalPayment should return correct string")
}
