// 1. 策略接口
export interface PaymentStrategy {
    pay(amount: number): void;
}

// 2. 具体策略实现
export class CreditCardPayment implements PaymentStrategy {
    private cardNumber: string;
    private name: string;
    private cvv: string;

    constructor(cardNumber: string, name: string, cvv: string) {
        this.cardNumber = cardNumber;
        this.name = name;
        this.cvv = cvv;
    }

    pay(amount: number) {
        console.log(`使用信用卡支付了${amount}元`);
        // 实际的信用卡支付逻辑
    }
}

export class AlipayPayment implements PaymentStrategy {
    private email: string;

    constructor(email: string) {
        this.email = email;
    }

    pay(amount: number) {
        console.log(`使用支付宝支付了${amount}元`);
        // 实际的支付宝支付逻辑
    }
}

export class WeChatPayment implements PaymentStrategy {
    private id: string;

    constructor(id: string) {
        this.id = id;
    }

    pay(amount: number) {
        console.log(`使用微信支付了${amount}元`);
        // 实际的微信支付逻辑
    }
}

// 3. 上下文
export class ShoppingCart {
    private paymentStrategy: PaymentStrategy | null = null;
    private items: { name: string, price: number }[] = [];

    addItem(name: string, price: number) {
        this.items.push({ name, price });
    }

    setPaymentStrategy(paymentStrategy: PaymentStrategy) {
        this.paymentStrategy = paymentStrategy;
    }

    checkout() {
        if (!this.paymentStrategy) {
            throw new Error('请先选择支付方式');
        }

        const amount = this.calculateTotal();
        this.paymentStrategy.pay(amount);
    }

    private calculateTotal(): number {
        return this.items.reduce((total, item) => total + item.price, 0);
    }
}

// 使用示例
if (require.main === module) {
    const cart = new ShoppingCart();
    cart.addItem('键盘', 1150);
    cart.addItem('鼠标', 1050);

    // 用户选择使用信用卡支付
    cart.setPaymentStrategy(new CreditCardPayment('1234-5678-9012-3456', '张三', '123'));
    cart.checkout();

    // 用户改变主意，选择使用支付宝支付
    cart.setPaymentStrategy(new AlipayPayment('zhangsan@example.com'));
    cart.checkout();
}