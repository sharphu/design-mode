package state

import (
	"testing"
)

// TestOrderState 测试订单状态转换
func TestOrderState(t *testing.T) {
	// 创建新的订单并处理
	order := NewOrder()

	t.Run("Process Pending State", func(t *testing.T) {
		order.Process() // 处理待处理状态
	})

	t.Run("Process Processed State", func(t *testing.T) {
		order.Process() // 处理已处理状态
	})

	t.Run("Process Shipped State", func(t *testing.T) {
		order.Process() // 处理已发货状态
	})
}
