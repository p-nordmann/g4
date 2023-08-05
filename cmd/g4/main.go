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
	board, _ := bitsim.FromString("ryry4|rrr5|yyy5|8|8|8|yyr5|4yyr1")
	p := tea.NewProgram(
		AppModel{
			spec:       spec,
			game:       bitsim.Game{Board: board, Mover: g4.Yellow},
			keyHandler: KeyHandler{keyMap: defaultKeymap},
		},
		tea.WithAltScreen(),
	)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
	}
}
