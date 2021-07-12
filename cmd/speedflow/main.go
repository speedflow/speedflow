package main

import (
	"os"

	"github.com/speedflow/speedflow/internal/speedflow/command"
)

func main() {
	command.Execute(os.Stdin, os.Stdout, os.Stderr)
}
