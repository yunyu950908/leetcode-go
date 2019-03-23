package golang

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	result := ""
	first := strs[0]
	for i := range first {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != first[i] {
				return result
			}
		}
		result += string(first[i])
	}
	return result
}
