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

type ErrorGameOver Color

func (err ErrorGameOver) Error() string {
	switch Color(err) {
	case Yellow:
		return "game is over - yellow wins"
	case Red:
		return "game is over - red wins"
	default:
		return "game is over - draw"
	}
}

type ErrorInvalidMove struct{}

func (err ErrorInvalidMove) Error() string {
	return "invalid move"
}
