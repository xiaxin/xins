package core

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackerPackAndUnpack(t *testing.T) {
	packer := NewPacker()

	msg := NewMessage(1, []byte("test"))
	packedBytes, err := packer.Pack(msg)
	assert.NoError(t, err)
	assert.NotNil(t, packedBytes)
	assert.Equal(t, packedBytes[8:], []byte("test"))

	r := bytes.NewBuffer(packedBytes)
	newMsg, err := packer.Unpack(r)
	assert.NoError(t, err)
	assert.NotNil(t, newMsg)
	assert.EqualValues(t, reflect.Indirect(reflect.ValueOf(msg.ID())).Interface(), newMsg.ID())
	assert.Equal(t, newMsg.Data(), msg.Data())

}
