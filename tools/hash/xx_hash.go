package hash

import "github.com/cespare/xxhash"

func XXHashSum64(d []byte) uint64 {
	return xxhash.Sum64(d)
}
func XXHashSum64String(s string) uint64 {
	return xxhash.Sum64String(s)
}
