package bits

import (
	"testing"
)

func TestFromString(t *testing.T) {
	one := bitboard(1)
	ones8 := bitboard(255)
	examples := []struct {
		in  string
		out bitboard
	}{
		{
			in:  "8|8|8|8|8|8|8|8",
			out: 0,
		},
		{
			in:  "x7|8|8|8|8|8|8|8",
			out: one,
		},
		{
			in:  "x6x|8|8|8|8|xxxxxxxx|8|8",
			out: one | one<<7 | ones8<<40,
		},
		{
			in:  "x321x|53|44|8|8|xxxxxxxx|8|111111111",
			out: one | one<<7 | ones8<<40,
		},
	}
	for k, ex := range examples {
		if out, err := bitboardFromString(ex.in); err != nil || out != ex.out {
			t.Errorf("example %d: got (%v, %v) but want (%v, %v)", k, out, err, ex.out, nil)
		}
	}
}

func TestFromStringError(t *testing.T) {
	examples := []string{
		"x7|8|8|8|8|8|8",
		"x6x|8|8|8|8|xxxxxxxxx|8|8",
		"z7|8|8|8|8|8|8|8",
	}
	for k, ex := range examples {
		if out, err := bitboardFromString(ex); err == nil {
			t.Errorf("example %d: got (%v, %v) but expected error", k, out, err)
		}
	}
}

func TestNorth(t *testing.T) {
	examples := []struct {
		in  string
		out string
	}{
		{
			in:  "x7|8|8|8|8|8|8|8",
			out: "1x6|8|8|8|8|8|8|8",
		},
		{
			in:  "7x|8|8|8|8|8|8|8",
			out: "8|8|8|8|8|8|8|8",
		},
		{
			in:  "8|1x6|8|8|8|8|8|8",
			out: "8|2x5|8|8|8|8|8|8",
		},
		{
			in:  "8|8|3x4|8|8|8|8|8",
			out: "8|8|4x3|8|8|8|8|8",
		},
	}
	for k, ex := range examples {
		in, _ := bitboardFromString(ex.in)
		out, _ := bitboardFromString(ex.out)
		if in.North() != out {
			t.Errorf("example %d: got %v but want %v", k, in.North(), out)
		}
	}
}

func TestWest(t *testing.T) {
	examples := []struct {
		in  string
		out string
	}{
		{
			in:  "x7|8|8|8|8|8|8|8",
			out: "8|8|8|8|8|8|8|8",
		},
		{
			in:  "8|x7|8|8|8|8|8|8",
			out: "x7|8|8|8|8|8|8|8",
		},
		{
			in:  "8|8|8|8|8|8|1x6|8",
			out: "8|8|8|8|8|1x6|8|8",
		},
		{
			in:  "8|8|8|8|8|8|8|8",
			out: "8|8|8|8|8|8|8|8",
		},
	}
	for k, ex := range examples {
		in, _ := bitboardFromString(ex.in)
		out, _ := bitboardFromString(ex.out)
		if in.West() != out {
			t.Errorf("example %d: got %v but want %v", k, in.West(), out)
		}
	}
}

func TestSouth(t *testing.T) {
	examples := []struct {
		in  string
		out string
	}{
		{
			in:  "x7|8|8|8|x7|8|8|8",
			out: "8|8|8|8|8|8|8|8",
		},
		{
			in:  "7x|8|8|8|8|8|8|8",
			out: "6x1|8|8|8|8|8|8|8",
		},
		{
			in:  "8|1x6|8|8|2x5|8|8|8",
			out: "8|x7|8|8|1x6|8|8|8",
		},
		{
			in:  "8|8|xx6|8|8|8|8|8",
			out: "8|8|x7|8|8|8|8|8",
		},
	}
	for k, ex := range examples {
		in, _ := bitboardFromString(ex.in)
		out, _ := bitboardFromString(ex.out)
		if in.South() != out {
			t.Errorf("example %d: got %v but want %v", k, in.South(), out)
		}
	}
}

