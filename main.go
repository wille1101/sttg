package main

import (
	"fmt"
	"os"

	"github.com/wille1101/sttg/config"
	"github.com/wille1101/sttg/tui"
)

func main() {
	if err := config.LoadCon(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	p := tui.NewProgram()
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
