/*
概念与定义：迭代器模式允许通过统一的方式遍历集合对象的元素，而无需暴露集合的内部表示。迭代器可以逐个访问集合中的每个元素。这样做的好处是将遍历逻辑与集合的内部
结构分离，增强了代码的灵活性和可维护性。
*/

package iterator

// Iterator 迭代器接口
type Iterator interface {
	Next() bool
	Value() interface{}
}

// Item 代表集合中的一个元素
type Item struct {
	value interface{}
}

// ItemCollection 具体集合实现
type ItemCollection struct {
	items []Item
}

// NewItemCollection 创建新的 ItemCollection
func NewItemCollection() *ItemCollection {
	return &ItemCollection{
		items: []Item{},
	}
}

// AddItem 向集合中添加元素
func (c *ItemCollection) AddItem(value interface{}) {
	c.items = append(c.items, Item{value: value})
}

// CreateIterator 创建迭代器
func (c *ItemCollection) CreateIterator() Iterator {
	return &ItemIterator{
		collection: c,
		index:      0,
	}
}

// ItemIterator 具体迭代器实现
type ItemIterator struct {
	collection *ItemCollection
	index      int
}

// Next 移动到下一个元素
func (i *ItemIterator) Next() bool {
	if i.index < len(i.collection.items) {
		i.index++
		return i.index <= len(i.collection.items)
	}
	return false
}

// Value 获取当前元素的值
func (i *ItemIterator) Value() interface{} {
	if i.index == 0 || i.index > len(i.collection.items) {
		return nil
	}
	return i.collection.items[i.index-1].value
}
