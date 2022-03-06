package commands

import (
	"flag"
	"mtsbank_golang/tests_mocks/coder/contract"
)

type EncodeCommand struct {
	fs *flag.FlagSet

	encoder     contract.Encoder
	source      string
	destination string
}

func (ec *EncodeCommand) Init(args []string) error {
	return ec.fs.Parse(args)
}

func (ec *EncodeCommand) Name() string {
	return ec.fs.Name()
}

func (ec *EncodeCommand) Run() (err error) {

	err = ec.encoder.Encode(ec.source)
	if err != nil {
		return
	}
	err = ec.encoder.SaveToFile(ec.destination)
	if err != nil {
		return
	}

	return
}

func NewEncodeCommand(encoder contract.Encoder) *EncodeCommand {
	ec := &EncodeCommand{fs: flag.NewFlagSet("enc", flag.ContinueOnError), encoder: encoder}
	ec.fs.StringVar(&ec.source, "source", "source.txt", "Source file")
	ec.fs.StringVar(&ec.destination, "destination", "sign.txt", "Destination file")

	return ec
}
