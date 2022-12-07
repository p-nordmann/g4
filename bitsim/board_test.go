package bitsim

import (
	"g4"
	"testing"
)

func TestBoardFromString(t *testing.T) {
	one := bitboard(1)
	ones8 := bitboard(255)
	examples := []struct {
		in  string
		out Board
	}{
		{
			in: "8|8|8|8|8|8|8|8",
			out: Board{
				yellowBits: 0,
				redBits:    0,
			},
		},
		{
			in: "y7|8|8|8|8|8|8|8",
			out: Board{
				yellowBits: one,
				redBits:    0,
			},
		},
		{
			in: "r6r|8|8|8|8|yyyyyyyy|8|8",
			out: Board{
				yellowBits: ones8 << 40,
				redBits:    one | one<<7,
			},
		},
		{
			in: "8|8|8|8|rrryr3|ryyyyyr1|r7|yr6",
			out: Board{
				yellowBits: (one<<3)<<32 | // column 4
					(one<<1|one<<2|one<<3|one<<4|one<<5)<<40 | // column 5
					(one)<<56, // column 7
				redBits: (one|one<<1|one<<2|one<<4)<<32 | // column 4
					(one|one<<6)<<40 | // column 5
					(one)<<48 | // column 6
					(one<<1)<<56, // column 7
			},
		},
	}
	for k, ex := range examples {
		if out, err := boardFromString(ex.in); err != nil || out != ex.out {
			t.Errorf("example %d: got (%v, %v) but want (%v, %v)", k, out, err, ex.out, nil)
		}
	}
}

func TestBoardFromStringError(t *testing.T) {
	examples := []string{
		"y7|ry6|r7|8|8|8|8",
		"yrryyrry|8|8|8|8|rrryyyyryr|8|8",
		"zry5|8|8|8|8|8|8|8",
	}
	for k, ex := range examples {
		if out, err := boardFromString(ex); err == nil {
			t.Errorf("example %d: got (%v, %v) but expected error", k, out, err)
		}
	}
}

// Tests that Board.String gives back the original string.
func TestBoardString(t *testing.T) {
	examples := []string{
		"8|8|8|8|8|8|8|8",
		"y7|8|8|8|8|8|8|8",
		"r6r|8|8|8|8|yyyyyyyy|8|8",
		"8|8|8|8|rrryr3|ryyyyyr1|r7|yr6",
		"y6y|8|8|8|8|yyy2ryy|8|8",
	}
	for _, ex := range examples {
		board, err := boardFromString(ex)
		if err != nil {
			t.Errorf("expected valid board but got error: %v", err)
		}
		s := board.String()
		if s != ex {
			t.Errorf("expected '%s' board but got '%s'", ex, s)
		}
	}
}

func TestBoardHeights(t *testing.T) {
	examples := []struct {
		in  string
		out [8]int
	}{
		{
			in:  "8|8|8|8|8|ryr5|r7|y7",
			out: [8]int{0, 0, 0, 0, 0, 3, 1, 1},
		},
		{
			in:  "8|ry6|8|8|rry5|8|8|8",
			out: [8]int{0, 2, 0, 0, 3, 0, 0, 0},
		},
		{
			in:  "8|8|8|8|rrryr3|ryyyyyr1|r7|yr6",
			out: [8]int{0, 0, 0, 0, 5, 7, 1, 2},
		},
		{
			in:  "8|yr6|yr6|yrrr4|yyr5|8|8|8",
			out: [8]int{0, 2, 2, 4, 3, 0, 0, 0},
		},
		{
			in:  "8|8|8|8|8|rryyyy2|r7|r7",
			out: [8]int{0, 0, 0, 0, 0, 6, 1, 1},
		},
		{
			in:  "8|ryryryry|8|8|rry5|8|8|8",
			out: [8]int{0, 8, 0, 0, 3, 0, 0, 0},
		},
	}
	for k, ex := range examples {
		in, err := boardFromString(ex.in)
		if err != nil {
			t.Errorf("example %d: FromString returned an error %v", k, err)
		}
		if in.heights() != ex.out {
			t.Errorf("example %d: got %v but want %v", k, in.heights(), ex.out)
		}
	}
}

func TestBoardHasYellowConnect4(t *testing.T) {
	examples := []struct {
		in  string
		out bool
	}{
		{
			in:  "8|8|rryyyyrr|rrrr4|yyy5|y7|8|8",
			out: true,
		},
		{
			in:  "8|ry6|ry6|yyrr4|ry6|8|8|8",
			out: true,
		},
		{
			in:  "8|8|ryry4|yyyrrr2|ryrr4|y7|r7|8",
			out: true,
		},
		{
			in:  "y7|ry6|rry5|rrry4|yyy5|8|8|8",
			out: true,
		},
		{
			in:  "8|8|ryry4|yyyrrr2|yrrr4|y7|r7|8",
			out: false,
		},
		{
			in:  "y7|ry6|rry5|8|yyy5|8|8|8",
			out: false,
		},
	}
	for k, ex := range examples {
		b, _ := boardFromString(ex.in)
		if b.hasYellowConnect4() != ex.out {
			t.Errorf("example %d: got %v but want %v", k, b.hasYellowConnect4(), ex.out)
		}
	}
}

