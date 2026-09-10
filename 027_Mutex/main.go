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
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()
	// Here modification happen so, we use mutex lock here for views value
	p.mu.Lock()
	p.views += 1
	// after this p.iews += 1 the operation complete. So, now views resource should be unlock
	// p.mu.Unlock()

		// we can write Unlock() below this resource but what happen if that resource get any kind of error, then this task never complete and our VIEWS resource STUCK in LOCK for lifetime.
