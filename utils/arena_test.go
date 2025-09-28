package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	height := 20
	k := "key_1"
	v := ValueStruct{
		Meta:      1,
		Value:     []byte("no step,no miles"),
		ExpiresAt: 1234567890,
		Version:   1,
	}
	area := newArena(1000)
	nodeOffset := area.putNode(height)

	keyOffset := area.putKey([]byte(k))
	keySize := len(k)

	valueOffset := area.putVal(v)
	valueSize := v.EncodedSize()

	nodeTarget := area.getNode(nodeOffset)
	keyTarget := area.getKey(keyOffset, uint16(keySize))
	valueTarget := area.getVal(valueOffset, uint32(valueSize))
	assert.Equal(t, height, len(nodeTarget.tower))
	assert.Equal(t, k, keyTarget)
	assert.Equal(t, v.Value, valueTarget.Value)
}
