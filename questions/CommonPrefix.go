package questions

func LongestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	for i, c := range strs[0] {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != byte(c) {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
