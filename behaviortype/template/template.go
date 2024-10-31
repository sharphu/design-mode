/*
模板模式（Template Method Pattern）定义一个操作的框架，并允许子类实现其中的部分步骤。模板模式的核心是一个基础结构体（或类）提供骨架方法，而具
体实现步骤由子类定义。
模板模式示例，模拟制作饮品的过程，其中 Beverage 是模板接口，包含固定的 PrepareRecipe 方法。不同的饮品（如茶和咖啡）实现 Beverage 接口的具体
步骤方法
*/
package template

import "fmt"

// Beverage 接口，定义了饮品制作的模板方法和步骤
type Beverage interface {
	PrepareRecipe() // 模板方法，定义饮品制作流程
	BoilWater()     // 煮水
	Brew()          // 冲泡
	PourInCup()     // 倒入杯中
	AddCondiments() // 添加调料
}

// BaseBeverage 模板结构体，定义 PrepareRecipe 模板方法
type BaseBeverage struct {
	Beverage
}

// PrepareRecipe 模板方法，定义制作流程
func (b *BaseBeverage) PrepareRecipe() {
	b.BoilWater()
	b.Brew()
	b.PourInCup()
	b.AddCondiments()
}

// Tea 结构体，实现具体的冲泡茶的步骤
type Tea struct {
	BaseBeverage // 嵌入 BaseBeverage 以实现模板方法
}

// NewTea 创建 Tea 实例
func NewTea() *Tea {
	tea := &Tea{}
	tea.BaseBeverage.Beverage = tea // 让 BaseBeverage 知道具体的 Beverage 是 Tea
	return tea
}

// 实现 Tea 的冲泡和添加调料方法
func (t *Tea) BoilWater() {
	fmt.Println("Boiling water for tea...")
}

func (t *Tea) Brew() {
	fmt.Println("Steeping the tea...")
}

func (t *Tea) PourInCup() {
	fmt.Println("Pouring tea into cup...")
}

func (t *Tea) AddCondiments() {
	fmt.Println("Adding lemon to tea...")
}

// Coffee 结构体，实现具体的冲泡咖啡的步骤
type Coffee struct {
	BaseBeverage // 嵌入 BaseBeverage 以实现模板方法
}

// NewCoffee 创建 Coffee 实例
func NewCoffee() *Coffee {
	coffee := &Coffee{}
	coffee.BaseBeverage.Beverage = coffee // 让 BaseBeverage 知道具体的 Beverage 是 Coffee
	return coffee
}

// 实现 Coffee 的冲泡和添加调料方法
func (c *Coffee) BoilWater() {
	fmt.Println("Boiling water for coffee...")
}

func (c *Coffee) Brew() {
	fmt.Println("Dripping coffee through filter...")
}

func (c *Coffee) PourInCup() {
	fmt.Println("Pouring coffee into cup...")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("Adding sugar and milk to coffee...")
}
