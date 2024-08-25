package binary

import (
	"log/slog"
	"slices"
)

/*
Pattern A pattern is a representation of the binary stream representing the client that we will search for.
We use an int16 type so that we can represent -1 as a wildcard search, and a valid pattern should only ever
encompass the values [-1, 255].
*/
type Pattern []int16

func (p *Pattern) Empty() []byte {
	return make([]byte, len(*p))
}

func StringToPattern(s string) Pattern {
	p := make(Pattern, len(s))
	for i := 0; i < len(s); i++ {
		p[i] = int16(s[i])
	}
	return p
}

func Patch(in *[]byte, find Pattern, replace []byte) {
	for i := 0; i < len(*in)-len(find); i++ {
		cmp := (*in)[i : i+len(find)]
		if slices.EqualFunc(cmp, find, func(b byte, i int16) bool {
			if i == -1 {
				return true
			}

			return int16(b) == i
		}) {

			slog.Debug("found pattern", "offset", i)
			for j := i; j < i+len(replace); j++ {
				(*in)[j] = replace[j-i]
			}
		}
	}
}
