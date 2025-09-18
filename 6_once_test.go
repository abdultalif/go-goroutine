package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

// Once
// ● Once adalah fitur di Go-Lang yang bisa kita gunakan untuk memastikan bahwa sebuah function di eksekusi hanya sekali
// ● Jadi berapa banyak pun goroutine yang mengakses, bisa dipastikan bahwa goroutine yang pertama yang bisa mengeksekusi function nya.
// ● Goroutine yang lain akan di hiraukan, artinya function tidak akan dieksekusi lagi

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			// jika menggunakna once.Do maka goroutine yg masuk hanya 1
			once.Do(OnlyOnce)
			// OnlyOnce()
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter:", counter)
}
