/*
策略模式（Strategy Pattern）是一种行为设计模式，允许定义一系列算法，将每种算法封装到独立的策略类中，并使这些策略类可以互相替换。这样可以让客户
端代码在运行时选择具体的算法。
策略模式示例。我们用一个支付系统来演示策略模式。不同的支付方式（如信用卡支付和 PayPal 支付）都是独立的策略，客户端可以动态选择其中一种支付方式。
*/
package strategy

import "fmt"

// PaymentStrategy 定义支付策略接口
type PaymentStrategy interface {
	Pay(amount float64) string
}

// CreditCardPayment 具体策略类，实现信用卡支付
type CreditCardPayment struct {
	Name   string
	CardNo string
}

// NewCreditCardPayment CreditCardPayment 的构造函数
func NewCreditCardPayment(name, cardNo string) *CreditCardPayment {
	return &CreditCardPayment{Name: name, CardNo: cardNo}
}

// Pay 使用信用卡支付实现支付接口
func (c *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Credit Card (Name: %s, CardNo: %s)", amount, c.Name, c.CardNo)
}

// PayPalPayment 具体策略类，实现 PayPal 支付
type PayPalPayment struct {
	Email string
}

// NewPayPalPayment PayPalPayment 的构造函数
func NewPayPalPayment(email string) *PayPalPayment {
	return &PayPalPayment{Email: email}
}

// Pay 使用 PayPal 实现支付接口
func (p *PayPalPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using PayPal (Email: %s)", amount, p.Email)
}

// PaymentContext 上下文类，用于设置和执行支付策略
type PaymentContext struct {
	strategy PaymentStrategy
}

// NewPaymentContext PaymentContext 的构造函数
func NewPaymentContext(strategy PaymentStrategy) *PaymentContext {
	return &PaymentContext{strategy: strategy}
}

// Pay 执行支付，使用当前策略
func (p *PaymentContext) Pay(amount float64) string {
	if p.strategy == nil {
		return "Payment strategy not set!"
	}
	return p.strategy.Pay(amount)
}
