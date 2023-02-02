package c_lib

import "sync"

type Container struct {
	mu           sync.Mutex
	Exec_counter int
}
