package mediator

import "testing"

// MockUser 模拟用户用于测试
type MockUser struct {
	name     string
	mediator Mediator
	messages []string
}

// NewMockUser 创建一个新的 MockUser 实例
func NewMockUser(name string, mediator Mediator) *MockUser {
	return &MockUser{name: name, mediator: mediator, messages: []string{}}
}

// SendMessage 模拟发送消息
func (u *MockUser) SendMessage(message string) {
	u.mediator.SendMessage(message, u)
}

// ReceiveMessage 模拟接收消息
func (u *MockUser) ReceiveMessage(message string) {
	u.messages = append(u.messages, message)
}

func TestChatRoom(t *testing.T) {
	chatRoom := NewChatRoom()

	// 创建用户实例
	user1 := NewMockUser("User1", chatRoom)
	user2 := NewMockUser("User2", chatRoom)
	user3 := NewMockUser("User3", chatRoom)

	// 注册用户到聊天室
	chatRoom.RegisterUser(user1)
	chatRoom.RegisterUser(user2)
	chatRoom.RegisterUser(user3)

	// 测试发送消息
	t.Run("User1 sends message", func(t *testing.T) {
		user1.SendMessage("Hello everyone!")
		if len(user2.messages) != 1 || user2.messages[0] != "Hello everyone!" {
			t.Errorf("Expected User2 to receive 'Hello everyone!', but got %v", user2.messages)
		}
		if len(user3.messages) != 1 || user3.messages[0] != "Hello everyone!" {
			t.Errorf("Expected User3 to receive 'Hello everyone!', but got %v", user3.messages)
		}
	})

	t.Run("User2 sends message", func(t *testing.T) {
		user2.SendMessage("Hi User1!")
		if len(user1.messages) != 1 || user1.messages[0] != "Hi User1!" {
			t.Errorf("Expected User1 to receive 'Hi User1!', but got %v", user1.messages)
		}
		if len(user3.messages) != 2 || user3.messages[1] != "Hi User1!" {
			t.Errorf("Expected User3 to receive 'Hi User1!', but got %v", user3.messages)
		}
	})
}
