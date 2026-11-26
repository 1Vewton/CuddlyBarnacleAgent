package operations

// IsIncluded checks if certain element is in a slice
func IsIncluded[T any](
	dataSlice []T,
	element T,
	isEqualFunc func(T, T) bool,
) bool {
	for _, data := range dataSlice {
		if isEqualFunc(data, element) {
			return true
		}
	}
	return false
}
