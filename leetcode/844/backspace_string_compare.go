package main

func backspaceCompare(s string, t string) bool {
	ss, sb, sc := [200]uint8{}, 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '#' {
			sb++
		} else if sb > 0 {
			sb--
		} else {
			ss[sc] = s[i]
			sc++
		}
	}
	ts, tb, tc := [200]uint8{}, 0, 0
	for i := len(t) - 1; i >= 0; i-- {
		if t[i] == '#' {
			tb++
		} else if tb > 0 {
			tb--
		} else {
			ts[tc] = t[i]
			tc++
		}
	}

	return ss == ts
}
