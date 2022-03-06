package coder

import (
	"io/ioutil"
	"mtsbank_golang/tests_mocks/signature/contract"
	"os"
)

type FileEncoder struct {
	//filePath  string
	signature contract.Signature
}

func NewFileEncoder(signature contract.Signature) *FileEncoder {
	return &FileEncoder{signature: signature}
}

func (fe *FileEncoder) Encode(filePath string) (err error) {

	file, err := os.Open(filePath)
	if err != nil {
		return
	}

	defer file.Close()

	err = fe.signature.HashFile(file.Name(), file)
	if err != nil {
		return
	}

	return nil
}

func (fe *FileEncoder) SaveToFile(path string) (err error) {
	return ioutil.WriteFile(path, fe.signature.SignatureByte(), 0644)
}
