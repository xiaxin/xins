package protocol

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackerPackUnpack(t *testing.T) {
	packer := NewDefaultPacker()

	t.Run("Pack and Unpack", func(t *testing.T) {
		message := NewMessage(1, []byte("HelloWorld"))

		packBytes, err := packer.Pack(message)

		assert.NoError(t, err)

		unpackMessage, err := packer.Unpack(bytes.NewBuffer(packBytes))

		assert.NoError(t, err)

		assert.Equal(t, message.Data(), unpackMessage.Data())

	})

}
