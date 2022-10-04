package main

import (
	"bufio"
	"os"

	"github.com/alvii147/gitmsg/cli/cli"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	cli.Run(reader)
}
