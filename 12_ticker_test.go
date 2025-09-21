package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

// time.Ticker
// ● Ticker adalah representasi kejadian yang berulang berbeda dengan timer, kalo timer sekali kejadian
// ● Ketika waktu ticker sudah expire, maka event akan dikirim ke dalam channel
// ● Untuk membuat ticker, kita bisa menggunakan time.NewTicker(duration)
// ● Untuk menghentikan ticker, kita bisa menggunakan Ticker.Stop()
// ● Untuk mendapatkan data ticker, kita bisa menggunakan ticker.C
// ● Untuk mereset ticker, kita bisa menggunakan ticker.Reset()

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
		fmt.Println("Ticker stopped")
	}()

	for time := range ticker.C {
		fmt.Println("Ticker event at", time)
	}

}

// time.Tick()
// ● Kadang kita tidak butuh data Ticker nya, kita hanya butuh channel nya saja
// ● Jika demikian, kita bisa menggunakan function timer.Tick(duration), function ini tidak akan mengembalikan Ticker, hanya mengembalikan channel timer nya saja

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for tick := range channel {
		fmt.Println("Tick at", tick)
	}

}
