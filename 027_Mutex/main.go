package main

import (
	"fmt"
	"sync"
)

// When we do multiThreading then to avoide the raise condition we use MUTEX

/* What is raise condition
When multiple processes use the same or single resources then modification on that resources is bot atomic.
There is a hight chancet that if Process 1 chance the resource then process 2 also change the resource so it create conflict in the final result.
*/

type post struct {
	views int
	// Here we add mutex in views resource.
	mu sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
