package questions

import "strings"

/*
Input: pattern = "abba", str = "dog cat cat dog"
Output: true
*/

func wordPattern(pattern string, str string) bool {
	stringArr := strings.Split(str, " ")
	patternArr := strings.Split(pattern, "")
	patternMap := make(map[string]string)
	stringMap:= make(map[string]string)
	if len(patternArr) != len(stringArr) {
		return false
	}
	for i := 0; i < len(patternArr); i++ {
		if val, ok := patternMap[patternArr[i]]; ok {
			if val != stringArr[i] {
				return false
			}
		} else {
			patternMap[patternArr[i]] = stringArr[i]
		}
		if val,ok:= stringMap[stringArr[i]];ok {
			if val != patternArr[i] {
				return false
			}
		} else {
			stringMap[stringArr[i]]= patternArr[i]
		}
	}
	return true
}
