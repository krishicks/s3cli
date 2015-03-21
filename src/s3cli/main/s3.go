package main

import (
	"fmt"
	"os"
	"runtime"

	s3cliapp "s3cli/app"
)

func main() {
	// set the number of processors to use to the number of cpus for parallelization of concurrent transfers
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := s3cliapp.New()

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}
