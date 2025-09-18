package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

// Map
// ● Go-Lang memiliki sebuah struct beranama sync.Map
// ● Map ini mirip Go-Lang map, namun yang membedakan, Map ini aman untuk menggunaan concurrent menggunakan goroutine
// ● Ada beberapa function yang bisa kita gunakan di Map :
// ○ Store(key, value) untuk menyimpan data ke Map
// ○ Load(key) untuk mengambil data dari Map menggunakan key
// ○ Delete(key) untuk menghapus data di Map menggunakan key
// ○ Range(function(key, value)) digunakan untuk melakukan iterasi seluruh data di Map

// AddToMap adalah fungsi yang akan dijalankan oleh setiap goroutine.
// Parameternya:
// - data  → pointer ke sync.Map, struktur map thread-safe untuk concurrent access
// - value → integer yang akan dimasukkan ke dalam map
// - group → pointer ke WaitGroup, digunakan untuk sinkronisasi goroutine
func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	// defer memastikan Done() selalu dipanggil di akhir fungsi,
	// sehingga WaitGroup counter akan berkurang meskipun ada error/panic
	defer group.Done()

	// Menyimpan pasangan (key, value) ke dalam sync.Map.
	// sync.Map aman digunakan oleh banyak goroutine secara bersamaan.
	data.Store(value, value)
}

// TestMap adalah fungsi unit test yang akan dijalankan oleh `go test`
// Nama fungsi harus diawali dengan "Test" agar dikenali testing framework.
func TestMap(t *testing.T) {
	// Membuat instance baru sync.Map (map aman untuk concurrency)
	data := &sync.Map{}

	// Membuat WaitGroup untuk menunggu semua goroutine selesai
	group := &sync.WaitGroup{}

	// Meluncurkan 100 goroutine untuk menyimpan data secara paralel
	for i := 0; i < 100; i++ {
		// Tambah counter WaitGroup sebelum memulai goroutine
		group.Add(1)

		// Menjalankan goroutine yang memanggil AddToMap
		// Setiap goroutine akan menyimpan angka i ke dalam map
		go AddToMap(data, i, group)
	}

	// Menunggu sampai semua goroutine memanggil Done() → counter kembali ke 0
	group.Wait()

	// Mengiterasi seluruh isi map secara thread-safe
	// Range menerima callback yang dipanggil untuk setiap entry
	data.Range(func(key, value interface{}) bool {
		// Menampilkan key dan value dari setiap data yang tersimpan
		fmt.Println(key, ":", value)

		// return true berarti iterasi lanjut ke entry berikutnya
		// jika return false → berhenti iterasi
		return true
	})
}
