package observer

import "testing"

func TestObserver(t *testing.T) {
	officialAccount := NewWeChatOfficialAccount("golang")

	user1 := NewWechatUser("zhangsan")
	user1.Subscribe(officialAccount)

	user2 := NewWechatUser("lisi")
	user2.Subscribe(officialAccount)

	officialAccount.Publish("关于 golang 设计模式的文章...")
}
