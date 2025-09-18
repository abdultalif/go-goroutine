package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Pool
// ● Pool adalah implementasi design pattern bernama object pool pattern. jadi kalo sebelumnya sudah pernah belajar tentang design pattern, jadi di golang untuk implementasi bernama object pool pattern. ini digunakan untuk concurency dan paralel programming
// ● Sederhananya, design pattern Pool ini digunakan untuk menyimpan data, selanjutnya untuk
// menggunakan datanya, kita bisa mengambil dari Pool, dan setelah selesai menggunakan datanya,
// kita bisa menyimpan kembali ke Pool nya
// ● Implementasi Pool di Go-Lang ini sudah aman dari problem race condition

func TestPool(t *testing.T) {
	pool := sync.Pool{}

	pool.Put("Abdul")
	pool.Put("Talif")
	pool.Put("test")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}
	time.Sleep(3 * time.Second)
	fmt.Println("Selessai")
}
