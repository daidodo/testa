package assert

import "sync"

const kHOOK = true

type common struct {
	su     sync.RWMutex
	output []byte // Output generated by test or benchmark.
}
