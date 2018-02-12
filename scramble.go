package stringutil

import (
	"math/rand"
	"regexp"
)

var wordsRegexp, _ = regexp.Compile("([\\w]+)")
var onlyDigitsRegexp, _ = regexp.Compile("([\\d]+)")

// Scramble returns its argument string keeping first and last letters untouched
func Scramble(s string) string {
	in := []byte(s)
	out := wordsRegexp.ReplaceAllFunc(in, scrambleByRandom)
	return string(out)
}

func scrambleByRandom(s []byte) []byte {
	if len(s) <= 3 {
		return s
	}

	if onlyDigitsRegexp.MatchString(string(s)) {
		return s
	}

	inBtw := []byte(s[1 : len(s)-1])

	x := make([]int, len(inBtw))
	for i := range x {
		x[i] = i
	}

	i := 0
	for i < len(x) {
		j := rand.Intn(len(x))
		used := false
		k := i - 1
		if i == 0 {
			k = i
		}
		for !used && k >= 0 {
			if x[k] == j {
				used = true
			}
			k = k - 1
		}
		if !used {
			x[i] = j
			i = i + 1
		}
	}

	r := make([]byte, len(x)+2)
	r[0] = s[0]
	r[len(r)-1] = s[len(s)-1]
	for i := range x {
		r[i+1] = inBtw[x[i]]
	}

	return r
}
