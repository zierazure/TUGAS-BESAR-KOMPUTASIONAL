package main

import "fmt"

func main() {
	// Mendefinisikan struktur data tiket untuk mengelompokkan informasi
	type Tiket struct {
		Kategori string
		Harga    float64
		Stok     int
	}

	// Inisialisasi data tiket
	tiketVIP := Tiket{"VIP (Front Row)", 2500000, 5}
	tiketFest := Tiket{"Festival", 120000, 10}

	// Tampilan awal untuk pengguna
	fmt.Println("✨WELCOME TO DHERA-TICKET✨")
	fmt.Println("Pilihan Tiket Gemilang Fest💫")
	fmt.Printf("1. %s - Rp%.0f (Sisa: %d)\n", tiketVIP.Kategori, tiketVIP.Harga, tiketVIP.Stok)
	fmt.Printf("2. %s - Rp%.0f (Sisa: %d)\n", tiketFest.Kategori, tiketFest.Harga, tiketFest.Stok)

	// Mengambil data identitas dan pesanan tiket dari pembeli
	var nama string
	var pilihan int
	var jumlah int

	fmt.Print("\nNama kamu: ")
	fmt.Scan(&nama)
	fmt.Print("\nPilih kategori (1/2): ")
	fmt.Scan(&pilihan)
	fmt.Print("Mau beli berapa tiket? ")
	fmt.Scan(&jumlah)

	// Inisialisasi variabel bantuan untuk kalkulasi
	var hargaTotal float64
	namaPilihan := ""
	stokCukup := false

	// Validasi berdasarkan kategori yang dipilih
	if pilihan == 1 {
		// Menambah pengecekan jumlah tiket harus lebih dari 0
		if jumlah > 0 && jumlah <= tiketVIP.Stok {
			hargaTotal = float64(jumlah) * tiketVIP.Harga
			namaPilihan = tiketVIP.Kategori
			stokCukup = true
		}
	} else if pilihan == 2 {
		// Menambah pengecekan jumlah tiket harus lebih dari 0
		if jumlah > 0 && jumlah <= tiketVIP.Stok {
			hargaTotal = float64(jumlah) * tiketFest.Harga
			namaPilihan = tiketFest.Kategori
			stokCukup = true
		}
	}

	// Error handling jika input jumlah tidak valid (0 atau minus)
	if jumlah <= 0 {
		fmt.Println("\n🚫 EH? Masa belinya 0 atau minus?")
		fmt.Println("Masukin jumlah yang bener dong!")
		return
	}
	// Program langsung berhenti di sini

	// Output akhir berdasarkan hasil validasi stok
	if stokCukup {
		pajak := hargaTotal * 0.1 // Pajak 10%
		fmt.Println("\n YAY!! Booking Berhasil!")
		fmt.Printf("Nama: %s\nTiket: %s\nJumlah: %d\nTotal Bayar (+Pajak 10%%): Rp%.0f\n", nama, namaPilihan, jumlah, hargaTotal+pajak)
		fmt.Println("Segera lakukan pembayaran ya!")
		fmt.Println("Sampai jumpa di Gemilang Fest👋🏻💫")

		// Pesan jika stok tidak mencukupi
	} else {
		fmt.Println("\nMohon maaf, anda telat! Tiketnya habis atau jumlah melebihi tiket yang tersisa.")
		fmt.Println("\nCoba war lagi besok!")
		fmt.Println("\nTerima kasih💫")
	}
}
