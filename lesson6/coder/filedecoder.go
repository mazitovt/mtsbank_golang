package coder

import (
	"io/ioutil"
	"mtsbank_golang/lesson6/signature/contract"
	"os"
	"strconv"
)

type FileDecoder struct {
	//sourcePath      string
	//signaturePath   string
	signatureOrigin contract.Signature
	signature       contract.Signature
}

func NewFileDecoder(signatureOrigin contract.Signature) *FileDecoder {
	return &FileDecoder{signatureOrigin: signatureOrigin, signature: contract.NewSignature(signatureOrigin)}
}

func (fe *FileDecoder) Decode(sourcePath, signaturePath string) (err error) {

	buf, err := ioutil.ReadFile(signaturePath)
	if err != nil {
		return
	}

	err = fe.signatureOrigin.ParseString(string(buf))
	if err != nil {
		return
	}

	file, err := os.Open(sourcePath)
	if err != nil {
		return
	}

	defer file.Close()

	err = fe.signature.HashFile(file)
	if err != nil {
		return
	}

	return nil
}

func (fe *FileDecoder) SaveToFile(outFile string) (err error) {
	return ioutil.WriteFile(outFile, []byte(strconv.FormatBool(fe.IsFileChanged())), 0644)
}

func (fe *FileDecoder) IsFileChanged() (isChanged bool) {
	return fe.signatureOrigin.Equal(fe.signature)
}
