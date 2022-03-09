package main

import (
	"mtsbank_golang/file_signature/packing"
	"mtsbank_golang/file_signature/signature"
	"testing"
)

func BenchmarkReflectPack(b *testing.B) {

	sig1 := signature.NewSignatureSha256FromFile()
	sig1.ParseString("2022-02-27T21:02:09+05:00::11::sourceName.txt====sign====a����3\u001B+���g}�H�(\u0016:A�5)ވ�r1���\u0004")

	for i := 0; i < b.N; i++ {
		_, _ = packing.PackSignature(sig1)
	}
}

func BenchmarkReflectUnpack(b *testing.B) {

	sig1 := signature.NewSignatureSha256FromFile()
	by := []byte{
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

	for i := 0; i < b.N; i++ {
		_ = packing.UnpackSignature(sig1, by)
	}
}

func BenchmarkCodegenUnpack(b *testing.B) {

	sig1 := signature.NewSignatureSha256FromFile()
	by := []byte{
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

	for i := 0; i < b.N; i++ {
		_ = sig1.Unpack(by)
	}
}
