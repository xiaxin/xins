package xins

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type Packer interface {
	Pack(*Message) ([]byte, error)
	Unpack(io.Reader) (*Message, error)
}

var _ Packer = &DefaultPacker{}

type DefaultPacker struct{}

func NewDefaultPacker() *DefaultPacker {
	return &DefaultPacker{}
}

func (d *DefaultPacker) bytesOrder() binary.ByteOrder {
	return binary.LittleEndian
}

func (d *DefaultPacker) Pack(msg *Message) ([]byte, error) {
	dataSize := len(msg.Data())

	dataBuff := bytes.NewBuffer([]byte{})

	//写dataLen
	if err := binary.Write(dataBuff, binary.LittleEndian, uint32(dataSize)); err != nil {
		return nil, err
	}

	//写msgID
	if err := binary.Write(dataBuff, binary.LittleEndian, uint32(msg.ID())); err != nil {
		return nil, err
	}

	//写data数据
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.Data()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

func (d *DefaultPacker) Unpack(reader io.Reader) (*Message, error) {
	headBytes := make([]byte, 4+4)
	if _, err := io.ReadFull(reader, headBytes); err != nil {
		if err == io.EOF {
			return nil, err
		}
		return nil, fmt.Errorf("read size and id err: %s", err)
	}

	var (
		dataSize uint32
		dataID   uint32
	)

	headReader := bytes.NewReader(headBytes)

	if err := binary.Read(headReader, binary.LittleEndian, &dataSize); err != nil {
		return nil, err
	}

	if err := binary.Read(headReader, binary.LittleEndian, &dataID); err != nil {
		return nil, err
	}

	data := make([]byte, dataSize)
	if _, err := io.ReadFull(reader, data); err != nil {
		return nil, fmt.Errorf("read data err: %s", err)
	}
	return NewMessage(dataID, data), nil
}
