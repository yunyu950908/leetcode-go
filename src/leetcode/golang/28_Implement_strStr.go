package golang

//func strStr(haystack string, needle string) int {
//	return strings.Index(haystack, needle)
//}

func strStr(haystack string, needle string) int {
	var ih, in, lh, ln int
	lh, ln = len(haystack), len(needle)
	for ih < lh && in < ln {
		if haystack[ih] == needle[in] {
			ih++
			in++
		} else {
			ih -= in - 1
			in = 0
		}
	}
	if in == ln {
		return ih - ln
	} else {
		return -1
	}
}

// TODO: KMP implementation
