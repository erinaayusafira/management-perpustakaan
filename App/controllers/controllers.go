package controllers

import(
	"time"
	"io"
	"os"
	"fmt"
	"path/filepath"
	"html/template"
	"encoding/json"
	"net/http"
	
	"perpustakaan-member/app/models"
	"perpustakaan-member/app/repositories"
)


// endpoint DATA BUKU
func GetDataBuku(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories.GetBuku())
}

// endpoint DATA MEMBER
func GetDataMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repositories.GetMember())
}

// endpoint DATA STOCK
func GetDataStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repositories.GetStock())
}


//endpoint INSERT PINJAM
func InsertDataPinjam(w http.ResponseWriter, r *http.Request){

	var pinjam models.Pinjam

	pinjam.Id_member 	= r.FormValue("id_member")
	pinjam.Kode_buku 	= r.FormValue("kode_buku")
	pinjam.Qty 			= r.FormValue("qty")
	pinjam.Tgl_pinjam  	= time.Now()

	// Validation
	if len(pinjam.Id_member) == 0 {
		json.NewEncoder(w).Encode("Please suplay valid id member")
		return
	}

	if len(pinjam.Kode_buku) == 0 {
		json.NewEncoder(w).Encode("Please suplay valid kode buku")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repositories.CreatePinjam(pinjam))
}
		

//endpoint INSERT KEMBALI
func InsertDataKembali(w http.ResponseWriter, r *http.Request){

	var kembali models.Kembali
	
		kembali.Id_member 		= r.FormValue("id_member")
		kembali.Kode_buku 		= r.FormValue("kode_buku")
		kembali.Id_pinjam 		= r.FormValue("id_pinjam")
		kembali.Qty 			= r.FormValue("qty")
		kembali.Tgl_kembali		= time.Now()

		// Validation
	if len(kembali.Id_member) == 0 {
		json.NewEncoder(w).Encode("Please suplay valid id member")
		return
	}

	if len(kembali.Kode_buku) == 0 {
		json.NewEncoder(w).Encode("Please suplay valid kode buku")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repositories.CreateKembali(kembali))
}
		

// endpoint DETAIL PINJAM
func GetDataDetailPinjam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repositories.GetDetailPinjam())

}	

// endpoint DETAIL KEMBALI
func GetDataDetailKembali(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repositories.GetDetailKembali())
}	

type M map[string]interface{}

func Index(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("views/index.html"))
	if err := tmpl.Execute(w, nil);
	err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UploadFile(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		http.Error(w, "Only accept POST method", http.StatusBadRequest)
		return
	}

	basePath, _ := os.Getwd()
	reader, err := r.MultipartReader() 
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for {
		part, err := reader.NextPart() 
		if err == io.EOF {	
			break
		}
		fileLocation := filepath.Join(basePath, "files", part.FileName())
		dst, err := os.Create(fileLocation)
		if dst != nil {
			defer dst.Close()
		}
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := io.Copy(dst, part); 
		err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Write([]byte(`all files uploaded`))
}

func ListFiles(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("views/view.html"))
	if err := tmpl.Execute(w, nil);
	err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandleListFiles(w http.ResponseWriter, r *http.Request){
	files := []M{}
	basePath,_ := os.Getwd() 
	filesLocation := filepath.Join(basePath, "files") 

	err := filepath.Walk(filesLocation, func(path string, info os.FileInfo, err error) error{
		if err != nil {
			return err 
		}
		if info.IsDir(){
			return nil
		}
		files = append(files, M{"filename": info.Name(), "path": path})
		return nil 
	})
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(files)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(res)
}

func Download(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm();
	err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	path := r.FormValue("path")
	f, err := os.Open(path)
	if f != nil{
		defer f.Close()
	}
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	contentDisposition := fmt.Sprintf("attachment; filename=%s", f.Name()) 
	w.Header().Set("Content-Disposition", contentDisposition)

	if _, err := io.Copy(w, f); 
	err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
