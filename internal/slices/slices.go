package slices

func Contains[E comparable](haystack []E, needle E) bool {
	for _, e := range haystack {
		if e == needle {
			return true
		}
	}
	return false
}
