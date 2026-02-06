package partitions

import "strconv"

// makeByteKeys makes n keys like "k0","k1"... as []byte.
// Allocations happen once, outside the benchmark timer.
func makeByteKeys(n int) [][]byte {
	keys := make([][]byte, n)
	for i := 0; i < n; i++ {
		// Build as bytes without fmt. This still allocates once per key (string + []byte),
		// but that's outside timing which is what we want.
		s := "k" + strconv.Itoa(i)
		keys[i] = []byte(s)
	}
	return keys
}

func makeHotByteKeys(n int) [][]byte {
	// n should be a power of two (256/1024/4096) so i&(n-1) is cheap.
	return makeByteKeys(n)
}
