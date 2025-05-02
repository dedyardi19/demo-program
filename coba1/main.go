package main

import (
	"bufio"
	"coba1/model"
	"coba1/node"
	"coba1/view"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func menuUtama() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Menu Buku")
		fmt.Println("2. Menu Penerbit")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih menu: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input tidak valid. Masukkan angka.")
			continue
		}

		switch choice {
		case 1:
			menuBuku()
		case 2:
			menuPenerbit()
		case 3:
			fmt.Println("Terima kasih! Program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}

func menuBuku() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== MENU BUKU ===")
		fmt.Println("1. Tambah Data Buku")
		fmt.Println("2. Tampilkan Data Buku")
		fmt.Println("3. Update Data Buku")
		fmt.Println("4. Hapus Data Buku")
		fmt.Println("5. Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input tidak valid. Masukkan angka.")
			continue
		}

		switch choice {
		case 1:
			view.InsertBuku()
		case 2:
			view.ViewBuku()
		case 3:
			view.UpdateBuku()
		case 4:
			view.DeleteBuku()
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}

func menuPenerbit() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== MENU PENERBIT ===")
		fmt.Println("1. Tambah Data Penerbit")
		fmt.Println("2. Tampilkan Data Penerbit")
		fmt.Println("3. Update Data Penerbit")
		fmt.Println("4. Hapus Data Penerbit")
		fmt.Println("5. Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input tidak valid. Masukkan angka.")
			continue
		}

		switch choice {
		case 1:
			view.InsertPenerbit()
		case 2:
			view.ViewPenerbit()
		case 3:
			view.UpdatePenerbit()
		case 4:
			view.DeletePenerbit()
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}

func main() {

	Buk1 := node.Penerbit{
		ID:      1,
		Nama:    "Gramedia Pustaka Utama",
		Alamat:  "Jakarta",
		Telepon: "0218888",
	}

	Buk2 := node.Penerbit{
		ID:      2,
		Nama:    "Penerbit Erlangga",
		Alamat:  "Jakarta",
		Telepon: "0229999",
	}

	Buk3 := node.Penerbit{
		ID:      3,
		Nama:    "Mizan Publishing",
		Alamat:  "Bandung",
		Telepon: "0274333",
	}

	Buk4 := node.Penerbit{
		ID:      4,
		Nama:    "Elex Media Komputindo",
		Alamat:  "Jakarta",
		Telepon: "0274354",
	}

	Buk5 := node.Penerbit{
		ID:      5,
		Nama:    "Greenbook.ID",
		Alamat:  "Yogyakarta",
		Telepon: "0274369",
	}

	model.CreatePenerbit(Buk1)
	model.CreatePenerbit(Buk2)
	model.CreatePenerbit(Buk3)
	model.CreatePenerbit(Buk4)
	model.CreatePenerbit(Buk5)

	fmt.Println(model.ReadPenerbit())

	menuUtama()
}
