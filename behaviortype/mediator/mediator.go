/*
概念与定义：中介者模式（Mediator Pattern）是一种行为设计模式，它定义了一个中介对象，来封装一系列对象之间的交互，使对象之间不需要直接引用彼此，
从而松散耦合。这种模式有助于简化对象之间的通信，将交互逻辑集中到一个中介者对象中。
应用场景：
	1.多个对象之间复杂通信：当多个对象之间有复杂的交互关系时，中介者可以简化其通信。
	2.解耦：减少对象之间的直接依赖关系，便于独立开发和维护。
	3.GUI组件交互：在用户界面开发中，不同组件之间的相互作用可以通过中介者来管理，避免直接依赖。
例子：假设我们有一个聊天室系统，用户通过聊天室进行消息交换。聊天室作为中介者，负责管理用户之间的消息传递。
*/

package mediator

import "fmt"

// User 是一个用户接口，定义了发送和接收消息的方法
type User interface {
	SendMessage(message string)
	ReceiveMessage(message string)
}

// Mediator 是中介者接口，用于管理用户间的消息传递
type Mediator interface {
	RegisterUser(user User)
	SendMessage(message string, sender User)
}

// ChatRoom 是一个具体的中介者，实现了 Mediator 接口
type ChatRoom struct {
	users []User
}

// NewChatRoom 创建一个新的 ChatRoom 实例
func NewChatRoom() *ChatRoom {
	return &ChatRoom{users: []User{}}
}

// RegisterUser 注册用户到聊天室
func (c *ChatRoom) RegisterUser(user User) {
	c.users = append(c.users, user)
}

// SendMessage 发送消息到所有用户，除了发送者自己
func (c *ChatRoom) SendMessage(message string, sender User) {
	for _, user := range c.users {
		if user != sender {
			user.ReceiveMessage(message)
		}
	}
}

// ConcreteUser 是一个具体的用户，实现了 User 接口
type ConcreteUser struct {
	name     string
	mediator Mediator
}

// NewConcreteUser 创建一个新的 ConcreteUser 实例
func NewConcreteUser(name string, mediator Mediator) *ConcreteUser {
	return &ConcreteUser{name: name, mediator: mediator}
}

// SendMessage 通过中介者发送消息
func (u *ConcreteUser) SendMessage(message string) {
	fmt.Printf("[%s] 发送消息: %s\n", u.name, message)
	u.mediator.SendMessage(message, u)
}

// ReceiveMessage 接收消息
func (u *ConcreteUser) ReceiveMessage(message string) {
	fmt.Printf("[%s] 接收到消息: %s\n", u.name, message)
}
