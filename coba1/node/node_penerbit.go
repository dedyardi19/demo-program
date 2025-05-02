package node

type Penerbit struct {
	ID      int
	Nama    string
	Alamat  string
	Telepon string
}

type ListPenerbit struct {
	Data Penerbit
	Link *ListPenerbit
}
