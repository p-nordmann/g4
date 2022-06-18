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

// MoveType indicates the kind of a move.
//
// A move can be one of:
//	- Tilt: changing the direction of the gravity.
//	- Token: placing a new token on top of one column.
type MoveType byte

const (
	Token MoveType = iota
	Tilt
)

type Move struct {
	// MoveType indicates whether the move is a tilt or a token.
	moveType MoveType
	// Gravity indicates the new direction for a tilt move.
	gravity Direction
	// Column indicates the column that was played for a token move.
	column int
	// Color indicates the color that was played for a token move.
	color Color
}

func Base() Move {
	return Move{}
}

func (m Move) Token() Move {
	m.moveType = Token
	return m
}

func (m Move) Tilt() Move {
	m.moveType = Tilt
	return m
}

func (m Move) Column(column int) Move {
	m.column = column
	return m
}

func (m Move) Color(color Color) Move {
	m.color = color
	return m
}

func (m Move) Gravity(gravity Direction) Move {
	m.gravity = gravity
	return m
}

func TokenMove(color Color, column int) Move {
	return Base().Token().Color(color).Column(column)
}

func TiltMove(direction Direction) Move {
	return Base().Tilt().Gravity(direction)
}
