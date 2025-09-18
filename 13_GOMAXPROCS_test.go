package golang_goroutines

import (
	"fmt"
	"runtime"
	"testing"
)

// GOMAXPROCS
// ● Sebelumnya diawal kita sudah bahas bahwa goroutine itu sebenarnya dijalankan di dalam Thread. jadi tidak ada perbandingan goroutine dengan thread karena goroutine sendiri dijalankan didalam thread
// ● Pertanyaannya, seberapa banyak Thread yang ada di Go-Lang ketika aplikasi kita berjalan?
// ● Untuk mengetahui berapa jumlah Thread, kita bisa menggunakan GOMAXPROCS, yaitu sebuah function di package runtime yang bisa kita gunakan untuk mengubah jumlah thread atau mengambil jumlah thread yang ada saat ini
// ● Secara default, jumlah thread di Go-Lang itu sebanyak jumlah CPU di komputer kita.
// ● Kita juga bisa melihat berapa jumlah CPU kita dengan menggunakan function runtime.NumCpu()

func TestGetGomaxprocs(t *testing.T) {

	// contoh untuk menambahkan goroutine di contohh ini saya coba tambah 100
	// group := sync.WaitGroup{}
	// for i := 0; i < 100; i++ {
	// 	group.Add(1)
	// 	go func() {
	// 		time.Sleep(3 * time.Second)
	// 		group.Done()
	// 	}()

	// }

	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU: ", totalCPU)

	// Total thread
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread: ", totalThread)

	totalgoroutine := runtime.NumGoroutine()
	fmt.Println("Total goroutine: ", totalgoroutine)

	// group.Wait()
}

func TestChangeThreadNumber(t *testing.T) {

	// ubah thread
	runtime.GOMAXPROCS(5)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread: ", totalThread)
}

// Peringatan
// ● Menambah jumlah thread tidak berarti membuat aplikasi kita menjadi lebih cepat
// ● Karena pada saat yang sama, 1 CPU hanya akan menjalankan 1 goroutine dengan 1 thread
// ● Oleh karena ini, jika ingin menambah throughput aplikasi, disarankan lakukan vertical scaling (dengan menambah jumlah CPU) atau horizontal scaling (menambah node baru)
