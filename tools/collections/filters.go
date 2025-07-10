package collections

// Filter a slice and return a new slice with  wanted return value
func Filter[T any](ori []T, f func(arg T) bool) []T {
	if ori == nil {
		return nil
	}
	if len(ori) == 0 {
		return []T{}
	}
	filtered := make([]T, 0, len(ori))
	for _, v := range ori {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

// FilterMap a map and return a new map with  wanted return value
func FilterMap[K comparable, V any](ori map[K]V, f func(k K, v V) bool) map[K]V {
	if ori == nil {
		return nil
	}
	if len(ori) == 0 {
		return map[K]V{}
	}
	filtered := make(map[K]V, len(ori))
	for k, v := range ori {
		if f(k, v) {
			filtered[k] = v
		}
	}
	return filtered
}
