package commands

import (
	"github.com/golang/mock/gomock"
	"mtsbank_golang/file_signature/coder"
	"mtsbank_golang/file_signature/mock"
	"mtsbank_golang/file_signature/signature"
	"testing"
)

func TestDecoderCommand(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEncoder := mock.NewMockDecoder(ctrl)

	sig := signature.NewSignatureSha256FromFile()
	sigCompare := signature.NewSignatureSha256FromFile()

	decoder := coder.NewFileDecoder(sig, sigCompare)

	cmd := NewDecodeCommand(decoder)

	_ = mockEncoder
	_ = cmd

}
