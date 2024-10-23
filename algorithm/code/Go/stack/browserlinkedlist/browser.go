package main

import (
	"fmt"
)

// Node 表示双向链表中的一个节点
type Node struct {
	URL  string // 当前 URL
	Prev *Node  // 指向前一个节点
	Next *Node  // 指向下一个节点
}

// BrowserHistory 表示浏览器历史
type BrowserHistory struct {
	Current *Node // 当前节点
}

// History 接口定义浏览器历史的行为
type History interface {
	Visit(url string)
	Back() string
	Forward() string
	CurrentURL() string
}

// 确保 BrowserHistory 实现了 History 接口
var _ History = (*BrowserHistory)(nil)

// NewBrowserHistory 初始化一个新的 BrowserHistory
func NewBrowserHistory() *BrowserHistory {
	return &BrowserHistory{}
}

// Visit 添加一个新的 URL 到历史中
func (b *BrowserHistory) Visit(url string) {
	newNode := &Node{URL: url}
	if b.Current != nil {
		newNode.Prev = b.Current
		b.Current.Next = newNode
	}
	b.Current = newNode
}

// Back 后退到历史中的前一个 URL
func (b *BrowserHistory) Back() string {
	if b.Current == nil || b.Current.Prev == nil {
		return "没有之前的 URL"
	}
	b.Current = b.Current.Prev
	return b.Current.URL
}

// Forward 前进到历史中的下一个 URL
func (b *BrowserHistory) Forward() string {
	if b.Current == nil || b.Current.Next == nil {
		return "没有下一个 URL"
	}
	b.Current = b.Current.Next
	return b.Current.URL
}

// CurrentURL 返回当前的 URL
func (b *BrowserHistory) CurrentURL() string {
	if b.Current == nil {
		return "没有当前 URL"
	}
	return b.Current.URL
}

func main() {
	history := NewBrowserHistory()

	history.Visit("https://www.google.com")
	history.Visit("https://www.facebook.com")
	history.Visit("https://www.twitter.com")

	fmt.Println("当前 URL:", history.CurrentURL()) // Twitter

	fmt.Println("后退到:", history.Back()) // Facebook
	fmt.Println("后退到:", history.Back()) // Google
	fmt.Println("后退到:", history.Back()) // 没有之前的 URL

	fmt.Println("前进到:", history.Forward()) // Facebook
	fmt.Println("前进到:", history.Forward()) // Twitter
	fmt.Println("前进到:", history.Forward()) // 没有下一个 URL
}
