package main

var m = map[uint8][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	if digits == "" {
		return res
	}

	letters := m[digits[0]]
	lastCombinations := letterCombinations(digits[1:])
	for i := 0; i < len(letters); i++ {
		if len(lastCombinations) == 0 {
			res = append(res, letters[i])
		} else {
			for j := 0; j < len(lastCombinations); j++ {
				res = append(res, letters[i]+lastCombinations[j])
			}
		}
	}

	return res
}
