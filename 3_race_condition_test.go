package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

// race condition
// Masalah Dengan Goroutine
// ● Saat kita menggunakan goroutine, dia tidak hanya berjalan secara concurrent, tapi bisa parallel juga, karena bisa ada beberapa thread yang berjalan secara parallel untuk menjalankan goroutine
// ● Hal ini sangat berbahaya ketika kita melakukan manipulasi data variable yang sama oleh beberapa goroutine secara bersamaan
// ● Hal ini bisa menyebabkan masalah yang namanya Race Condition, jadi kalo ada ngelakuin perubahan data di beberapa goroutine terhadap sati variable yg sama maka akan kejadian race condition
// ● ini disemua Bahasa pemrograman yg memiliki parallel atau Concurrency udh pasti problem race condition ini ada
func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()

	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}

// tiap function TestRaceCondition di run hasilnnya akan selalu berubah dan tidak akan sampai 100.000 ubah karena punya variable yg di sharing ke beberapa goroutine bahkan 1000 goroutine seperti yg tertulis di for dan mengakses variable x yg sama, lalu kita manipulasi variable tersebut.
// jadi bisa saja ada satu goroutine mengakses data atau value yg punya value tetap sama juga contoh:
// 	x ada di posisi 500, ternyata ada 2 goroutine yg mengakses secara berbarengan, karena kan paralel ya di beberapa processor, artinya ada dua mungkin juga lebih goroutine yg akases nya seperti ini
// x = 500
// x = 500
// jadi ada kondisi waktu dan presisinya sangat tepat, dmn ketika ngelakuin conter itu sama nilai nya seperti 2 contoh di atas, jadi otomatis nilai nya akan hilang sebagian dan tidak akan pernah mencapai 100.0000.
// jadi ini bahaya nya menggunakan concurency jadi ada problem yg namanya race condition (balapan antar goroutine untuk mengubah variabel yg sama)
