package cmd

import (
	"errors"
	"os"

	s3cliclient "s3cli/client"
)

type putCmd struct {
	client s3cliclient.Client
}

func newPut(s3Client s3cliclient.Client) Cmd {
	return putCmd{client: s3Client}
}

func (cmd putCmd) Run(args []string) (err error) {
	if len(args) < 2 {
		return errors.New("Not enough arguments, expected source file and destination path")
	}

	source := args[0]
	destination := args[1]

	file, err := os.Open(source)
	if err != nil {
		return err
	}

	defer file.Close()

	return cmd.client.Put(destination, file)
}
