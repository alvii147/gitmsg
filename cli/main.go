package main

import (
	"bufio"
	"gitmsg/cli"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	cli.Run(reader)
}
