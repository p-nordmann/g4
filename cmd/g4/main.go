package main

import (
	"flag"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

var descr string

func init() {
	flag.Parse()
	descr = flag.Arg(0)
}

// TODO: rough draft/outline of main game loop
func playLoop(descr string) error {
	p := tea.NewProgram(NewFrontend(descr))
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func main() {
	playLoop(descr)
}
