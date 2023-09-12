package utils

// todo: make generic
func Contains(haystack []string, needle string) bool {
	for _, e := range haystack {
		if e == needle {
			return true
		}
	}
	return false
}
