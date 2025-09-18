package golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// Channel adalah tempat komunikasi secara synchronous yang bisa dilakukan oleh goroutine. Di Channel terdapat pengirim dan penerima, biasanya pengirim dan penerima adalah goroutine yang berbeda. Saat melakukan pengiriman data ke Channel, goroutine akan ter-block, sampai ada yang menerima data tersebut. Maka dari itu, channel disebut sebagai alat komunikasi synchronous (blocking). Channel cocok sekali sebagai alternatif seperti mekanisme async await yang terdapat di beberapa
// bahasa pemrograman lain

// karakteristik channel
// Secara default channel hanya bisa menampung satu data, jika kita ingin menambahkan data lagi, harus menunggu data yang ada di channel diambil. Channel hanya bisa menerima satu jenis data. Channel bisa diambil dari lebih dari satu goroutine. Channel harus di close jika tidak digunakan, atau bisa menyebabkan memory leak (channel nya akan menggantung terus di memory jika tidak di close, direkomendasikan close channel)

// Membuat Channel
// ● Channel di Go-Lang direpresentasikan dengan tipe data chan
// ● Untuk membuat channel sangat mudah, kita bisa menggunakan make(), mirip ketika membuat map
// ● Namun saat pembuatan channel, kita harus tentukan tipe data apa yang bisa dimasukkan kedalam channel tersebut

// func TestCreateChannel1(t *testing.T) {
// 	channel := make(chan string)

// 	// ini cara untuk mengirim data ke channel
// 	channel <- "Talif"

// 	// sedangkan untuk menerima/ambil data seperti ini caranya di masukan ke variable
// 	data := <- channel

// 	// biasakan kalo sudah buat channel buat juga close nya di paling bawah karena itu best practice nya
// 	// close(channel)
// 	defer close(channel)
// }

func TestCreateChannel2(t *testing.T) {
	// Membuat channel dengan tipe data string, tapi jika sudah real case biasanya gunakan struct untuk tipe data channel nya
	channel := make(chan string)

	// Menutup channel
	defer close(channel)

	// jika function di awali dengan go berarti dia goroutine dan berjalan secara asyncronous
	// oh iya ini cara memanggil goroutine dengan anonimous function di bawah nanti ada cara menggunakan goroutine dengan function sebagai parameter
	go func() {
		time.Sleep(2 * time.Second)
		// jika mengirim data ke channel seperti ini pastikan data nya ada yg menerima kalo tidak error panic
		channel <- "Abdul Talif"
		fmt.Println("Selesai Mengirim Data ke channel")
	}()

	// jika mengambil data dari channel seperti ini pastikan ada data di dalam channel jika tidak ada error (deadlock) juga jadi mengirim dan menerima wajib ada bersamaan
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// berbedan dengan parameter biasa, parameter channel ini dari awal sudah pass by referance, kalo parameter lain kan pass by value default nya jadi harus menggunakan pointer (*) jika ingin di ubah jadi pass by referance. jika ingin tahu pass by referance dan pass by value apa buka catatan golang dasar
func GiveMeResponse(channel chan string) {
	// bengong atau diam selama 2 detik
	time.Sleep(2 * time.Second)

	// mengirim channel dengan string sesuai tipe data yg ketika membuat channel di function TestChannelAsParameter
	channel <- "Abdul Talif"
}

// Channel Sebagai Parameter
// ● Dalam kenyataan pembuatan aplikasi, seringnya kita akan mengirim channel ke function lain via parameter
// ● Sebelumnya kita tahu bahkan di Go-Lang by default, parameter adalah pass by value, artinya value akan diduplikasi lalu dikirim ke function parameter, sehingga jika kita ingin mengirim data asli, kitabiasa gunakan pointer (agar pass by reference).
// ● Berbeda dengan Channel, kita tidak perlu melakukan hal tersebut
func TestChannelAsParameter(t *testing.T) {
	// membuat channel dengan tipe data string
	channel := make(chan string)

	// memanggil function dengan parameter di goroutine
	go GiveMeResponse(channel)

	// menerima data dari channel, jika data tidak dati channel maka akan error deadlock
	data := <-channel
	fmt.Println(data)
	close(channel)
}

// Channel In dan Out
// ● Saat kita mengirim channel sebagai parameter, isi function tersebut bisa mengirim dan menerima data dari channel tersebut
// ● Kadang kita ingin memberi tahu terhadap function, misal bahwa channel tersebut hanya digunakan untuk mengirim data, atau hanya dapat digunakan untuk menerima data
// ● Hal ini bisa kita lakukan di parameter dengan cara menandai apakah channel ini digunakan untuk in(mengirim data) atau out (menerima data)

// jika parameter nya seperti ini channel chan <- string berarti menandakan parameter hanya dibolehkan untun mengirim channel
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Abdul Talif"
}

// jika parameter nya seperti ini channel <- chan string berarti parameter nya hanya diperbolehkan menerima data dari channel
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)

}

