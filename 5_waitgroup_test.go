package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// WaitGroup
// ● WaitGroup adalah fitur yang bisa digunakan untuk menunggu sebuah proses selesai dilakukan. biasanya si WaitGroup di gunakan untuk menunggu proses goroutine selesai, karena kan proses goroutine berjalan secara asycronous, sebelumnya kita nungguin nya menggunakan (time.Sleep) ini tidak direkomendasikan sebenarnya. karena kita kan ga tau goroutine itu selesai nya berapa lama. jadi ideal nya kalo mau menunggu itu menggunakan WaitGroup
// ● Hal ini kadang diperlukan, misal kita ingin menjalankan beberapa proses menggunakan goroutine, tapi kita ingin semua proses selesai terlebih dahulu sebelum aplikasi kita selesai. WaitGroup juga bisa digunakan lebih dari satu goroutine, misal ada 100 goroutine pengen di tunggu seratus seratus nya selesai baru program nya berhenti atau function selesai WaitGroup adalah yg cocok untuk digunakan
// ● Kasus seperti ini bisa menggunakan WaitGroup
// ● Untuk menandai bahwa ada proses goroutine, kita bisa menggunakan method Add(int), setelah proses goroutine selesai, kita bisa gunakan method Done(). done sebenarnya untuk menurunkan
// ● Untuk menunggu semua proses selesai, kita bisa menggunakan method Wait(). wait akan menunggu sampai jumlah si done sesuai apa yg ada di add

func RunAsynchronous(group *sync.WaitGroup) {
	// group.Done digunakan untuk mengurangi group.Add
	defer group.Done()

	// ini berfungsi untuk running 1 proses asyncronous
	group.Add(1)

	fmt.Println("Hello")
	// diam 1 detik untuk menunggu proses goroutine
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	// buat waitgroup nya
	group := &sync.WaitGroup{}

	// membuat perulangan sampai 100
	for i := 0; i < 100; i++ {
		// membuat goroutine sebanyak 10 untuk me running func RunAsynchronous
		go RunAsynchronous(group)
	}

	// untuk menunggu semua selesai dan group.Add nya 0 lagi karena terus di kurangi group.Done maka menggunakan group.Wait()
	group.Wait()
	fmt.Println("Selesai")
}
