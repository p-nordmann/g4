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
