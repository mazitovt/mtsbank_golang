package packing

import (
	"github.com/stretchr/testify/assert"
	"mtsbank_golang/lesson6/signature"
	"testing"
)

func TestPackSignature(t *testing.T) {

	signTxt := "2022-02-27T21:02:09+05:00::3::source.txt====sign====abc"
	sig := signature.NewSignatureSha256FromFile()
	sig.ParseString(signTxt)

	er := []byte{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 15,
		1, 0, 0, 0, 14, 217, 173, 151, 129, 0, 0, 0, 0, 1, 44,
		0, 0, 0, 1,
		51,
		0, 0, 0, 10,
		115, 111, 117, 114, 99, 101, 46, 116, 120, 116,
		0, 0, 0, 3,
		97, 98, 99,
	}

	b, e := PackSignature(sig)
	if e != nil {
		panic(e)
	}

	assert.Equal(t, b.Bytes(), er)

}

func TestPackEmptyStruct(t *testing.T) {

	b, e := PackSignature(&struct{}{})
	if e != nil {
		t.Error(e)
	}
	assert.Nil(t, b.Bytes())
}

func TestPackStruct(t *testing.T) {

	testCases := []struct {
		i uint
		s string
	}{
		{1, "abc"},
		{999, "abc"},
		{999, ""},
	}
	er := [][]byte{
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 3, 97, 98, 99},
		{0, 0, 0, 0, 0, 0, 3, 231, 0, 0, 0, 3, 97, 98, 99},
		{0, 0, 0, 0, 0, 0, 3, 231, 0, 0, 0, 0},
	}

	for i, testCase := range testCases {
		b, e := PackSignature(&testCase)
		if e != nil {
			t.Error(e)
		}
		assert.Equal(t, b.Bytes(), er[i])
	}
}

func TestUnpackStruct(t *testing.T) {
	er := []struct {
		i uint
		s string
	}{
		{1, "abc"},
		{999, "abc"},
		{999, ""},
	}
	testCases := [][]byte{
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 3, 97, 98, 99},
		{0, 0, 0, 0, 0, 0, 3, 231, 0, 0, 0, 3, 97, 98, 99},
		{0, 0, 0, 0, 0, 0, 3, 231, 0, 0, 0, 0},
	}

	var u struct {
		i uint
		s string
	}
	for i, testCase := range testCases {
		e := UnpackSignature(&u, testCase)
		if e != nil {
			t.Error(e)
		}
		assert.Equal(t, u, er[i])
	}
}

func TestUnpackEmptyStruct(t *testing.T) {

	var u struct{}
	e := UnpackSignature(&u, []byte{})
	if e != nil {
		t.Error(e)
	}
	assert.Equal(t, u, struct{}{})
}

func TestUnpackSignature(t *testing.T) {

	signTxt := "2022-02-27T21:02:09+05:00::3::source.txt====sign====abc"
	er := signature.NewSignatureSha256FromFile()
	er.ParseString(signTxt)

	testCase := []byte{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 15,
		1, 0, 0, 0, 14, 217, 173, 151, 129, 0, 0, 0, 0, 1, 44,
		0, 0, 0, 1,
		51,
		0, 0, 0, 10,
		115, 111, 117, 114, 99, 101, 46, 116, 120, 116,
		0, 0, 0, 3,
		97, 98, 99,
	}

	var u signature.SignatureSha256
	e := UnpackSignature(&u, testCase)
	if e != nil {
		panic(e)
	}

	assert.Equal(t, u, *er)

}

func TestPackAndUnpackSignature(t *testing.T) {
	signTxt := "2022-02-27T21:02:09+05:00::3::source.txt====sign====abc"
	testSig := signature.NewSignatureSha256FromFile()
	testSig.ParseString(signTxt)

	b, e := PackSignature(testSig)
	if e != nil {
		panic(e)
	}

	var ar signature.SignatureSha256
	e = UnpackSignature(&ar, b.Bytes())
	if e != nil {
		panic(e)
	}

	assert.True(t, testSig.Equal(&ar))
}
