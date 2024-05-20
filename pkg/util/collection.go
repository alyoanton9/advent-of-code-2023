package util

func SliceToSet[T comparable](slice []T) map[T]struct{} {
	set := make(map[T]struct{})

	for _, elem := range slice {
		set[elem] = struct{}{}
	}

	return set
}

// TODO replace with lo.CountValues
func Count[T comparable](elems []T) map[T]int {
	counter := make(map[T]int)
	for _, elem := range elems {
		counter[elem]++
	}

	return counter
}

func Keys[K comparable, V any](map_ map[K]V) []K {
	keys := make([]K, 0)
	for k, _ := range map_ {
		keys = append(keys, k)
	}

	return keys
}
