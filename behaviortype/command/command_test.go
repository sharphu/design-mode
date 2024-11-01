package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLightCommands(t *testing.T) {
	light := NewLight()
	onCommand := NewOnCommand(light)
	offCommand := NewOffCommand(light)
	remote := &RemoteControl{}

	t.Run("Turn Light On", func(t *testing.T) {
		remote.SetCommand(onCommand)
		remote.PressButton()
		assert.True(t, light.isOn, "Expected the light to be on")
	})

	t.Run("Turn Light Off", func(t *testing.T) {
		remote.SetCommand(offCommand)
		remote.PressButton()
		assert.False(t, light.isOn, "Expected the light to be off")
	})
}
