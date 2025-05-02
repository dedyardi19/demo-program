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

func InsertBuku() {
	reader := bufio.NewReader(os.Stdin)
	var id, tahun, idPenerbit int
	var judul, penulis string

	fmt.Print("Masukkan ID Buku: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, _ = strconv.Atoi(idStr)

	fmt.Print("Masukkan Judul Buku: ")
	judul, _ = reader.ReadString('\n')
	judul = strings.TrimSpace(judul)

	fmt.Print("Masukkan Penulis Buku: ")
	penulis, _ = reader.ReadString('\n')
	penulis = strings.TrimSpace(penulis)

	fmt.Print("Masukkan Tahun Terbit: ")
	tahunStr, _ := reader.ReadString('\n')
	tahunStr = strings.TrimSpace(tahunStr)
	tahun, _ = strconv.Atoi(tahunStr)

	fmt.Print("Masukkan ID Penerbit: ")
	idPenerbitStr, _ := reader.ReadString('\n')
	idPenerbitStr = strings.TrimSpace(idPenerbitStr)
	idPenerbit, _ = strconv.Atoi(idPenerbitStr)

	if model.SearchPenerbit(idPenerbit) {
		buku := node.Buku{
			ID:         id,
			Judul:      judul,
			Penulis:    penulis,
			Tahun:      tahun,
			IDPenerbit: idPenerbit,
		}
		model.CreateBuku(buku)
		fmt.Println("== Buku berhasil ditambahkan ==")
	} else {
		fmt.Println("== ID Penerbit tidak ditemukan ==")
	}
}

func ViewBuku() {
	fmt.Println("=== Daftar Buku ===")
	for i, b := range model.ReadBuku() {
		fmt.Println("Buku ke -", i+1)
		fmt.Println("ID\t\t:", b.ID)
		fmt.Println("Judul\t\t:", b.Judul)
		fmt.Println("Penulis\t\t:", b.Penulis)
		fmt.Println("Tahun Terbit\t:", b.Tahun)
		fmt.Println("Penerbit\t:", model.GetNamaPenerbit(b.IDPenerbit))
		fmt.Println()
	}
}

func UpdateBuku() {
	reader := bufio.NewReader(os.Stdin)
	var id, tahun, idPenerbit int
	var judul, penulis string

	fmt.Print("ID Buku yang akan diupdate: ")
	idStr, _ := reader.ReadString('\n')
	id, _ = strconv.Atoi(strings.TrimSpace(idStr))

	fmt.Print("Judul Baru: ")
	judul, _ = reader.ReadString('\n')
	judul = strings.TrimSpace(judul)

	fmt.Print("Penulis Baru: ")
	penulis, _ = reader.ReadString('\n')
	penulis = strings.TrimSpace(penulis)

	fmt.Print("Tahun Terbit Baru: ")
	tahunStr, _ := reader.ReadString('\n')
	tahun, _ = strconv.Atoi(strings.TrimSpace(tahunStr))

	fmt.Print("ID Penerbit Baru: ")
	idPenerbitStr, _ := reader.ReadString('\n')
	idPenerbit, _ = strconv.Atoi(strings.TrimSpace(idPenerbitStr))

	if model.SearchPenerbit(idPenerbit) {
		buku := node.Buku{
			ID:         id,
			Judul:      judul,
			Penulis:    penulis,
			Tahun:      tahun,
			IDPenerbit: idPenerbit,
		}
		if model.UpdateBuku(buku, id) {
			fmt.Println("== Buku berhasil diupdate ==")
		} else {
			fmt.Println("== Gagal update buku ==")
		}
	} else {
		fmt.Println("== ID Penerbit tidak ditemukan ==")
	}
}

func DeleteBuku() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("ID Buku yang akan dihapus: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	if model.DeleteBuku(id) {
		fmt.Println("== Buku berhasil dihapus ==")
	} else {
		fmt.Println("== Gagal menghapus buku ==")
	}
}
