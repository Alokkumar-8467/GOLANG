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

		// So we move the p.mu.Unlock() to defer() function.
	//  As we know that either the function run or give error or any deadlock condition the defer function run at END of that function.

		// Never Lock the Whole function or logic by MUTEX, it's a bad practice
	// Only lock that particular line that perform modification.
		// Sometime MUTEX careate a bottleNeck situation.

}

func main() {

	var wg sync.WaitGroup
		myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}


