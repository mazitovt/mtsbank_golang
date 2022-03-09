package signature

import (
	"bytes"
	"encoding/binary"
	"mtsbank_golang/file_signature/packing"
	"time"
)

func (in *SignatureSha256) Unpack(data []byte) error {

	r := bytes.NewReader(data)
	orderByte := binary.BigEndian

	// sizeUint
	var sizeUintRaw uint64
	binary.Read(r, binary.BigEndian, &sizeUintRaw)
	in.sizeUint = uint(sizeUintRaw)

	// date
	data, e := packing.ReadByteSlice(r, orderByte)
	if e != nil {
		return e
	}
	var t time.Time
	e = t.GobDecode(data)
	if e != nil {
		return e
	}

	in.date = t

	// size
	data, e = packing.ReadByteSlice(r, orderByte)
	if e != nil {
		return e
	}
	in.size = string(data)

	// name
	data, e = packing.ReadByteSlice(r, orderByte)
	if e != nil {
		return e
	}
	in.name = string(data)

	// hash
	data, e = packing.ReadByteSlice(r, orderByte)
	if e != nil {
		return e
	}
	in.hash = data
	return nil
}
