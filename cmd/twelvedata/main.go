package main

import (
	"os"

	"github.com/twelvedata/twelvedata-cli/internal/commands"
)

func main() {
	os.Exit(commands.Execute())
}
