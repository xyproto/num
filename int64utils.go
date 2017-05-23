package gofractions

// Return the smallest number of two given integers
func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// Return the absolute value of an integer
func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}
