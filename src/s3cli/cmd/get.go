package cmd

import (
	"errors"
	s3cliclient "s3cli/client"
)

type getCmd struct {
	client s3cliclient.Client
}

func newGet(s3Client s3cliclient.Client) (cmd Cmd) {
	return getCmd{client: s3Client}
}

func (cmd getCmd) Run(args []string) (err error) {
	if len(args) < 2 {
		return errors.New("Not enough arguments, expected remote path and destination path")
	}

	// remotePath := args[0]
	// localPath := args[1]

	return errors.New("not implemented")
}
