package packing

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
	"time"
	"unsafe"
)

func getUnexportedField(field reflect.Value) interface{} {
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()
}

func setUnexportedField(field reflect.Value, value interface{}) {
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).
		Elem().
		Set(reflect.ValueOf(value))
}

func writeByteSlice(buf *bytes.Buffer, orderByte binary.ByteOrder, data []byte) (err error) {

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

func readByteSlice(r io.Reader, orderByte binary.ByteOrder) (data []byte, err error) {

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

func PackSignature(u interface{}) (buf bytes.Buffer, err error) {
	orderByte := binary.BigEndian
	value := reflect.ValueOf(u).Elem()

	for i := 0; i < value.NumField(); i++ {

		fieldValue := getUnexportedField(reflect.ValueOf(u).Elem().Field(i))

		switch v := fieldValue.(type) {
		case string:
			err = writeByteSlice(&buf, orderByte, []byte(v))
			if err != nil {
				return
			}
		case uint:
			err = binary.Write(&buf, orderByte, uint64(v))
			if err != nil {
				return
			}
		case []byte:
			err = writeByteSlice(&buf, orderByte, v)
			if err != nil {
				return
			}
		case time.Time:
			encode, e := v.GobEncode()
			if e != nil {
				return buf, e
			}
			err = writeByteSlice(&buf, orderByte, encode)
			if err != nil {
				return
			}
		default:
			fmt.Println("no such type")
		}
	}
	return
}

func UnpackSignature(u interface{}, buf []byte) (err error) {

	r := bytes.NewReader(buf)

	orderByte := binary.BigEndian
	value := reflect.ValueOf(u).Elem()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		switch getUnexportedField(field).(type) {
		case string:
			data, e := readByteSlice(r, orderByte)
			if e != nil {
				return e
			}
			setUnexportedField(field, string(data))
		case uint:
			var data uint64
			err = binary.Read(r, orderByte, &data)
			if err != nil {
				return
			}
			setUnexportedField(field, uint(data))
		case []byte:
			data, e := readByteSlice(r, orderByte)
			if e != nil {
				return e
			}
			setUnexportedField(field, data)
		case time.Time:
			data, e := readByteSlice(r, orderByte)
			if e != nil {
				return e
			}
			var t time.Time
			err = t.GobDecode(data)
			if err != nil {
				return
			}
			setUnexportedField(field, t)
		default:
			fmt.Println("no such type")
		}
	}

	return
}
