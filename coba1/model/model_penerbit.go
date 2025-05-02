package model

import "coba1/node"

var DaftarPenerbit node.ListPenerbit

func CreatePenerbit(buk node.Penerbit) bool {
	tempLL := node.ListPenerbit{
		Data: buk,
		Link: nil,
	}
	if DaftarPenerbit.Link == nil {
		DaftarPenerbit.Link = &tempLL
		return true
	} else {
		temp := &DaftarPenerbit
		for temp.Link != nil {
			temp = temp.Link
		}
		temp.Link = &tempLL
		return true
	}
	return false
}

func ReadPenerbit() []node.Penerbit {
	daftarPenerbit := []node.Penerbit{}
	temp := &DaftarPenerbit
	for temp.Link != nil {
		daftarPenerbit = append(daftarPenerbit, temp.Link.Data)
		temp = temp.Link
	}
	return daftarPenerbit
}

func UpdatePenerbit(emp node.Penerbit, id int) bool {
	temp := DaftarPenerbit.Link
	for temp != nil {
		if temp.Data.ID == id {
			temp.Data = emp
			return true
		}
		temp = temp.Link
	}
	return false
}

func DeletePenerbit(id int) bool {
	temp := &DaftarPenerbit
	for temp.Link != nil {
		if temp.Link.Data.ID == id {
			temp.Link = temp.Link.Link
			return true
		}
		temp = temp.Link
	}
	return false
}

func SearchPenerbit(id int) bool {
	temp := DaftarPenerbit.Link
	for temp != nil {
		if temp.Data.ID == id {
			return true
		}
		temp = temp.Link
	}
	return false
}

func GetNamaPenerbit(id int) string {
	temp := DaftarPenerbit.Link
	for temp != nil {
		if temp.Data.ID == id {
			return temp.Data.Nama
		}
		temp = temp.Link
	}
	return ""
}
