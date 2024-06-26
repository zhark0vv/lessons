/*
Компоненты GMP
Goroutines (G)
Горутина — это легковесный поток выполнения, управляемый рантаймом Go.
Горутина намного легче, чем системный поток (OS thread), и может быть создана в огромных количествах (но лучше не злоупотреблять).
Каждая горутина имеет стек, который динамически изменяется в зависимости от потребностей.


Machine Threads (M)
Машинные потоки (M) — это системные потоки (OS threads), на которых выполняются горутины.
Могут быть непосредственно сопоставлены с потоками операционной системы.
Рантайм Go может динамически увеличивать или уменьшать количество потоков M в зависимости от нагрузки.

Processors (P)
Процессоры (P) управляют выполнением горутин, обеспечивая контекст для выполнения (например, готовый набор задач, очередь задач).
P управляют тем, какие горутины могут выполняться на потоках M.
Количество P определяется значением переменной GOMAXPROCS, которое по умолчанию равно количеству процессорных ядер в системе.
P также содержит локальный пул готовых горутин.

Как это работает вместе
Модель GMP работает следующим образом:

Создание горутин: Когда создается горутина, она помещается в очередь задач (run queue) одного из процессоров P.
Назначение процессора: Процессор P берет горутину из своей очереди задач и назначает ее на выполнение потоку M.
Выполнение горутины: Поток M выполняет горутину. Если горутина блокируется (например, ожидает завершения операции ввода-вывода), поток M может переключиться на выполнение другой горутины.
Системные вызовы: Если горутина делает системный вызов, который блокирует поток M, рантайм Go может создать новый поток M, чтобы продолжить выполнение других горутин.
Преемственность: При необходимости, P может "украсть" задачи из очередей других P, чтобы сбалансировать нагрузку (work stealing).
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// Из интересного: в Go кооперативно-вытесняющая многозадачность, что означает, что горутины сами передают управление друг другу.
	// До версии 1.14 была чисто кооперативная многозадачность, но с 1.14 добавились системные вызовы, которые позволяют Go выполнять системные вызовы в отдельных потоках.
	// Если запустить этот код с GOMAXPROCS(1), то горутины будут выполняться последовательно, так как один процессор не может выполнять несколько горутин одновременно.
	// На версии до 1.14 горутина 2 не запустилась бы никогда.
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			fmt.Println("Goroutine 1")
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Println("Goroutine 2:", i)
			time.Sleep(time.Millisecond * 100)
		}
	}()

	wg.Wait()
	fmt.Println("All goroutines finished")
}
