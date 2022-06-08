package bits_test

import (
	b "g4/bits"
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
		in, err := b.FromString(ex.in)
		if err != nil {
			t.Errorf("example %d: FromString returned an error %v", k, err)
		}
		if in.Heights() != ex.out {
			t.Errorf("example %d: got %v but want %v", k, in.Heights(), ex.out)
		}
	}
}
