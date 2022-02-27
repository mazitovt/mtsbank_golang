package contract

import (
	"io/fs"
	"reflect"
	"time"
)

type Signature interface {
	Date() time.Time
	Size() string
	Name() string
	Hash() []byte
	HashFile(string, fs.File) error
	SignatureByte() []byte
	Equal(s Signature) bool //(bool, error)
	ParseString(string) error
}

func NewSignature(signatureOrigin Signature) Signature {
	type1 := reflect.TypeOf(signatureOrigin).Elem()
	obj := reflect.New(type1)
	i := obj.Interface()
	return i.(Signature)
}
