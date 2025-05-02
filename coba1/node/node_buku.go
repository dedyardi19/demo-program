package node

type Buku struct {
	ID         int
	Judul      string
	Penulis    string
	Tahun      int
	IDPenerbit int
}

type ListBuku struct {
	Data Buku
	Link *ListBuku
}
