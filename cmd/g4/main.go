package main

import (
	"flag"
	"fmt"
	"g4/cmd/g4/frontend"

	tea "github.com/charmbracelet/bubbletea"
)

var port int
var url string

func init() {
	flag.IntVar(&port, "port", 80, "port to listen to")
	flag.Parse()
	url = flag.Arg(0)
}

// TODO: rough draft/outline of main game loop
func playLoop(url string, port int) error {

	p := tea.NewProgram(frontend.New(url, port), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func main() {
	fmt.Println("listening to:", fmt.Sprintf(":%d", port))
	fmt.Println("reaching at:", url)
	playLoop(url, port)
}
