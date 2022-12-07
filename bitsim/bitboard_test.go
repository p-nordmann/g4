package bitsim

import (
	"testing"
)

func TestBitboardFromString(t *testing.T) {
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

func TestBitboardFromStringError(t *testing.T) {
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
		if in.north() != out {
			t.Errorf("example %d: got %v but want %v", k, in.north(), out)
		}
	}
}

func TestNorth2(t *testing.T) {
	examples := []struct {
		in  string
		out string
	}{
		{
			in:  "x7|8|8|8|8|8|8|8",
			out: "2x5|8|8|8|8|8|8|8",
		},
		{
			in:  "6xx|8|8|8|8|8|8|5xxx",
			out: "8|8|8|8|8|8|8|7x",
		},
		{
			in:  "8|1x6|8|8|8|8|8|8",
			out: "8|3x4|8|8|8|8|8|8",
		},
		{
			in:  "8|8|3x4|8|8|8|8|8",
			out: "8|8|5x2|8|8|8|8|8",
		},
	}
	for k, ex := range examples {
		in, _ := bitboardFromString(ex.in)
		out, _ := bitboardFromString(ex.out)
		if in.north2() != out {
			t.Errorf("example %d: got %v but want %v", k, in.north2(), out)
		}
	}
}

func TestNorth3(t *testing.T) {
	examples := []struct {
		in  string
		out string
	}{
		{
			in:  "x7|8|8|8|8|8|8|8",
			out: "3x4|8|8|8|8|8|8|8",
		},
		{
			in:  "5xxx|8|8|8|8|8|8|4xxxx",
			out: "8|8|8|8|8|8|8|7x",
		},
		{
			in:  "8|1x6|8|8|8|8|8|8",
			out: "8|4x3|8|8|8|8|8|8",
		},
		{
			in:  "8|8|3x4|8|8|8|8|8",
			out: "8|8|6x1|8|8|8|8|8",
		},
	}
	for k, ex := range examples {
		in, _ := bitboardFromString(ex.in)
		out, _ := bitboardFromString(ex.out)
		if in.north3() != out {
			t.Errorf("example %d: got %v but want %v", k, in.north3(), out)
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
		if in.south() != out {
			t.Errorf("example %d: got %v but want %v", k, in.south(), out)
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
		if in.east() != out {
			t.Errorf("example %d: got %v but want %v", k, in.east(), out)
		}
	}
}

func TestEast2(t *testing.T) {
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
			out: "8|8|x7|8|8|8|8|8",
		},
		{
			in:  "8|8|8|8|8|1x6|8|8",
			out: "8|8|8|8|8|8|8|1x6",
		},
		{
			in:  "8|8|8|8|8|8|x7|3x4",
			out: "8|8|8|8|8|8|8|8",
		},
	}
	for k, ex := range examples {
		in, _ := bitboardFromString(ex.in)
		out, _ := bitboardFromString(ex.out)
		if in.east2() != out {
			t.Errorf("example %d: got %v but want %v", k, in.east2(), out)
		}
	}
}

func TestEast3(t *testing.T) {
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
			out: "8|8|8|x7|8|8|8|8",
		},
		{
			in:  "8|8|8|8|1x6|1x6|8|8",
			out: "8|8|8|8|8|8|8|1x6",
		},
		{
			in:  "8|8|8|8|8|7x|x7|3x4",
			out: "8|8|8|8|8|8|8|8",
		},
	}
	for k, ex := range examples {
		in, _ := bitboardFromString(ex.in)
		out, _ := bitboardFromString(ex.out)
		if in.east3() != out {
			t.Errorf("example %d: got %v but want %v", k, in.east3(), out)
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
		if in.northWest() != out {
			t.Errorf("example %d: got %v but want %v", k, in.northWest(), out)
		}
	}
}

func TestNorthWest2(t *testing.T) {
	examples := []struct {
		in  string
		out string
	}{
		{
			in:  "x7|8|8|8|x7|8|8|8",
			out: "8|8|2x5|8|8|8|8|8",
		},
		{
			in:  "8|x7|8|3x4|8|5xxx|8|8",
			out: "8|5x2|8|7x|8|8|8|8",
		},
		{
			in:  "8|8|x7|8|8|8|8|8",
			out: "2x5|8|8|8|8|8|8|8",
		},
	}
	for k, ex := range examples {
		in, _ := bitboardFromString(ex.in)
		out, _ := bitboardFromString(ex.out)
		if in.northWest2() != out {
			t.Errorf("example %d: got %v but want %v", k, in.northWest2(), out)
		}
	}
}