func TestBoardHasRedConnect4(t *testing.T) {
	examples := []struct {
		in  string
		out bool
	}{
		{
			in:  "8|8|rryyyyrr|rrrr4|yyy5|y7|8|8",
			out: true,
		},
		{
			in:  "8|ry6|ry6|ryyr4|ry6|8|8|8",
			out: true,
		},
		{
			in:  "8|8|ryrr4|yyryrr2|yryy4|r7|r7|8",
			out: true,
		},
		{
			in:  "y7|ry6|rrr5|ryrr4|yyyr4|8|8|8",
			out: true,
		},
		{
			in:  "8|8|ryry4|yyyrrr2|yrrr4|y7|r7|8",
			out: false,
		},
		{
			in:  "y7|ry6|rry5|8|yyy5|8|8|8",
			out: false,
		},
	}
	for k, ex := range examples {
		b, _ := boardFromString(ex.in)
		if b.hasRedConnect4() != ex.out {
			t.Errorf("example %d: got %v but want %v", k, b.hasRedConnect4(), ex.out)
		}
	}
}

func TestBoardRotateLeft(t *testing.T) {
	examples := []struct {
		in    string
		times int
		out   string
	}{
		{
			in:    "yr6|8|8|8|8|8|8|8",
			times: 1,
			out:   "8|8|8|8|8|8|r7|y7",
		},
		{
			in:    "yr6|8|8|8|8|8|8|8",
			times: 2,
			out:   "8|8|8|8|8|8|8|6ry",
		},
		{
			in:    "yr6|8|8|8|8|8|8|8",
			times: 3,
			out:   "7y|7r|8|8|8|8|8|8",
		},
		{
			in:    "yr6|8|8|8|8|8|8|8",
			times: 4,
			out:   "yr6|8|8|8|8|8|8|8",
		},
		{
			in:    "yr6|2y5|2r5|2r5|2y5|6r1|6y1|8",
			times: 1,
			out:   "8|5ry1|8|8|8|1yrry3|r7|y7",
		},
		{
			in:    "yr6|2y5|2r5|2r5|2y5|6r1|6y1|8",
			times: 2,
			out:   "8|1y6|1r6|5y2|5r2|5r2|5y2|6ry",
		},
		{
			in:    "y7|ry6|y1y5|r2y4|y3r3|r4r2|y5r1|rrrrrrrr",
			times: 1,
			out:   "7r|6rr|5r1r|4r2r|3y3r|2y4r|1y5r|yryryryr",
		},
		{
			in:    "y7|ry6|y1y5|r2y4|y3r3|r4r2|y5r1|rrrrrrrr",
			times: 2,
			out:   "rrrrrrrr|1r5y|2r4r|3r3y|4y2r|5y1y|6yr|7y",
		},
	}
	for k, ex := range examples {
		got, _ := boardFromString(ex.in)
		got = got.RotateLeft(ex.times)
		want, _ := boardFromString(ex.out)
		if got != want {
			t.Errorf("example %d: got != want", k)
		}
	}
}

func TestBoardApplyGravity(t *testing.T) {
	examples := []struct {
		in  string
		out string
	}{
		{
			in:  "y6r|y5r1|r4y2|rr3y2|rrry4|7y|r6r|8",
			out: "yr6|yr6|ry6|rry5|rrry4|y7|rr6|8",
		},
		{
			in:  "y1r1y1r1|1y1r1y1r|r1y1r1y1|8|rr1yy1rr|1rr1yy1r|8|r2y2ry",
			out: "yryr4|yryr4|ryry4|8|rryyrr2|rryyr3|8|ryry4",
		},
	}
	for k, ex := range examples {
		got, _ := boardFromString(ex.in)
		got = got.ApplyGravity()
		want, _ := boardFromString(ex.out)
		if got != want {
			t.Errorf("example %d: got != want", k)
		}
	}
}

func TestBoardAddToken(t *testing.T) {
	examples := []struct {
		in     string
		column int
		color  g4.Color
		out    string
	}{
		{
			in:     "8|8|8|8|8|8|8|8",
			column: 0,
			color:  g4.Yellow,
			out:    "y7|8|8|8|8|8|8|8",
		},
		{
			in:     "yyy5|8|8|8|8|8|8|8",
			column: 0,
			color:  g4.Red,
			out:    "yyyr4|8|8|8|8|8|8|8",
		},
		{
			in:     "ryryryry|8|8|8|8|8|8|8",
			column: 0,
			color:  g4.Red,
			out:    "ryryryry|8|8|8|8|8|8|8",
		},
		{
			in:     "ryryryry|rr6|yy6|8|8|8|8|8",
			column: 1,
			color:  g4.Yellow,
			out:    "ryryryry|rry5|yy6|8|8|8|8|8",
		},
	}
	for k, ex := range examples {
		got, _ := boardFromString(ex.in)
		got = got.AddToken(ex.column, ex.color)
		want, _ := boardFromString(ex.out)
		if got != want {
			t.Errorf("example %d: got != want", k)
		}
	}
}
