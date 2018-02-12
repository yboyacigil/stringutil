package stringutil

import "testing"

func TestScramble(t *testing.T) {
	cases := []struct {
		input, firstLetter, lastLetter, inBtwNotEq string
		number                                     bool
	}{
		{"hello", "h", "o", "ell", false},
		{"daddy", "d", "y", "add", false},
		{"according", "a", "g", "ccordin", false},
		{"mom", "m", "m", "o", false},
		{"1234", "1", "4", "23", true},
	}

	for _, c := range cases {
		actual := Scramble(c.input)
		t.Logf("Scramble(%q) == %q", c.input, actual)

		if actual[0:1] != c.firstLetter {
			t.Errorf("Scramble(%q) == %q, actual first letter is %q but expected %q", c.input, actual, actual[0:1], c.firstLetter)
		}
		if actual[len(actual)-1:] != c.lastLetter {
			t.Errorf("Scramble(%q) == %q, actual last letter is %q but expected %q", c.input, actual, actual[len(actual)-1:], c.lastLetter)
		}
		inBtwn := string(actual[1 : len(actual)-1])
		if len(inBtwn) > 1 && !c.number && inBtwn == c.inBtwNotEq {
			t.Errorf("Scramble(%q) == %q, actual in between should be different than %q", c.input, actual, c.inBtwNotEq)
		}
		if c.number && inBtwn != c.inBtwNotEq {
			t.Errorf("Scramble(%q) == %q, actual in between should NOT be different than %q", c.input, actual, c.inBtwNotEq)
		}
	}
}
