// browserHistory.test.ts

import BrowserHistory from './browserlinkedlist'; // 根据实际文件路径修改

describe('BrowserHistory', () => {
    let bHistory: BrowserHistory;

    beforeEach(() => {
        bHistory = new BrowserHistory();
    });

    test('访问 URL 应该更新当前 URL', () => {
        bHistory.visit("https://www.google.com");
        expect(bHistory.currentUrl()).toBe("https://www.google.com");

        bHistory.visit("https://www.facebook.com");
        expect(bHistory.currentUrl()).toBe("https://www.facebook.com");
    });

    test('后退功能应返回上一个 URL', () => {
        bHistory.visit("https://www.google.com");
        bHistory.visit("https://www.facebook.com");

        expect(bHistory.back()).toBe("https://www.google.com");
        expect(bHistory.currentUrl()).toBe("https://www.google.com");
    });

    test('后退到无前一个 URL 应返回提示信息', () => {
        expect(bHistory.back()).toBe("没有之前的 URL");
    });

    test('前进功能应返回下一个 URL', () => {
        bHistory.visit("https://www.google.com");
        bHistory.visit("https://www.facebook.com");
        bHistory.back(); // 现在在 Google

        expect(bHistory.forward()).toBe("https://www.facebook.com");
        expect(bHistory.currentUrl()).toBe("https://www.facebook.com");
    });

    test('前进到无下一个 URL 应返回提示信息', () => {
        bHistory.visit("https://www.google.com");
        bHistory.visit("https://www.facebook.com");

        expect(bHistory.forward()).toBe("没有下一个 URL");
    });

    test('没有当前 URL 应返回提示信息', () => {
        expect(bHistory.currentUrl()).toBe("没有当前 URL");
    });
});
