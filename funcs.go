package go_funcs

import "sync"

type fns struct {
	data      []string
	satisfies []bool
	once      sync.Once
}
