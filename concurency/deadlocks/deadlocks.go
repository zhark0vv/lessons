package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

/*
Здесь мы создаем заведомо заблокированные горутины и отслеживаем их через pprof:
http://localhost:6060/debug/pprof/goroutine?debug=1

1 @ 0x10014d628 0x100162bb8 0x100162b95 0x100183358 0x1001ac528 0x1001ac310 0x10045518c 0x100187954
#	0x100183357	sync.runtime_SemacquireMutex+0x27	/opt/homebrew/opt/go/libexec/src/runtime/sema.go:77
#	0x1001ac527	sync.(*Mutex).lockSlow+0x1e7		/opt/homebrew/opt/go/libexec/src/sync/mutex.go:171
#	0x1001ac30f	sync.(*Mutex).Lock+0x5f			/opt/homebrew/opt/go/libexec/src/sync/mutex.go:90
#	0x10045518b	main.main.func2+0xdb			/Users/vladislavzharkov/GolandProjects/lessons/concurency/deadlocks.go:24

Мы видим, что горутины заведомо "держат" друг друга.
*/
func main() {
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	var wg sync.WaitGroup
	var mu1, mu2 sync.Mutex

	wg.Add(2)

	go func() {
		defer wg.Done()
		mu1.Lock()
		fmt.Println("Goroutine 1 locked mu1")
		mu2.Lock()
		fmt.Println("Goroutine 1 locked mu2")
		mu2.Unlock()
		mu1.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu2.Lock()
		fmt.Println("Goroutine 2 locked mu2")
		mu1.Lock()
		fmt.Println("Goroutine 2 locked mu1")
		mu1.Unlock()
		mu2.Unlock()
	}()

	wg.Wait()
	fmt.Println("Main completed")
}
