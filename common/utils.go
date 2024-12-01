package common

func Count[T any](slice []T, cmp func(T) bool) int {
	count := 0
	for _, v := range slice {
		if cmp(v) {
			count++
		}
	}
	return count
}
