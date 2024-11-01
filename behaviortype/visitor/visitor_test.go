package visitor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAreaVisitor(t *testing.T) {
	t.Run("Calculate Circle Area", func(t *testing.T) {
		circle := NewCircle(5)
		areaVisitor := &AreaVisitor{}

		circle.Accept(areaVisitor)
		expectedArea := 3.14 * 5 * 5

		assert.Equal(t, expectedArea, areaVisitor.Area, "Expected Circle area to match")
	})

	t.Run("Calculate Rectangle Area", func(t *testing.T) {
		rectangle := NewRectangle(4, 6)
		areaVisitor := &AreaVisitor{}

		rectangle.Accept(areaVisitor)
		expectedArea := 4.0 * 6.0

		assert.Equal(t, expectedArea, areaVisitor.Area, "Expected Rectangle area to match")
	})
}
