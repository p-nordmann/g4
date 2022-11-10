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

package g4

// Game is supposed to hold the data about the current game state.
//
// It can generate a list of possible move in its current state and it can be updated using moves.
//
// It can be queried for display information.
//
// TODO: display information.
type Game interface {
	// Apply performs a move from a game state.
	// The new game state is returned.
	Apply(m Move) (Game, error)

	// Generate returns the list of possible moves.
	// If the game is over, an error is supposed to be returned.
	Generate() ([]Move, error)

	ToArray() ([8][8]Color, Direction)
}

type Board interface {
	// RotateLeft applies `times` left rotations on the board.
	//
	// It does not make the token drop according to new gravity.
	RotateLeft(times int) Board

	// ApplyGravity makes the token drop according to gravity.
	ApplyGravity() Board

	// AddToken adds a token on top of requested column.
	AddToken(column int, color Color) Board

	// HasYellowConnect4 returns whether the board has a connect 4 with yellow tokens.
	HasYellowConnect4() bool

	// HasRedConnect4 returns whether the board has a connect 4 with red tokens.
	HasRedConnect4() bool

	// Heights returns a list of heights for all the columns.
	Heights() [8]int

	ToArray() [8][8]Color
}

// TODO: split board into multiple interfaces?
// TODO: add a method String to boards.
// TODO: add a method image or string or array to game.
// TODO: build game history with boards and moves.
