package models

import "time"

//STRUCT
type Member struct {
	No_member		int			`json:"no_member"`
    Id_member	 	string		`json:"id_member"`
    Nama 			string		`json:"nama"`
    Umur 			string		`json:"umur"`
	Alamat 			string		`json:"alamat"`
	Telepon 		string		`json:"telepon"`
}

type Buku struct {
	Id_buku			int			`json:"id_buku"`
	Kode_buku 		string		`json:"kode_buku"`
    Judul_buku 		string		`json:"judul_buku"`
    Jumlah_halaman 	int			`json:"jumlah_halaman"`
	Penulis 		string		`json:"penulis"`
	Tahun_terbit 	int64		`json:"tahun_terbit"`
	Stok			int64		`json:"stok"`
}

type Stok struct{
	Kode_buku 		string		`json:"kode_buku"`
	Judul_buku    	string		`json:"judul_buku"`
	Stok			int			`json:"stok"`
}

type Pinjam struct {
	Id_member	 	string		`json:"id_member"`
	Kode_buku		string		`json:"kode_buku"`
	Qty		 		string		`json:"qty"`
	Tgl_pinjam		time.Time	`json:"tgl_pinjam"`
}

type Kembali struct {
	Id_member	 	string		`json:"id_member"`
	Kode_buku 		string		`json:"kode_buku"`
	Id_pinjam 		string 		`json:"id_pinjam"`	
	Qty		 		string		`json:"qty"`
	Tgl_kembali		time.Time	`json:"tgl_kembali"`
}

type Detail_Pinjam struct {
	Id_pinjam	 	int				`json:"id_pinjam"`
	Id_member	 	string			`json:"id_member"`
	Kode_buku		string			`json:"kode_buku"`
	Judul_buku	 	string			`json:"judul_buku"`
	Qty		 		int				`json:"qty"`
	Tgl_pinjam		string			`json:"tgl_pinjam"`
}

type Detail_Kembali struct {
	Id_kembali	 	int				`json:"id_kembali"`
	Id_member	 	string			`json:"id_member"`
	Kode_buku		string			`json:"kode_buku"`
	Judul_buku	 	string			`json:"judul_buku"`
	Qty		 		int				`json:"qty"`
	Id_pinjam		int				`json:"id_pinjam"`
	Tgl_pinjam		string			`json:"tgl_pinjam"`
	Tgl_kembali 	string 			`json:"tgl_kembali"`
}

