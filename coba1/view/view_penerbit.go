package view

import (
	"bufio"
	"coba1/model"
	"coba1/node"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InsertPenerbit() {
	reader := bufio.NewReader(os.Stdin)
	var id int
	var nama, alamat, telepon string

	fmt.Print("Masukkan ID Penerbit: ")
	idStr, _ := reader.ReadString('\n')
	id, _ = strconv.Atoi(strings.TrimSpace(idStr))

	fmt.Print("Masukkan Nama Penerbit: ")
	nama, _ = reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Masukkan Alamat: ")
	alamat, _ = reader.ReadString('\n')
	alamat = strings.TrimSpace(alamat)

	fmt.Print("Masukkan Telepon: ")
	telepon, _ = reader.ReadString('\n')
	telepon = strings.TrimSpace(telepon)

	p := node.Penerbit{
		ID:      id,
		Nama:    nama,
		Alamat:  alamat,
		Telepon: telepon,
	}

	if model.CreatePenerbit(p) {
		fmt.Println("== Penerbit berhasil ditambahkan ==")
	} else {
		fmt.Println("== Gagal menambahkan penerbit ==")
	}
}

func ViewPenerbit() {
	fmt.Println("\n=== Daftar Penerbit ===")
	for i, p := range model.ReadPenerbit() {
		fmt.Println("Penerbit ke -", i+1)
		fmt.Println("ID\t:", p.ID)
		fmt.Println("Nama\t:", p.Nama)
		fmt.Println("Alamat\t:", p.Alamat)
		fmt.Println("Telepon\t:", p.Telepon)
		fmt.Println()
	}
}

func UpdatePenerbit() {
	reader := bufio.NewReader(os.Stdin)
	var id int
	var nama, alamat, telepon string

	fmt.Print("ID Penerbit yang akan diupdate: ")
	idStr, _ := reader.ReadString('\n')
	id, _ = strconv.Atoi(strings.TrimSpace(idStr))

	fmt.Print("Nama Baru: ")
	nama, _ = reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Alamat Baru: ")
	alamat, _ = reader.ReadString('\n')
	alamat = strings.TrimSpace(alamat)

	fmt.Print("Telepon Baru: ")
	telepon, _ = reader.ReadString('\n')
	telepon = strings.TrimSpace(telepon)

	p := node.Penerbit{
		ID:      id,
		Nama:    nama,
		Alamat:  alamat,
		Telepon: telepon,
	}

	if model.UpdatePenerbit(p, id) {
		fmt.Println("== Penerbit berhasil diupdate ==")
	} else {
		fmt.Println("== Gagal update penerbit ==")
	}
}

func DeletePenerbit() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("ID Penerbit yang akan dihapus: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	if model.DeletePenerbit(id) {
		fmt.Println("== Penerbit berhasil dihapus ==")
	} else {
		fmt.Println("== Gagal menghapus penerbit ==")
	}
}
