package signature

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"io/fs"
	"mtsbank_golang/tests_mocks/signature/contract"
	"path"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const separator1 = "====sign===="
const separator2 = "::"

type SignatureSha256 struct {
	sizeUint uint
	date     time.Time
	size     string
	name     string
	hash     []byte
}

func NewSignatureSha256FromFile() *SignatureSha256 {
	return &SignatureSha256{}
}

func (s *SignatureSha256) hashSum(text string) []byte {
	sha := sha256.New()
	sha.Write([]byte(text))
	return sha.Sum(nil)
}

func (s *SignatureSha256) headString() string {
	return strings.Join([]string{s.Date().Format(time.RFC3339), s.Size(), s.Name()}, separator2)
}

func (s *SignatureSha256) HashFile(name string, file fs.File) (err error) {

	stat, err := file.Stat()
	if err != nil {
		return errors.New("no such file")
	}

	s.size = strconv.FormatInt(stat.Size(), 10)
	s.name = path.Base(name)
	s.date = stat.ModTime().Round(time.Second)

	var fileData = make([]byte, stat.Size())
	_, err = file.Read(fileData)

	if err != nil {
		return errors.New("error while reading file")
	}
	s.hash = s.hashSum(string(fileData))

	return
}

func (s *SignatureSha256) ParseString(signatureString string) (err error) {

	signElems := strings.Split(signatureString, separator1)

	if len(signElems) != 2 {
		err = errors.New("invalid argument signatureString")
		return
	}

	s.hash = []byte(signElems[1])

	data := strings.Split(signElems[0], separator2)
	if len(data) != 3 {
		err = errors.New("invalid data before sign separator1")
		return
	}

	s.date, err = time.Parse(time.RFC3339, data[0])
	if err != nil {
		return
	}

	s.size = data[1]
	s.name = data[2]

	return
}

func (s *SignatureSha256) SignatureByte() []byte {
	result := bytes.NewBufferString(s.headString())
	result.WriteString(separator1)
	result.Write(s.hash)
	return result.Bytes()
}

func (s *SignatureSha256) Equal(s1 contract.Signature) bool {
	return reflect.DeepEqual(s, s1)
}

func (s *SignatureSha256) Date() time.Time {
	return s.date
}

func (s *SignatureSha256) Size() string {
	return s.size
}

func (s *SignatureSha256) Name() string {
	return s.name
}

func (s *SignatureSha256) Hash() []byte {
	return s.hash
}

func (s *SignatureSha256) SetSizeUint(size uint) {
	s.sizeUint = size
}
