package main

import "fmt"

// 1. 策略接口
type PaymentStrategy interface {
	Pay(amount float64)
}

// 2. 具体策略实现
type CreditCardPayment struct {
	cardNumber string
	name       string
	cvv        string
}

func NewCreditCardPayment(cardNumber, name, cvv string) *CreditCardPayment {
	return &CreditCardPayment{
		cardNumber: cardNumber,
		name:       name,
		cvv:        cvv,
	}
}

func (c *CreditCardPayment) Pay(amount float64) {
	fmt.Printf("使用信用卡支付了%.2f元\n", amount)
	// 实际的信用卡支付逻辑
}

type AlipayPayment struct {
	email string
}

func NewAlipayPayment(email string) *AlipayPayment {
	return &AlipayPayment{
		email: email,
	}
}

func (a *AlipayPayment) Pay(amount float64) {
	fmt.Printf("使用支付宝支付了%.2f元\n", amount)
	// 实际的支付宝支付逻辑
}

type WeChatPayment struct {
	id string
}

func NewWeChatPayment(id string) *WeChatPayment {
	return &WeChatPayment{
		id: id,
	}
}

func (w *WeChatPayment) Pay(amount float64) {
	fmt.Printf("使用微信支付了%.2f元\n", amount)
	// 实际的微信支付逻辑
}

// 3. 上下文
type Item struct {
	name  string
	price float64
}

type ShoppingCart struct {
	paymentStrategy PaymentStrategy
	items           []Item
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{
		items: make([]Item, 0),
	}
}

func (s *ShoppingCart) AddItem(name string, price float64) {
	s.items = append(s.items, Item{name: name, price: price})
}

func (s *ShoppingCart) SetPaymentStrategy(paymentStrategy PaymentStrategy) {
	s.paymentStrategy = paymentStrategy
}

func (s *ShoppingCart) Checkout() error {
	if s.paymentStrategy == nil {
		return fmt.Errorf("请先选择支付方式")
	}

	amount := s.calculateTotal()
	s.paymentStrategy.Pay(amount)
	return nil
}

func (s *ShoppingCart) calculateTotal() float64 {
	var total float64
	for _, item := range s.items {
		total += item.price
	}
	return total
}

func main() {
	cart := NewShoppingCart()
	cart.AddItem("键盘", 250)
	cart.AddItem("鼠标", 150)

	// 用户选择使用信用卡支付
	cart.SetPaymentStrategy(NewCreditCardPayment("1234-5678-9012-3456", "张三", "123"))
	cart.Checkout()

	// 用户改变主意，选择使用支付宝支付
	cart.SetPaymentStrategy(NewAlipayPayment("zhangsan@example.com"))
	cart.Checkout()
}
