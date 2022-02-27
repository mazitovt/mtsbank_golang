package coder

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"mtsbank_golang/lesson6/mock"
	"os"
	"testing"
)

func TestFileEncoder(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSig := mock.NewMockSignature(ctrl)

	source := "source.txt"
	destination := "sign.txt"

	f, err := os.Create(source)
	if err != nil {
		return
	}

	f.Close()
	gomock.InOrder(
		mockSig.EXPECT().HashFile(gomock.Any(), gomock.Any()).Return(nil),
		mockSig.EXPECT().SignatureByte().Return([]byte("hash")),
	)

	encoder := NewFileEncoder(mockSig)
	err = encoder.Encode(source)
	assert.Nil(t, err)

	err = encoder.SaveToFile(destination)
	assert.Nil(t, err)

	des, err := os.Open(destination)
	if err != nil {
		return
	}
	assert.Nil(t, err)

	ar, _ := os.ReadFile(destination)
	assert.Equal(t, []byte("hash"), ar)

	des.Close()

	err = os.Remove(source)
	err = os.Remove(destination)
}
