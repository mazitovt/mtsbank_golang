package commands

import (
	"github.com/golang/mock/gomock"
	"mtsbank_golang/tests_mocks/coder"
	"mtsbank_golang/tests_mocks/mock"
	"mtsbank_golang/tests_mocks/signature"
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
