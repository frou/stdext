package stdext

import (
	"strconv"
)

// ParseInt parses the ASCII representation of a base-10 integer. Assumes the size of type int on this platform is 64 bits.
func ParseInt(s string) (int, error) {
	n, err := strconv.ParseInt(s, 10, 64)
	nplatform := int(n)
	// if size := unsafe.Sizeof(n); size != 8 {
	// 	panic(size)
	// }
	return nplatform, err
}
