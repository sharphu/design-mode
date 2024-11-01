/*
概念与定义：备忘录模式（Memento Pattern）是一种行为型设计模式。它用于在不破坏封装的前提下，捕获对象的内部状态，并在以后需要时将对象恢复到原来的状态。
应用场景：
	1.需要保存对象的状态并在后续还原，例如游戏中的存档。
	2.在事务处理中，支持回滚操作。
	3.避免直接暴露对象的内部状态。
例子：假设我们有一个 Game 对象，记录了游戏的进度（如关卡和分数），并希望在游戏状态保存或还原时使用备忘录模式。
*/

package memento

// Memento 备忘录接口，用于存储状态
type Memento interface {
	GetLevel() int
	GetScore() int
}

// gameMemento 是备忘录的具体实现，保存游戏状态
type gameMemento struct {
	level int
	score int
}

// NewGameMemento 创建新的 gameMemento 实例
func NewGameMemento(level, score int) *gameMemento {
	return &gameMemento{level: level, score: score}
}

func (m *gameMemento) GetLevel() int {
	return m.level
}

func (m *gameMemento) GetScore() int {
	return m.score
}

// Originator 负责创建和恢复备忘录的游戏对象
type Originator interface {
	Save() Memento
	Restore(memento Memento)
}

// Game 是 Originator 的实现，表示游戏状态
type Game struct {
	level int
	score int
}

// NewGame 创建新的 Game 实例
func NewGame(level, score int) *Game {
	return &Game{level: level, score: score}
}

// Save 保存当前状态到备忘录
func (g *Game) Save() Memento {
	return NewGameMemento(g.level, g.score)
}

// Restore 从备忘录中恢复状态
func (g *Game) Restore(memento Memento) {
	g.level = memento.GetLevel()
	g.score = memento.GetScore()
}

// Caretaker 管理备忘录的存储
type Caretaker struct {
	mementos []Memento
}

// NewCaretaker 创建新的 Caretaker 实例
func NewCaretaker() *Caretaker {
	return &Caretaker{}
}

// AddMemento 保存备忘录
func (c *Caretaker) AddMemento(m Memento) {
	c.mementos = append(c.mementos, m)
}

// GetMemento 获取指定索引的备忘录
func (c *Caretaker) GetMemento(index int) Memento {
	if index >= 0 && index < len(c.mementos) {
		return c.mementos[index]
	}
	return nil
}
