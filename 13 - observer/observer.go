package observer

import "fmt"

// Observable 被观察者接口
type Observable interface {
	Notify()
}

// Observer 观察者接口
type Observer interface {
	Update(*WeChatOfficialAccount)
}

var _ Observable = (*WeChatOfficialAccount)(nil)

// WeChatOfficialAccount 微信公众号
type WeChatOfficialAccount struct {
	Name       string
	NewArticle string
	subscriber []Observer
}

// NewWeChatOfficialAccount .
func NewWeChatOfficialAccount(name string) *WeChatOfficialAccount {
	return &WeChatOfficialAccount{
		Name:       name,
		subscriber: make([]Observer, 0),
	}
}

// AddFollower .
func (w *WeChatOfficialAccount) AddFollower(o Observer) {
	w.subscriber = append(w.subscriber, o)
}

// Publish 发布
func (w *WeChatOfficialAccount) Publish(newArticle string) {
	w.NewArticle = newArticle
	w.Notify()
}

// Notify 通知观察者们
func (w *WeChatOfficialAccount) Notify() {
	for _, s := range w.subscriber {
		s.Update(w)
	}
}

var _ Observer = (*WechatUser)(nil)

// WechatUser wechat用户
type WechatUser struct {
	Name string
}

func NewWechatUser(name string) *WechatUser {
	return &WechatUser{Name: name}
}

// Subscribe 订阅
func (u *WechatUser) Subscribe(woa *WeChatOfficialAccount) {
	woa.AddFollower(u)
}

// Update 接收通知
func (u *WechatUser) Update(w *WeChatOfficialAccount) {
	fmt.Printf("--------user: %s--------\n\t微信公众号：%s 新文章：%s\n\n", u.Name, w.Name, w.NewArticle)
}
