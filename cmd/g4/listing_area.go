package main

import (
	"fmt"
	"g4"
)

type listingArea struct {
	History []g4.Move
	Waiting bool
}

func (m listingArea) View() string {
	s := ""
	for k, move := range m.History {
		if k%2 == 0 {
			s += fmt.Sprintf("%d. %s", k/2, move)
		} else {
			s += fmt.Sprintf(",%s\n", move)
		}
	}
	return s
}
