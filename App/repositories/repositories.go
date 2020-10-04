package repositories

import(
	"database/sql"
	"perpustakaan-member/app/config"
	"perpustakaan-member/app/libraries"
	"perpustakaan-member/app/models"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

//KONEKSI DB
func init(){
	db, err = sql.Open(config.GetString("database.driverName"), config.GetString("database.dataSourceName"))
	libraries.CheckError(err)

	err = db.Ping()
	libraries.CheckError(err)
}


//GET DATA BUKU
func GetBuku()([]models.Buku) {
	rows, err := db.Query("select * from tb_buku")
	libraries.CheckError(err)
	
	var bukus []models.Buku

    for rows.Next() {	
        var each = models.Buku{}
    	var err = rows.Scan(&each.Id_buku, &each.Kode_buku,&each.Judul_buku, &each.Jumlah_halaman,&each.Penulis, &each.Tahun_terbit, &each.Stok )

     	libraries.CheckError(err)

		bukus = append(bukus, each)
			
	}
	return bukus
}	


// GET DATA MEMBER
func GetMember()([]models.Member) {
	rows, err := db.Query("select * from tb_member")
	libraries.CheckError(err)
	
	var members []models.Member

    for rows.Next() {
        var each = models.Member{}
        var err = rows.Scan(&each.No_member,&each.Id_member, &each.Nama, &each.Umur,&each.Alamat, &each.Telepon )

        libraries.CheckError(err)

		members = append(members, each)			
	}
    return members
}	


// GET DATA STOK
func GetStock()([]models.Stok) {
	rows, err := db.Query("select kode_buku, judul_buku, stok from tb_buku")
	libraries.CheckError(err)
	
	var stoks []models.Stok

    for rows.Next() {
        var each = models.Stok{}
        var err = rows.Scan(&each.Kode_buku, &each.Judul_buku, &each.Stok)

        libraries.CheckError(err)

		stoks = append(stoks, each)
    }
	return stoks
}	


//INSERT PINJAM
func CreatePinjam(pinjam models.Pinjam)(models.Pinjam){

		_, err = db.Exec("INSERT INTO tb_pinjaman (id_member, kode_buku, qty, tgl_pinjam) values (?,?,?,?)",
			pinjam.Id_member,
			pinjam.Kode_buku,
			pinjam.Qty,
			pinjam.Tgl_pinjam,
	)
		libraries.CheckError(err)

		return pinjam
}

// INSERT KEMBALI
func CreateKembali(kembali models.Kembali)(models.Kembali){
	
		_, err = db.Exec("INSERT INTO tb_kembalian (id_member, kode_buku, id_pinjam, qty, tgl_kembali) values (?,?,?,?,?)",
			kembali.Id_member,
			kembali.Kode_buku,
			kembali.Id_pinjam,
			kembali.Qty,
			kembali.Tgl_kembali,
	)
		libraries.CheckError(err)

		return kembali
}
		
// GET DETAIL PINJAM
func GetDetailPinjam()([]models.Detail_Pinjam) {

	rows, err := db.Query("select p.id_pinjam, p.id_member, p.kode_buku, b.judul_buku, p.qty, p.tgl_pinjam from tb_pinjaman p inner join tb_buku b on p.kode_buku = b.kode_buku order by id_pinjam")
	
	libraries.CheckError(err)
	
	var detail_pinjam []models.Detail_Pinjam

    for rows.Next() {
        var each = models.Detail_Pinjam{}
        var err = rows.Scan(&each.Id_pinjam, &each.Id_member, &each.Kode_buku, &each.Judul_buku, &each.Qty, &each.Tgl_pinjam)

        libraries.CheckError(err)

		detail_pinjam = append(detail_pinjam, each)
			
	}
	return detail_pinjam
}	

// GET DETAIL KEMBALI
func GetDetailKembali()([]models.Detail_Kembali) {
	
	rows, err := db.Query("select k.id_kembali, k.id_member, b.kode_buku, b.judul_buku, k.qty, k.id_pinjam, p.tgl_pinjam, k.tgl_kembali from tb_kembalian k inner join tb_pinjaman p on k.id_pinjam = p.id_pinjam inner join tb_buku b on k.kode_buku = b.kode_buku order by id_kembali")
	libraries.CheckError(err)
	
	var detail_kembali []models.Detail_Kembali

    for rows.Next() {
    	var each = models.Detail_Kembali{}
    	var err = rows.Scan(&each.Id_kembali,&each.Id_member, &each.Kode_buku, &each.Judul_buku, &each.Qty,&each.Id_pinjam, &each.Tgl_pinjam, &each.Tgl_kembali)

        libraries.CheckError(err)

		detail_kembali = append(detail_kembali, each)
			
	}
	return detail_kembali
}	


