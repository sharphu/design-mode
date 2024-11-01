/*
概念与定义：访问者模式（Visitor Pattern）是一种行为型设计模式。它允许在不更改对象结构的前提下，向对象添加新的功能。通过访问者模式，可以将对对象
的操作与对象本身的结构分离，从而实现扩展功能的灵活性。
应用场景：
	1.对结构稳定的对象进行不同种类操作，比如导出不同格式的数据。
	2.希望将数据结构与操作解耦，方便后续添加新功能。
*/

package visitor

import "fmt"

// Visitor 访问者接口，定义了对不同元素的访问操作
type Visitor interface {
	VisitCircle(*Circle)
	VisitRectangle(*Rectangle)
}

// Element 元素接口，定义接受访问者的方法
type Element interface {
	Accept(Visitor)
}

// Circle 圆形元素
type Circle struct {
	Radius float64
}

// NewCircle 创建新的 Circle 实例
func NewCircle(radius float64) *Circle {
	return &Circle{Radius: radius}
}

// Accept 让访问者访问 Circle
func (c *Circle) Accept(v Visitor) {
	v.VisitCircle(c)
}

// Rectangle 矩形元素
type Rectangle struct {
	Width, Height float64
}

// NewRectangle 创建新的 Rectangle 实例
func NewRectangle(width, height float64) *Rectangle {
	return &Rectangle{Width: width, Height: height}
}

// Accept 让访问者访问 Rectangle
func (r *Rectangle) Accept(v Visitor) {
	v.VisitRectangle(r)
}

// AreaVisitor 面积计算访问者
type AreaVisitor struct {
	Area float64
}

// VisitCircle 计算圆形的面积
func (v *AreaVisitor) VisitCircle(c *Circle) {
	v.Area = 3.14 * c.Radius * c.Radius
	fmt.Printf("Circle Area: %f\n", v.Area)
}

// VisitRectangle 计算矩形的面积
func (v *AreaVisitor) VisitRectangle(r *Rectangle) {
	v.Area = r.Width * r.Height
	fmt.Printf("Rectangle Area: %f\n", v.Area)
}