// Buffered Channel
// ● Seperti yang dijelaskan sebelumnya, bahwa secara default channel itu hanya bisa menerima 1 data
// ● Artinya jika kita menambah data ke-2, maka kita akan diminta menunggu sampai data ke-1 ada yang mengambil
// ● Kadang-kadang ada kasus dimana pengirim lebih cepat dibanding penerima, dalam hal ini jika kita menggunakan channel, maka otomatis pengirim akan ikut lambat juga
// ● Untuknya ada Buffered Channel, yaitu buffer yang bisa digunakan untuk menampung data antrian di Channel

// Buffer Capacity
// ● Kita bebas memasukkan berapa jumlah kapasitas antrian di dalam buffer
// ● Jika kita set misal 5, artinya kita bisa menerima 5 data di buffer.
// ● Jika kita mengirim data ke 6, maka kita diminta untuk menunggu sampai buffer ada yang kosong
// ● Ini cocok sekali ketika memang goroutine yang menerima data lebih lambat dari yang mengirim data

// contoh:
// ch := make(chan int, 3) //channel dengan buffer kapasitas 3
// angka 3 menunjukan jumlah slot yg bisa menampung data di dalam channel
// jadi goroutine bisa mengirim data tanpa harus menunggu goroutine lain menerima, selama kapasitas belum penuh
// jika buffer penuh, pengirim akan block sampau ada data yang di ambil
// jika buffer kosong. receive(penerima) akan block sampai ada data yg masuk
func TestBufferedChannel(t *testing.T) {
	// membuat channel dengan buffer 2
	channel := make(chan string, 2)
	defer close(channel)

	go func() {
		channel <- "Abdul"
		channel <- "Talif"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("selesai")

}

// Range Channel
// ● Kadang-kadang ada kasus sebuah channel dikirim data secara terus menerus oleh pengirim
// ● Dan kadang tidak jelas kapan channel tersebut akan berhenti menerima data
// ● Salah satu yang bisa kita lakukan adalah dengan menggunakan perulangan range ketika menerima data dari channel
// ● Ketika sebuah channel di close(), maka secara otomatis perulangan tersebut akan berhenti
// ● Ini lebih sederhana dari pada kita melakukan pengecekan channel secara manual
func TestRangeChannel(t *testing.T) {
	// membuat channel dengan tipe string
	channel := make(chan string)

	// goroutine dengan anonimous function
	go func() {
		// Masukan data channel dengan loop dengan convert int jadi string dengan itoa
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		// close channel pastikan ini dilakukan jika tidak perulangan for range dibawah tidak akan berhenti
		close(channel)
	}()

	// Mengambil data dari channel dengan range
	for data := range channel {
		// print data
		fmt.Println(data)
	}

	fmt.Println("Selesai")
}

// Select Channel
// ● Kadang ada kasus dimana kita membuat beberapa channel, dan menjalankan beberapa goroutine untuk tiap channel
// ● Lalu kita ingin mendapatkan data dari semua channel tersebut, nah kalo kita menggunakan perulangan for range (range channel) itu kan agak sedikit ribet ya, karen dia  hanya bisa buat satu channel, kalo missal channel nya lebih dari 1 itu aga menyulitkan jika menggunakan for range karena dia hanyak ngeblock di satu channel saja
// ● Untuk melakukan hal tersebut, kita bisa menggunakan select channel di Go-Lang
// ● Dengan select channel, kita bisa memilih data tercepat dari beberapa channel, jika data dating secara bersamaan di beberapa channel. jadi jika nanti kalo select channel sebutin channel nya ada apa aja, nanti tiap ada data yg masuk ke salah satu channel nya maka otomatis akan di ambil dan dieksekusi perintah kode kita, kalo missal ada 2 data yg datang secara bersamaan nanti akan di ambil dulu secara random salah satu nya, baru nanti select selanjutnya akan di ambil dari channel lainya, jadi sekali kita melakukan select dia akan sekali juga ngambil data dari channel
// ● kalo pengen ngambil semua data di dalam beberapa channel tersebut kita harus menggunakan perulangan, nanti didalam perulangan nya menggunakan select

// Default Select
// ● Apa yang terjadi jika kita melakukan select terhadap channel yang ternyata tidak ada datanya? maka by default akan terjadi deadlock
// ● Maka kita akan menunggu sampai data ada
// ● Kadang mungkin kita ingin melakukan sesuatu jika misal semua channel tidak ada datanya Ketika kita melakukan select channel
// ● Dalam select, kita bisa menambahkan default, dimana ini akan dieksekusi jika memang di semua channel yang kita select tidak ada datanya
// ● Sebenernya mudah saja tinggal tambahkan default di case di atas itu sudah termasuk default select contoh tapi saya tidak akan menulis ulang code nya dari awal
func TestSelectChannel(t *testing.T) {

	// buat dua channel
	channel1 := make(chan string)
	channel2 := make(chan string)

	// close channel
	defer close(channel1)
	defer close(channel2)

	// 2 goroutine
	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	// menggunakan loop infinite karen tidak menggunakan kondisi
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++

		// ini jika di dalam select bisa disebut dengan select default
		default:
			fmt.Println("Menunggu")
		}

		// kondisi untuk menghentikan loop karena counter nya hanya 2 dan menggunakan break
		if counter == 2 {
			break
		}
	}
}