func TestEast(t *testing.T) {
	examples := []struct {
		in  string
		out string
	}{
		{
			in:  "8|8|8|8|8|8|8|8",
			out: "8|8|8|8|8|8|8|8",
		},
		{
			in:  "x7|8|8|8|8|8|8|8",
			out: "8|x7|8|8|8|8|8|8",
		},
		{
			in:  "8|8|8|8|8|1x6|8|8",
			out: "8|8|8|8|8|8|1x6|8",
		},
		{
			in:  "8|8|8|8|8|8|8|3x4",
			out: "8|8|8|8|8|8|8|8",
		},
	}
	for k, ex := range examples {
		in, _ := bitboardFromString(ex.in)
		out, _ := bitboardFromString(ex.out)
		if in.East() != out {
			t.Errorf("example %d: got %v but want %v", k, in.East(), out)
		}
	}
}

func TestNorthWest(t *testing.T) {
	examples := []struct {
		in  string
		out string
	}{
		{
			in:  "x7|8|8|8|x7|8|8|8",
			out: "8|8|8|1x6|8|8|8|8",
		},
		{
			in:  "8|x7|8|3x4|8|6xx|8|8",
			out: "1x6|8|4x3|8|7x|8|8|8",
		},
	}
	for k, ex := range examples {
		in, _ := bitboardFromString(ex.in)
		out, _ := bitboardFromString(ex.out)
		if in.NorthWest() != out {
			t.Errorf("example %d: got %v but want %v", k, in.NorthWest(), out)
		}
	}
}

func TestNorthEast(t *testing.T) {
	examples := []struct {
		in  string
		out string
	}{
		{
			in:  "8|8|8|8|8|x5x1|x7|x7",
			out: "8|8|8|8|8|8|1x5x|1x6",
		},
		{
			in:  "x7|8|8|8|8|7x|8|8",
			out: "8|1x6|8|8|8|8|8|8",
		},
	}
	for k, ex := range examples {
		in, _ := bitboardFromString(ex.in)
		out, _ := bitboardFromString(ex.out)
		if in.NorthEast() != out {
			t.Errorf("example %d: got %v but want %v", k, in.NorthEast(), out)
		}
	}
}

func TestHasConnect4(t *testing.T) {
	examples := []struct {
		in  string
		out bool
	}{
		{
			in:  "8|8|8|8|8|x5x1|x7|x7",
			out: false,
		},
		{
			in:  "8|1x6|8|8|2x5|8|8|8",
			out: false,
		},
		{
			in:  "8|8|8|8|xxx1x3|x5x1|x7|1x6",
			out: false,
		},
		{
			in:  "8|1x6|1x6|1xxx4|2x5|8|8|8",
			out: false,
		},
		{
			in:  "8|8|8|8|8|2xxxx2|x7|x7",
			out: true,
		},
		{
			in:  "8|4xxxx|8|8|2x5|8|8|8",
			out: true,
		},
		{
			in:  "8|8|x7|1x6|2x5|3x4|x7|x7",
			out: true,
		},
		{
			in:  "8|7x|6xx|5x2|4x3|8|8|8",
			out: true,
		},
		{
			in:  "8|7x|7x|5x2|4x3|7x|7x|7x",
			out: false,
		},
		{
			in:  "8|7x|7x|5x2|7x|7x|7x|7x",
			out: true,
		},
	}
	for k, ex := range examples {
		in, _ := bitboardFromString(ex.in)
		if in.HasConnect4() != ex.out {
			t.Errorf("example %d: got %v but want %v", k, in.HasConnect4(), ex.out)
		}
	}
}

func TestCount(t *testing.T) {
	examples := []struct {
		in  string
		out int
	}{
		{
			in:  "8|8|8|8|8|x5x1|x7|x7",
			out: 4,
		},
		{
			in:  "8|1x6|8|8|2x5|8|8|8",
			out: 2,
		},
		{
			in:  "8|8|8|8|xxx1x3|x5x1|x7|1x6",
			out: 8,
		},
		{
			in:  "8|1x6|1x6|1xxx4|2x5|8|8|8",
			out: 6,
		},
		{
			in:  "8|8|8|8|8|2xxxx2|x7|x7",
			out: 6,
		},
		{
			in:  "8|4xxxx|8|8|2x5|8|8|8",
			out: 5,
		},
	}
	for k, ex := range examples {
		in, _ := bitboardFromString(ex.in)
		if in.Count() != ex.out {
			t.Errorf("example %d: got %v but want %v", k, in.Count(), ex.out)
		}
	}
}
