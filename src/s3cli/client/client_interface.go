package client

import (
	"os"
)

type Client interface {
	Put(string, *os.File) error
}
