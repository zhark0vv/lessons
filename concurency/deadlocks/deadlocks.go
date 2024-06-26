package main

import (
	"fmt"
	"log"
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
	var (
		mu      sync.Mutex
		wg      sync.WaitGroup
		counter int
	)

	wg.Add(3)
	go func() {
		defer wg.Done()
		log.Fatal(http.ListenAndServe("localhost:6060", nil))
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		counter++
		fmt.Println("First goroutine finished, counter:", counter)
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		counter++
		fmt.Println("Second goroutine finished, counter:", counter)
	}()

	wg.Wait()
}
