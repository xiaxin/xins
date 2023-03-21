package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type packer struct{}

func NewPacker() *packer {
	return &packer{}
}

func (d *packer) bytesOrder() binary.ByteOrder {
	return binary.LittleEndian
}

func (d *packer) Pack(msg *Message) ([]byte, error) {
	dataSize := len(msg.Data())

	dataBuff := bytes.NewBuffer([]byte{})

	//写 Len
	if err := binary.Write(dataBuff, d.bytesOrder(), uint32(dataSize)); err != nil {
		return nil, err
	}

	//写 ID TODO
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.ID()); err != nil {
		return nil, err
	}

	//写 Data
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.Data()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

func (d *packer) Unpack(reader io.Reader) (*Message, error) {
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
