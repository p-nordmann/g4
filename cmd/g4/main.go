package main

import (
	"flag"
	"fmt"
	"g4"
	"g4/bitsim"

	tea "github.com/charmbracelet/bubbletea"
)

var descr string

func init() {
	flag.Parse()
	descr = flag.Arg(0)
}

func main() {
	board, _ := bitsim.FromString("ryry4|rrr5|yyy5|8|8|8|yyr5|4yyr1")
	p := tea.NewProgram(
		AppModel{
			descr:         descr,
			game:          bitsim.Game{Board: board, Mover: g4.Yellow},
			board:         BoardModel{game: bitsim.Game{Board: board, Mover: g4.Yellow}},
			preview:       BoardModel{game: bitsim.Game{Board: board, Mover: g4.Yellow}},
			selectedModel: selectedPreview,
		},
		tea.WithAltScreen(),
	)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
	}
}
