package signature

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"mtsbank_golang/lesson6/signature/contract"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const separator = "====sign===="

type SignatureSha256 struct {
	date time.Time
	size string
	name string
	hash []byte
}

func NewSignatureSha256FromFile() *SignatureSha256 {
	return &SignatureSha256{}
}

func (s *SignatureSha256) HashFile(file *os.File) (err error) {
	stat, err := file.Stat()
	if err != nil {
		return errors.New("no such file")
	}

	s.size = strconv.FormatInt(stat.Size(), 10)
	s.name = path.Base(file.Name())
	s.date = stat.ModTime()

	var fileData = make([]byte, stat.Size())
	_, err = file.Read(fileData)

	if err != nil {
		return errors.New("error while reading file")
	}
	s.hash = s.makeHashSum(string(fileData))

	return
}

func (s *SignatureSha256) headString() string {
	return strings.Join([]string{s.Date().Format("2006-01-02 15-04-05"), s.Size(), s.Name()}, ":")
}

func (s *SignatureSha256) makeHashSum(text string) []byte {
	sha := sha256.New()
	sha.Write([]byte(text))
	return sha.Sum(nil)
}

func (s *SignatureSha256) SignatureByte() []byte {
	result := bytes.NewBufferString(s.headString())
	result.WriteString(separator)
	result.Write(s.hash)
	return result.Bytes()
}

func (s *SignatureSha256) Equals(s1 *contract.Signature) bool {
	return reflect.DeepEqual(s.hash, (*s1).Hash())
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