func TestNorthWest3(t *testing.T) {
	examples := []struct {
		in  string
		out string
	}{
		{
			in:  "x7|8|8|8|x7|8|8|8",
			out: "8|3x4|8|8|8|8|8|8",
		},
		{
			in:  "8|x7|8|3x4|8|4xxxx|8|8",
			out: "6x1|8|7x|8|8|8|8|8",
		},
		{
			in:  "8|8|8|x7|8|8|8|8",
			out: "3x4|8|8|8|8|8|8|8",
		},
	}
	for k, ex := range examples {
		in, _ := bitboardFromString(ex.in)
		out, _ := bitboardFromString(ex.out)
		if in.northWest3() != out {
			t.Errorf("example %d: got %v but want %v", k, in.northWest3(), out)
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
		if in.northEast() != out {
			t.Errorf("example %d: got %v but want %v", k, in.northEast(), out)
		}
	}
}

func TestNorthEast2(t *testing.T) {
	examples := []struct {
		in  string
		out string
	}{
		{
			in:  "8|8|8|8|8|x5x1|x7|x7",
			out: "8|8|8|8|8|8|8|2x5",
		},
		{
			in:  "x7|8|8|8|8|7x|8|8",
			out: "8|8|2x5|8|8|8|8|8",
		},
		{
			in:  "8|8|8|8|8|4xxxx|8|8",
			out: "8|8|8|8|8|8|8|6xx",
		},
	}
	for k, ex := range examples {
		in, _ := bitboardFromString(ex.in)
		out, _ := bitboardFromString(ex.out)
		if in.northEast2() != out {
			t.Errorf("example %d: got %v but want %v", k, in.northEast2(), out)
		}
	}
}

func TestNorthEast3(t *testing.T) {
	examples := []struct {
		in  string
		out string
	}{
		{
			in:  "8|8|8|8|8|x5x1|x7|x7",
			out: "8|8|8|8|8|8|8|8",
		},
		{
			in:  "x7|8|8|8|8|7x|8|8",
			out: "8|8|8|3x4|8|8|8|8",
		},
		{
			in:  "8|8|8|8|4xxxx|8|8|8",
			out: "8|8|8|8|8|8|8|7x",
		},
	}
	for k, ex := range examples {
		in, _ := bitboardFromString(ex.in)
		out, _ := bitboardFromString(ex.out)
		if in.northEast3() != out {
			t.Errorf("example %d: got %v but want %v", k, in.northEast3(), out)
		}
	}
}

func TestBitboardHasConnect4(t *testing.T) {
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
		if in.hasConnect4() != ex.out {
			t.Errorf("example %d: got %v but want %v", k, in.hasConnect4(), ex.out)
		}
	}
}

func TestBitboardCount(t *testing.T) {
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
		if in.count() != ex.out {
			t.Errorf("example %d: got %v but want %v", k, in.count(), ex.out)
		}
	}
}

func TestBitboardGetColumn(t *testing.T) {
	examples := []struct {
		in     string
		column int
		out    byte
	}{
		{
			in:     "8|xx6|8|8|8|8|8|8",
			column: 0,
			out:    0,
		},
		{
			in:     "8|xx6|8|8|8|8|8|8",
			column: 1,
			out:    0b00000011,
		},
		{
			in:     "8|xx6|x1x1x1x1|8|8|8|8|8",
			column: 2,
			out:    0b01010101,
		},
		{
			in:     "8|xx6|8|8|8|8|8|xxxxxxxx",
			column: 7,
			out:    255,
		},
	}
	for k, ex := range examples {
		bb, _ := bitboardFromString(ex.in)
		out := bb.getColumn(ex.column)
		if out != ex.out {
			t.Errorf("example %d: got %v but want %v", k, out, ex.out)
		}
	}
}

func TestBitboardRotateLeft(t *testing.T) {
	examples := []struct {
		in  string
		out string
	}{
		{
			in:  "x7|8|8|8|8|8|8|8",
			out: "8|8|8|8|8|8|8|x7",
		},
		{
			in:  "1x6|8|8|8|8|8|8|8",
			out: "8|8|8|8|8|8|x7|8",
		},
		{
			in:  "x7|xx6|x7|8|8|8|8|8",
			out: "8|8|8|8|8|8|1x6|xxx5",
		},
		{
			in:  "xx6|8|8|8|8|8|8|8",
			out: "8|8|8|8|8|8|x7|x7",
		},
		{
			in:  "xx6|2x5|2x5|2x5|2x5|6x1|6x1|8",
			out: "8|5xx1|8|8|8|1xxxx3|x7|x7",
		},
		{
			in:  "x7|xx6|x1x5|x2x4|x3x3|x4x2|x5x1|xxxxxxxx",
			out: "7x|6xx|5x1x|4x2x|3x3x|2x4x|1x5x|xxxxxxxx",
		},
	}
	for k, ex := range examples {
		got, _ := bitboardFromString(ex.in)
		got = got.rotateLeft()
		want, _ := bitboardFromString(ex.out)
		if got != want {
			t.Errorf("example %d: got %v but want %v", k, got, want)
		}
	}
}
