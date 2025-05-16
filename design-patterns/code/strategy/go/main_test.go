package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// TestCreditCardPayment 测试信用卡支付
func TestCreditCardPayment(t *testing.T) {
	// 捕获标准输出
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// 执行支付
	payment := NewCreditCardPayment("1234-5678-9012-3456", "张三", "123")
	payment.Pay(100.00)

	// 恢复标准输出并获取输出内容
	w.Close()
	os.Stdout = oldStdout
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// 验证输出
	expected := "使用信用卡支付了100.00元\n"
	if output != expected {
		t.Errorf("期望输出 %q, 实际得到 %q", expected, output)
	}
}

// TestAlipayPayment 测试支付宝支付
func TestAlipayPayment(t *testing.T) {
	// 捕获标准输出
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// 执行支付
	payment := NewAlipayPayment("zhangsan@example.com")
	payment.Pay(200.00)

	// 恢复标准输出并获取输出内容
	w.Close()
	os.Stdout = oldStdout
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// 验证输出
	expected := "使用支付宝支付了200.00元\n"
	if output != expected {
		t.Errorf("期望输出 %q, 实际得到 %q", expected, output)
	}
}

// TestWeChatPayment 测试微信支付
func TestWeChatPayment(t *testing.T) {
	// 捕获标准输出
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// 执行支付
	payment := NewWeChatPayment("wx123456")
	payment.Pay(300.00)

	// 恢复标准输出并获取输出内容
	w.Close()
	os.Stdout = oldStdout
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// 验证输出
	expected := "使用微信支付了300.00元\n"
	if output != expected {
		t.Errorf("期望输出 %q, 实际得到 %q", expected, output)
	}
}

// TestShoppingCartAddItem 测试添加商品
func TestShoppingCartAddItem(t *testing.T) {
	cart := NewShoppingCart()
	cart.AddItem("键盘", 250)
	cart.AddItem("鼠标", 150)

	if len(cart.items) != 2 {
		t.Errorf("期望商品数量为 2, 实际为 %d", len(cart.items))
	}

	if cart.items[0].name != "键盘" || cart.items[0].price != 250 {
		t.Errorf("商品1信息错误, 期望 {键盘 250}, 实际为 {%s %.2f}", cart.items[0].name, cart.items[0].price)
	}

	if cart.items[1].name != "鼠标" || cart.items[1].price != 150 {
		t.Errorf("商品2信息错误, 期望 {鼠标 150}, 实际为 {%s %.2f}", cart.items[1].name, cart.items[1].price)
	}
}

// TestShoppingCartCalculateTotal 测试计算总价
func TestShoppingCartCalculateTotal(t *testing.T) {
	cart := NewShoppingCart()
	cart.AddItem("键盘", 250)
	cart.AddItem("鼠标", 150)

	total := cart.calculateTotal()
	expected := 400.0

	if total != expected {
		t.Errorf("期望总价为 %.2f, 实际为 %.2f", expected, total)
	}
}

// TestShoppingCartCheckoutNoStrategy 测试未设置支付方式时的结账
func TestShoppingCartCheckoutNoStrategy(t *testing.T) {
	cart := NewShoppingCart()
	cart.AddItem("键盘", 250)

	err := cart.Checkout()
	if err == nil {
		t.Error("未设置支付方式应返回错误，但没有")
	}

	if !strings.Contains(err.Error(), "请先选择支付方式") {
		t.Errorf("错误信息不符合预期，应包含'请先选择支付方式'，实际为：%s", err.Error())
	}
}

// TestShoppingCartCheckout 测试结账功能
func TestShoppingCartCheckout(t *testing.T) {
	// 捕获标准输出
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	cart := NewShoppingCart()
	cart.AddItem("键盘", 250)
	cart.AddItem("鼠标", 150)
	cart.SetPaymentStrategy(NewCreditCardPayment("1234-5678-9012-3456", "张三", "123"))

	err := cart.Checkout()

	// 恢复标准输出并获取输出内容
	w.Close()
	os.Stdout = oldStdout
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if err != nil {
		t.Errorf("结账过程出现错误: %v", err)
	}

	expected := "使用信用卡支付了400.00元\n"
	if output != expected {
		t.Errorf("期望输出 %q, 实际得到 %q", expected, output)
	}
}

// TestShoppingCartChangeStrategy 测试更换支付方式
func TestShoppingCartChangeStrategy(t *testing.T) {
	cart := NewShoppingCart()
	cart.AddItem("键盘", 250)

	// 设置支付宝支付方式
	alipay := NewAlipayPayment("zhangsan@example.com")
	cart.SetPaymentStrategy(alipay)

	// 检查支付方式是否正确设置
	if cart.paymentStrategy != alipay {
		t.Error("支付方式设置失败")
	}

	// 更换为微信支付
	wechat := NewWeChatPayment("wx123456")
	cart.SetPaymentStrategy(wechat)

	// 检查支付方式是否已更换
	if cart.paymentStrategy != wechat {
		t.Error("支付方式更换失败")
	}
}
