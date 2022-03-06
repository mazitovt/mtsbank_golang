package coder

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"mtsbank_golang/file_signature/mock"
	"os"
	"testing"
)

func TestFileDecoder(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSigOrigin := mock.NewMockSignature(ctrl)
	mockSigCompare := mock.NewMockSignature(ctrl)

	sourceName := "source.txt"
	signatureName := "sign.txt"
	destinationName := "out.txt"

	defer func() {
		os.Remove(sourceName)
		os.Remove(signatureName)
		os.Remove(destinationName)
	}()

	f, err := os.Create(sourceName)
	if err != nil {
		return
	}

	f.Close()

	f1, err := os.Create(signatureName)
	if err != nil {
		return
	}

	f1.Close()
	gomock.InOrder(
		mockSigOrigin.EXPECT().ParseString("").Return(nil),
		mockSigCompare.EXPECT().HashFile(gomock.Any(), gomock.Any()).Return(nil),
		mockSigOrigin.EXPECT().Equal(gomock.Any()).Return(true),
	)

	encoder := NewFileDecoder(mockSigOrigin, mockSigCompare)
	err = encoder.Decode(sourceName, signatureName)
	assert.Nil(t, err)

	err = encoder.SaveToFile(destinationName)
	assert.Nil(t, err)

	des, err := os.Open(destinationName)
	if err != nil {
		return
	}
	assert.Nil(t, err)

	ar, _ := os.ReadFile(destinationName)
	assert.Equal(t, []byte("true"), ar)

	des.Close()

}
