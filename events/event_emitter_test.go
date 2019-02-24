package events

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewEventEmitter(t *testing.T) {
	emitter := NewEventEmitter()
	emitter.On("click", func(args ...interface{}) {
		assert.Equal(t, 1, args[0])
	})
	emitter.Emit("click", 1)
	// sleep for goroutine run
	time.Sleep(time.Second)
	emitter.Off("click")
	assert.Equal(t, 0, len(emitter.handlers))
}
