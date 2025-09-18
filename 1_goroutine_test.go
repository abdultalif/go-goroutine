package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

// Pengenalan Parallel Programming
// ● Saat ini kita hidup di era multicore, dimana jarang sekali kita menggunakan prosesor yang single core
// ● Semakin canggih perangkat keras, maka software pun akan mengikuti, dimana sekarang kita bisa dengan mudah membuat proses parallel di aplikasi.
// ● Parallel programming sederhananya adalah memecahkan suatu masalah dengan cara membaginya menjadi yang lebih kecil, dan dijalankan secara bersamaan pada waktu yang bersamaan pula

// Contoh Parallel
// ● Menjalankan beberapa aplikasi sekaligus di sistem operasi kita (office, text editor, browser, dan lain-lain)
// ● Beberapa koki menyiapkan makanan di restoran, dimana tiap koki membuat makanan masing-masing
// ● Antrian di Bank, dimana tiap teller melayani nasabah nya masing-masing

// Process vs Thread
// Process 													Thread
// Process adalah sebuah eksekusi program 					Thread adalah segmen dari process
// Process mengkonsumsi memory besar 						Thread menggunakan memory kecil
// Process saling terisolasi dengan process lain			Thread bisa saling berhubungan jika dalam process yang sama
// Process lama untuk dijalankan dihentikan 				Thread cepat untuk dijalankan dan dihentikan

// Parallel vs Concurrency (penjelasan sebelum masuk goroutine)
// ● Berbeda dengan paralel (menjalankan beberapa pekerjaan secara bersamaan), concurrency adalah menjalankan beberapa pekerjaan secara bergantian
// ● Dalam parallel kita biasanya membutuhkan banyak Thread, sedangkan dalam concurrency, kita hanya membutuhkan sedikit Thread

// nah golang menggunakna concurrency

// Contoh Concurrency
// ● Saat kita makan di cafe, kita bisa makan, lalu ngobrol, lalu minum, makan lagi, ngobrol lagi, minum lagi, dan seterusnya. Tetapi kita tidak bisa pada saat yang bersamaan minum, makan dan ngobrol, hanya bisa melakukan satu hal pada satu waktu, namun bisa berganti kapanpun kita mau.

// Goroutine adalah sebuah thread ringan yang dikelola oleh Go Runtime. Ukuran Goroutine sangat kecil, sekitar 2kb, jauh lebih kecil dibandingkan Thread yang bisa sampai 1mb atau 1000kb. Namun tidak seperti thread yang berjalan parallel, goroutine berjalan secara concurrent.

// Cara Kerja Goroutine
// ● Sebenarnya, Goroutine dijalankan oleh Go Scheduler dalam thread, dimana jumlah thread nya sebanyak GOMAXPROCS (biasanya sejumlah core CPU)
// ● Jadi sebenarnya tidak bisa dibilang Goroutine itu pengganti Thread, karena Goroutine sendiri berjalan di atas Thread
// ● Namun yang mempermudah kita adalah, kita tidak perlu melakukan manajemen Thread secara manual, semua sudah diatur oleh Go Scheduler

// Dalam Go-Scheduler, kita akan mengenal beberapa terminologi
// ● G : Goroutine
// ● M : Thread (Machine)
// ● P : Processor

// Membuat goroutine
// Untuk membuat goroutine di Golang sangatlah sederhana, Kita hanya cukup menambahkan perintah go sebelum memanggil function yang akan kita jalankan dalam goroutine. Saat sebuah function kita jalankan dalam goroutine, function tersebut akan berjalan secara asynchronous, artinya tidak akan ditunggu sampai function tersebut selesai. Aplikasi akan lanjut berjalan ke kode program selanjutnya tanpa menunggu goroutine yang kita buat selesai.

// Seperti yang sebelumnya dijelaskan, bahwa goroutine itu sangat ringan (2kb). Kita bisa membuat ribuan, bahkan sampai jutaan goroutine tanpa takut boros memory(RAM). Tidak seperti thread yang ukurannya berat, goroutine sangatlah ringan

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	// jika pemanggilan function nya di tambahkan go maka funnction nya dirunning menggunakan goroutine dan function nya berjalan secara asycronous
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

// Tets untuk membuat banyak goroutine dan karena hasilnya berjalan secara asyncronous hasilnya akan muncul sesuai logika kita tapi tidak berurutan
func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(10 * time.Second)
}
