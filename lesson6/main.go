package main

import (
	"flag"
	"fmt"
	"log"
	"mtsbank_golang/lesson6/mycrypt"
	"mtsbank_golang/lesson6/signature"
	"mtsbank_golang/lesson6/signature/contract"
)

type Encoder interface {
	Encode() string
}

type FileEncoder struct {
	signature contract.Signature
}

type Decoder interface {
	Decode() bool
}

type Dec struct {
}

func (d *Dec) Decode() bool {
	return false
}

func main() {

	var outFile string

	fileSource := flag.String("source-file", "", "File source")
	flag.StringVar(&outFile, "out-file", "sign.txt", "File output")
	flag.Parse()

	args := flag.Args()
	action := args[0]

	fmt.Println(action, *fileSource, outFile)
	switch action {
	case "enc":

		encoder := mycrypt.NewEncoder(*fileSource, signature.NewSignatureSha256FromFile())
		err := encoder.EncryptSha256()
		if err != nil {
			panic(err)
		}
		err = encoder.SaveToFile(outFile)
		if err != nil {
			panic(err)
		}

	case "dec":
		return
	default:
		log.Fatal("Use enc of dec param")
	}
}
