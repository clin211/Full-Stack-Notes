import {
    CreditCardPayment,
    AlipayPayment,
    WeChatPayment,
    ShoppingCart
} from './main';

// Mock console.log to track outputs
let consoleOutput: string[] = [];
const originalLog = console.log;
beforeEach(() => {
    consoleOutput = [];
    console.log = jest.fn((...args) => {
        consoleOutput.push(args.join(' '));
    });
});
afterEach(() => {
    console.log = originalLog;
});

describe('Payment Strategies', () => {
    describe('CreditCardPayment', () => {
        test('should log correct payment message', () => {
            const payment = new CreditCardPayment('1234-5678-9012-3456', '张三', '123');
            payment.pay(100);
            expect(consoleOutput[0]).toBe('使用信用卡支付了100元');
        });
    });

    describe('AlipayPayment', () => {
        test('should log correct payment message', () => {
            const payment = new AlipayPayment('zhangsan@example.com');
            payment.pay(200);
            expect(consoleOutput[0]).toBe('使用支付宝支付了200元');
        });
    });

    describe('WeChatPayment', () => {
        test('should log correct payment message', () => {
            const payment = new WeChatPayment('wx123456');
            payment.pay(300);
            expect(consoleOutput[0]).toBe('使用微信支付了300元');
        });
    });
});

describe('ShoppingCart', () => {
    let cart: ShoppingCart;

    beforeEach(() => {
        cart = new ShoppingCart();
    });

    test('should add items correctly', () => {
        cart.addItem('键盘', 250);
        cart.addItem('鼠标', 150);

        // Accessing private property for testing
        const items = (cart as any).items;

        expect(items.length).toBe(2);
        expect(items[0]).toEqual({ name: '键盘', price: 250 });
        expect(items[1]).toEqual({ name: '鼠标', price: 150 });
    });

    test('should calculate total correctly', () => {
        cart.addItem('键盘', 250);
        cart.addItem('鼠标', 150);

        // Accessing private method for testing
        const total = (cart as any).calculateTotal();

        expect(total).toBe(400);
    });

    test('should throw error when checking out without strategy', () => {
        cart.addItem('键盘', 250);

        expect(() => {
            cart.checkout();
        }).toThrow('请先选择支付方式');
    });

    test('should checkout with credit card payment', () => {
        cart.addItem('键盘', 250);
        cart.addItem('鼠标', 150);
        cart.setPaymentStrategy(new CreditCardPayment('1234-5678-9012-3456', '张三', '123'));

        cart.checkout();

        expect(consoleOutput[0]).toBe('使用信用卡支付了400元');
    });

    test('should change payment strategy', () => {
        cart.addItem('键盘', 250);

        // First use credit card
        const creditCard = new CreditCardPayment('1234-5678-9012-3456', '张三', '123');
        cart.setPaymentStrategy(creditCard);
        cart.checkout();
        expect(consoleOutput[0]).toBe('使用信用卡支付了250元');

        // Then switch to Alipay
        consoleOutput = []; // Clear previous output
        const alipay = new AlipayPayment('zhangsan@example.com');
        cart.setPaymentStrategy(alipay);
        cart.checkout();
        expect(consoleOutput[0]).toBe('使用支付宝支付了250元');

        // Finally switch to WeChat
        consoleOutput = []; // Clear previous output
        const wechat = new WeChatPayment('wx123456');
        cart.setPaymentStrategy(wechat);
        cart.checkout();
        expect(consoleOutput[0]).toBe('使用微信支付了250元');
    });
});