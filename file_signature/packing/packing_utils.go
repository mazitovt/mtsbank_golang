package packing

import (
	"bytes"
	"encoding/binary"
	"io"
)

func WriteByteSlice(buf *bytes.Buffer, orderByte binary.ByteOrder, data []byte) (err error) {

	err = binary.Write(buf, orderByte, uint32(len(data)))
	if err != nil {
		return
	}

	err = binary.Write(buf, orderByte, data)
	if err != nil {
		return
	}

	return
}

func ReadByteSlice(r io.Reader, orderByte binary.ByteOrder) (data []byte, err error) {

	var length uint32

	err = binary.Read(r, orderByte, &length)
	if err != nil {
		return
	}

	data = make([]byte, length)
	err = binary.Read(r, orderByte, &data)
	if err != nil {
		return
	}

	return
}
