/*
状态模式允许一个对象在其内部状态变化时改变其行为.
例子，展示一个 Order（订单）类的状态转换。
*/

package state

import "fmt"

// State 状态接口
type State interface {
	HandleOrder(*Order)
}

// Order 订单结构体
type Order struct {
	state State
}

// NewOrder 创建新的订单
func NewOrder() *Order {
	order := &Order{}
	order.state = &PendingState{} // 初始状态为 Pending
	return order
}

// SetState 设置订单状态
func (o *Order) SetState(state State) {
	o.state = state
}

// Process 处理订单
func (o *Order) Process() {
	o.state.HandleOrder(o)
}

// PendingState 待处理状态
type PendingState struct{}

// HandleOrder 处理待处理状态
func (p *PendingState) HandleOrder(order *Order) {
	fmt.Println("Order is in pending state.")
	order.SetState(&ProcessedState{}) // 转换到 Processed 状态
}

// ProcessedState 已处理状态
type ProcessedState struct{}

// HandleOrder 处理已处理状态
func (p *ProcessedState) HandleOrder(order *Order) {
	fmt.Println("Order has been processed.")
	order.SetState(&ShippedState{}) // 转换到 Shipped 状态
}

// ShippedState 已发货状态
type ShippedState struct{}

// HandleOrder 处理已发货状态
func (s *ShippedState) HandleOrder(order *Order) {
	fmt.Println("Order has been shipped.")
	// 可以继续转换到其他状态或保持当前状态
}
