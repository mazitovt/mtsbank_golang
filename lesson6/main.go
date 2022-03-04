package main

import (
	"errors"
	"fmt"
	"mtsbank_golang/lesson6/coder"
	"mtsbank_golang/lesson6/commands"
	"mtsbank_golang/lesson6/commands/contract"
	"mtsbank_golang/lesson6/signature"
	"os"
)

func ParseCommands(args []string, cmds ...contract.Command) (err error) {

	if len(args) < 1 {
		err = errors.New("invalid")
		return
	}

	subcmd := args[0]

	for _, cmd := range cmds {
		if cmd.Name() == subcmd {
			cmd.Init(args[1:])
			return cmd.Run()
		}
	}

	return errors.New("dfghf")
}

func main() {

	sig := signature.NewSignatureSha256FromFile()
	sigCompare := signature.NewSignatureSha256FromFile()

	decoder := coder.NewFileDecoder(sig, sigCompare)
	encoder := coder.NewFileEncoder(sig)

	cmd := []contract.Command{
		commands.NewEncodeCommand(encoder),
		commands.NewDecodeCommand(decoder),
	}

	if err := ParseCommands(os.Args[1:], cmd...); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
