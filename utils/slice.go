package utils

func InSlice(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func RevertSlice(slice []any) []any {
	copySlice := make([]any, 0)
	for idx := len(slice) - 1; idx >= 0; idx -= 1 {
		copySlice = append(copySlice, slice[idx])
	}

	return copySlice
}
