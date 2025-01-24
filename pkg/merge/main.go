package merge

func MergeArray1(dest []string, src []string) (result []string) {
	result = make([]string, len(dest)+len(src))
	copy(result, dest)
	copy(result[len(dest):], src)
	return
}
