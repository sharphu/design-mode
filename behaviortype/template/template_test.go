package template

import (
	"testing"
)

// TestTea_PrepareRecipe 测试 Tea 的 PrepareRecipe 方法
func TestTea_PrepareRecipe(t *testing.T) {
	tea := NewTea()

	// 验证 PrepareRecipe 是否按顺序调用各步骤
	tea.PrepareRecipe() // 输出测试，看是否按预期输出
}

// TestCoffee_PrepareRecipe 测试 Coffee 的 PrepareRecipe 方法
func TestCoffee_PrepareRecipe(t *testing.T) {
	coffee := NewCoffee()

	// 验证 PrepareRecipe 是否按顺序调用各步骤
	coffee.PrepareRecipe() // 输出测试，看是否按预期输出
}
