package mycrypt

import (
	"io/ioutil"
	"mtsbank_golang/lesson6/signature/contract"
	"os"
)

type FileEncoder struct {
	filePath  string
	signature contract.Signature
}

func NewEncoder(filePath string, signature contract.Signature) *FileEncoder {
	return &FileEncoder{filePath: filePath, signature: signature}
}

func (enc *FileEncoder) EncryptSha256() (err error) {

	file, err := os.Open(enc.filePath)
	if err != nil {
		return
	}

	defer file.Close()

	err = enc.signature.HashFile(file)
	if err != nil {
		return
	}

	return nil
}

func (enc *FileEncoder) SaveToFile(path string) (err error) {
	return ioutil.WriteFile(path, enc.signature.SignatureByte(), 0644)
}
