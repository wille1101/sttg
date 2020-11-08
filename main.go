package main

import (
	"fmt"
	"os"

	"github.com/wille1101/sttg/tui"
)

func main() {
	p := tui.NewProgram()
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
