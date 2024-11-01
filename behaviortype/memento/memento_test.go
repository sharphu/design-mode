package memento

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGameMemento(t *testing.T) {
	t.Run("Save and Restore Game State", func(t *testing.T) {
		game := NewGame(1, 100)
		caretaker := NewCaretaker()

		// 保存初始状态
		caretaker.AddMemento(game.Save())
		assert.Equal(t, 1, game.level)
		assert.Equal(t, 100, game.score)

		// 更新状态并保存
		game.level = 2
		game.score = 200
		caretaker.AddMemento(game.Save())
		assert.Equal(t, 2, game.level)
		assert.Equal(t, 200, game.score)

		// 恢复到初始状态
		game.Restore(caretaker.GetMemento(0))
		assert.Equal(t, 1, game.level, "Expected level to be restored to 1")
		assert.Equal(t, 100, game.score, "Expected score to be restored to 100")

		// 恢复到更新后的状态
		game.Restore(caretaker.GetMemento(1))
		assert.Equal(t, 2, game.level, "Expected level to be restored to 2")
		assert.Equal(t, 200, game.score, "Expected score to be restored to 200")
	})
}
