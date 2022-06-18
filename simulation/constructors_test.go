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

// func TestFromBoard(t *testing.T) {
// 	examples := []struct {
// 		in    string
// 		color g4.Color
// 		out   Game
// 		err   error
// 	}{
// 		{
// 			in:    "UP yr/r//////",
// 			color: g4.Yellow,
// 			out: Game{
// 				board: Board{
// 					gravity: g4.UP,
// 					board:   exampleBoard,
// 				},
// 				colorWithTheMove: g4.Yellow,
// 			},
// 			err: nil,
// 		},
// 		{
// 			in:    "UP yr/r//////",
// 			color: g4.Empty,
// 			out:   Game{},
// 			err:   fmt.Errorf("unexpected color: %v", g4.Empty),
// 		},
// 		{
// 			in:    "UP yr/r//////",
// 			color: g4.Color(255),
// 			out:   Game{},
// 			err:   fmt.Errorf("unexpected color: %v", g4.Color(255)),
// 		},
// 	}
// 	for k, ex := range examples {
// 		board, err := FromString(ex.in)
// 		if err != nil {
// 			t.Errorf("example %d: error in FromString: %v", k, err)
// 		}
// 		out, err := FromBoard(board, ex.color)
// 		if err != nil && ex.err != nil {
// 			if err.Error() != ex.err.Error() {
// 				t.Errorf("example %d: invalid error: got %v but want %v", k, err, ex.err)
// 			}
// 		} else if err != nil || ex.err != nil {
// 			t.Errorf("example %d: invalid error: got %v but want %v", k, err, ex.err)
// 		}
// 		if out != ex.out {
// 			t.Errorf("example %d: invalid Game", k)
// 		}
// 	}
// }
