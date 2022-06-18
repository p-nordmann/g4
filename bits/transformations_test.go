package bits_test

import (
	"g4"
	"g4/bits"
	"testing"
)

func TestRotateLeft(t *testing.T) {
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
		got, _ := bits.FromString(ex.in)
		got = got.RotateLeft(ex.times).(bits.Board)
		want, _ := bits.FromString(ex.out)
		if got != want {
			t.Errorf("example %d: got != want", k)
		}
	}
}

func TestApplyGravity(t *testing.T) {
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
		got, _ := bits.FromString(ex.in)
		got = got.ApplyGravity().(bits.Board)
		want, _ := bits.FromString(ex.out)
		if got != want {
			t.Errorf("example %d: got != want", k)
		}
	}
}

func TestAddToken(t *testing.T) {
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
		got, _ := bits.FromString(ex.in)
		got = got.AddToken(ex.column, ex.color).(bits.Board)
		want, _ := bits.FromString(ex.out)
		if got != want {
			t.Errorf("example %d: got != want", k)
		}
	}
}
