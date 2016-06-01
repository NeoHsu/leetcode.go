package isomorphicStrings

func isIsomorphic(s string, t string) (isIsomorphic bool) {
	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)

	isIsomorphic = true
	for i := 0; i < len(s); i++ {
		sValue, sOk := sMap[s[i]]
		tValue, tOk := tMap[t[i]]
		if sOk || tOk {
			if sValue != t[i] || tValue != s[i] {
				isIsomorphic = false
				return
			}
		} else {
			sMap[s[i]], tMap[t[i]] = t[i], s[i]
		}
	}
	return
}
