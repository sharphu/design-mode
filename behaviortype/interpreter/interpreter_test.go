package interpreter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterpreter(t *testing.T) {
	t.Run("Addition", func(t *testing.T) {
		parser := NewParser("2 3 +")
		result := parser.Interpret()
		assert.Equal(t, 5, result, "Expected 2 + 3 to equal 5")
	})

	t.Run("Subtraction", func(t *testing.T) {
		parser := NewParser("10 4 -")
		result := parser.Interpret()
		assert.Equal(t, 6, result, "Expected 10 - 4 to equal 6")
	})

	t.Run("Addition and Subtraction", func(t *testing.T) {
		parser := NewParser("5 3 + 2 -")
		result := parser.Interpret()
		assert.Equal(t, 6, result, "Expected 5 + 3 - 2 to equal 6")
	})

	t.Run("Complex Expression", func(t *testing.T) {
		parser := NewParser("7 3 + 2 - 4 +")
		result := parser.Interpret()
		assert.Equal(t, 12, result, "Expected 7 + 3 - 2 + 4 to equal 12")
	})
}
