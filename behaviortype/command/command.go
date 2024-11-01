/*
概念与定义：命令模式（Command Pattern）是一种行为型设计模式，它将请求封装成一个对象，使得用户可以参数化地控制请求的执行、撤销和重做操作。
应用场景：
	1.需要将操作参数化，支持不同操作的调度。
	2.实现撤销、重做功能，如文本编辑器的操作记录。
	3.任务调度系统和操作日志功能。
例子：假设我们实现一个简单的电灯开关控制系统，命令包括“打开”和“关闭”电灯。
*/

package command

import "fmt"

// Command 命令接口，定义执行命令的方法
type Command interface {
	Execute()
}

// Light 是接收者，执行具体操作
type Light struct {
	isOn bool
}

// NewLight 创建新的 Light 实例
func NewLight() *Light {
	return &Light{isOn: false}
}

// On 打开电灯
func (l *Light) On() {
	l.isOn = true
	fmt.Println("The light is on.")
}

// Off 关闭电灯
func (l *Light) Off() {
	l.isOn = false
	fmt.Println("The light is off.")
}

// OnCommand 是打开电灯的具体命令
type OnCommand struct {
	light *Light
}

// NewOnCommand 创建新的 OnCommand 实例
func NewOnCommand(light *Light) *OnCommand {
	return &OnCommand{light: light}
}

// Execute 执行打开电灯的命令
func (c *OnCommand) Execute() {
	c.light.On()
}

// OffCommand 是关闭电灯的具体命令
type OffCommand struct {
	light *Light
}

// NewOffCommand 创建新的 OffCommand 实例
func NewOffCommand(light *Light) *OffCommand {
	return &OffCommand{light: light}
}

// Execute 执行关闭电灯的命令
func (c *OffCommand) Execute() {
	c.light.Off()
}

// RemoteControl 是调用者，用于触发命令
type RemoteControl struct {
	command Command
}

// SetCommand 设置命令
func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

// PressButton 执行当前设置的命令
func (r *RemoteControl) PressButton() {
	r.command.Execute()
}
