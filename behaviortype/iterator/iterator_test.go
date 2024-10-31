package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItemCollection(t *testing.T) {
	collection := NewItemCollection()
	collection.AddItem("Item 1")
	collection.AddItem("Item 2")
	collection.AddItem("Item 3")

	iterator := collection.CreateIterator()

	t.Run("Iterate Over Items", func(t *testing.T) {
		expectedItems := []interface{}{"Item 1", "Item 2", "Item 3"}
		var items []interface{}

		for iterator.Next() {
			items = append(items, iterator.Value())
		}

		assert.Equal(t, expectedItems, items)
	})
}
