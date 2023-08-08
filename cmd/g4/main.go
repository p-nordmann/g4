package main

import (
	"flag"
	"fmt"
	"g4"
	"g4/bitsim"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	flag.Parse()
	spec := flag.Arg(0)
	board, _ := bitsim.FromString(bitsim.StartingPosition)
	p := tea.NewProgram(
		AppModel{
			spec:       spec,
			game:       bitsim.Game{Board: board, Mover: g4.Yellow},
			keyHandler: KeyHandler{keyMap: defaultKeymap},
			gameStatus: inProgress,
			connStatus: connecting,
		},
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
	)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
	}
}
