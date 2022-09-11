package main

func isIsomorphic(s string, t string) bool {
	sm := make(map[uint8]int)
	tm := make(map[uint8]int)

	for i := 0; i < len(s); i++ {
		if _, ok := sm[s[i]]; !ok {
			sm[s[i]] += i
		}
		if _, ok := tm[t[i]]; !ok {
			tm[t[i]] += i
		}
	}

	for i := 0; i < len(s); i++ {
		if sv, sOk := sm[s[i]]; sOk {
			if tv, tOk := tm[t[i]]; tOk {
				if sv != tv {
					return false
				}
			}
		}
	}

	return true
}
