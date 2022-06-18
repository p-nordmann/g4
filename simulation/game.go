/*
G4 is an open-source board game inspired by the popular game of connect-4.
Copyright (C) 2022  Pierre-Louis Nordmann

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package simulation

import (
	"fmt"
	"g4"
)

type Game struct {
	board     g4.Board
	color     g4.Color
	direction g4.Direction
}

func FromBoard(board g4.Board, color g4.Color, direction g4.Direction) (Game, error) {
	switch color {
	case g4.Red:
	case g4.Yellow:
	default:
		return Game{}, fmt.Errorf("unexpected color: %v", color)
	}
	return Game{
		board:     board,
		color:     color,
		direction: direction,
	}, nil
}

// Generate computes the list of possible moves from a given position.
func (g Game) Generate() ([]g4.Move, error) {
	var moves []g4.Move

	// Look for connect 4s.
	hasYellowConnect4 := g.board.HasYellowConnect4()
	hasRedConnect4 := g.board.HasRedConnect4()
	if hasYellowConnect4 && hasRedConnect4 {
		return moves, g4.ErrorGameOver(g4.Empty) // Draw.
	} else if hasYellowConnect4 {
		return moves, g4.ErrorGameOver(g4.Yellow)
	} else if hasRedConnect4 {
		return moves, g4.ErrorGameOver(g4.Red)
	}

	// Check whether board is full.
	heights := g.board.Heights()
	if heights[0] == 8 && heights[1] == 8 && heights[2] == 8 && heights[3] == 8 && heights[4] == 8 && heights[5] == 8 && heights[6] == 8 && heights[7] == 8 {
		return moves, g4.ErrorGameOver(g4.Empty) // Draw.
	}

	// Tilt moves.
	if g.direction != g4.UP {
		moves = append(moves, g4.TiltMove(g4.UP))
	}
	if g.direction != g4.LEFT {
		moves = append(moves, g4.TiltMove(g4.LEFT))
	}
	if g.direction != g4.DOWN {
		moves = append(moves, g4.TiltMove(g4.DOWN))
	}
	if g.direction != g4.RIGHT {
		moves = append(moves, g4.TiltMove(g4.RIGHT))
	}

	// Token moves.
	for column, height := range heights {
		if height < 8 {
			moves = append(moves, g4.TokenMove(g.color, column))
		}
	}

	return moves, nil
}

func (g Game) Apply(move g4.Move) (g4.Game, error) {
	return nil, nil
}
