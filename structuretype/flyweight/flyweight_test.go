package flyweight

import (
	"testing"
)

// TestCircleFactory_Singleton 测试 CircleFactory 是否正确复用相同颜色的 Circle 实例
func TestCircleFactory_Singleton(t *testing.T) {
	factory := NewCircleFactory()

	// 获取两次相同颜色的 Circle 实例
	redCircle1 := factory.GetCircle("red")
	redCircle2 := factory.GetCircle("red")

	// 确保两次获取到的是同一个实例
	if redCircle1 != redCircle2 {
		t.Error("Expected redCircle1 and redCircle2 to be the same instance")
	}

	// 验证 Circle 实例的数量是否正确
	if len(factory.circleMap) != 1 {
		t.Errorf("Expected 1 unique Circle instance, got %d", len(factory.circleMap))
	}
}

// TestCircleFactory_UniqueColors 测试 CircleFactory 是否正确创建不同颜色的 Circle 实例
func TestCircleFactory_UniqueColors(t *testing.T) {
	factory := NewCircleFactory()

	// 获取不同颜色的 Circle 实例
	redCircle := factory.GetCircle("red")
	blueCircle := factory.GetCircle("blue")
	greenCircle := factory.GetCircle("green")

	// 确保不同颜色的实例是不同的
	if redCircle == blueCircle || redCircle == greenCircle || blueCircle == greenCircle {
		t.Error("Expected all Circle instances to be unique for different colors")
	}

	// 验证 Circle 实例的数量是否正确
	if len(factory.circleMap) != 3 {
		t.Errorf("Expected 3 unique Circle instances, got %d", len(factory.circleMap))
	}
}

// TestCircle_Draw 测试 Circle 的 Draw 方法
func TestCircle_Draw(t *testing.T) {
	factory := NewCircleFactory()
	circle := factory.GetCircle("yellow")

	// 在 Draw 方法中可以简单地检查是否运行成功，或通过mock输出
	circle.Draw(10, 10)
	// 没有具体的测试断言，只需确保 Draw 不会发生错误即可
}
