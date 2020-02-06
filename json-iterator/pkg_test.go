package json_iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshal(t *testing.T) {
	_, err := Marshal()
	assert.Nil(t, err)
}

func BenchmarkMarshal(b *testing.B) {
	_, _ = Marshal()
}

func TestRegisterEncoder(t *testing.T) {
	_, err := RegisterEncoder()
	assert.Nil(t, err)
}
