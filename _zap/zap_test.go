package _zap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	_, err := NewLogger(
		WithLoggingLevel(0),
		WithEncoded("console"),
		WithColor(),
		WithAddCaller(),
		WithAddStacktrace(-1),
	)
	assert.NoError(t, err)
}
