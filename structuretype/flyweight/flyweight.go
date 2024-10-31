/*
享元模式（Flyweight Pattern）用于减少内存消耗，通过共享相同的数据来减少内存中的对象数量。它通常会创建一个工厂或管理器来存储和复用共享的对象
例子中，我们模拟一个图形绘制系统，绘制各种形状和颜色的图形。为了优化内存消耗，相同颜色的图形将共享一个对象实例，而不是为每个图形创建新的实例
*/
package flyweight

import "fmt"

// Shape 接口，定义图形的绘制行为
type Shape interface {
	Draw(x, y int)
}

// Circle 结构体，表示具体享元类
type Circle struct {
	color string
}

// Draw 在指定位置绘制带颜色的圆
func (c *Circle) Draw(x, y int) {
	fmt.Printf("Drawing a %s circle at (%d, %d)\n", c.color, x, y)
}

// CircleFactory 工厂，用于创建和管理 Circle 实例
type CircleFactory struct {
	circleMap map[string]*Circle
}

// GetCircle 获取具有指定颜色的 Circle，如果不存在则创建新的
func (f *CircleFactory) GetCircle(color string) *Circle {
	if circle, exists := f.circleMap[color]; exists {
		return circle
	}
	// 如果不存在，则创建新的 Circle 实例并存储
	circle := &Circle{color: color}
	f.circleMap[color] = circle
	fmt.Printf("Creating a new %s circle\n", color)
	return circle
}

// NewCircleFactory 创建一个新的 CircleFactory
func NewCircleFactory() *CircleFactory {
	return &CircleFactory{circleMap: make(map[string]*Circle)}
}
