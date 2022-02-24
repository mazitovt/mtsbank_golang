package contract

import (
	"os"
	"time"
)

type Signature interface {
	Date() time.Time
	Size() string
	Name() string
	Hash() []byte
	HashFile(file *os.File) error
	SignatureByte() []byte
	Equals(s *Signature) bool //(bool, error)
}
