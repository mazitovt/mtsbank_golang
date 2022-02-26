package commands

import (
	"flag"
	"mtsbank_golang/lesson6/coder/contract"
)

type DecodeCommand struct {
	fs *flag.FlagSet

	decoder contract.Decoder

	source      string
	signature   string
	destination string
}

func (dc *DecodeCommand) Init(args []string) error {
	return dc.fs.Parse(args)
}

func (dc *DecodeCommand) Name() string {
	return dc.fs.Name()
}

func (dc *DecodeCommand) Run() (err error) {

	err = dc.decoder.Decode(dc.source, dc.signature)
	if err != nil {
		return
	}
	err = dc.decoder.SaveToFile(dc.destination)
	if err != nil {
		return
	}

	return
}

func NewDecodeCommand(decoder contract.Decoder) *DecodeCommand {
	ec := &DecodeCommand{fs: flag.NewFlagSet("dec", flag.ContinueOnError), decoder: decoder}
	ec.fs.StringVar(&ec.source, "action", "source.txt", "Source file")
	ec.fs.StringVar(&ec.signature, "signature", "sign.txt", "Signature file")
	ec.fs.StringVar(&ec.destination, "destination", "out.txt", "Destination file")

	return ec
}
