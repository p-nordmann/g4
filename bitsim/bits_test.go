package bitsim_test

import (
	bitsim "g4/bitsim"
	"testing"
)

func TestHeights(t *testing.T) {
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
		in, err := bitsim.FromString(ex.in)
		if err != nil {
			t.Errorf("example %d: FromString returned an error %v", k, err)
		}
		if in.Heights() != ex.out {
			t.Errorf("example %d: got %v but want %v", k, in.Heights(), ex.out)
		}
	}
}

func TestHasYellowConnect4(t *testing.T) {
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
		b, _ := bitsim.FromString(ex.in)
		if b.HasYellowConnect4() != ex.out {
			t.Errorf("example %d: got %v but want %v", k, b.HasYellowConnect4(), ex.out)
		}
	}
}

func TestHasRedConnect4(t *testing.T) {
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
		b, _ := bitsim.FromString(ex.in)
		if b.HasRedConnect4() != ex.out {
			t.Errorf("example %d: got %v but want %v", k, b.HasRedConnect4(), ex.out)
		}
	}
}
