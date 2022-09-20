package timeseriescache

type CircularCache struct {
	cache    []CacheElem
	startIdx int
}
