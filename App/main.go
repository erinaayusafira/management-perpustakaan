package main

import(
	"net/http"
	"fmt"

	"perpustakaan-member/app/controllers"
	"perpustakaan-member/app/config"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/GetDataBuku", controllers.GetDataBuku).Methods("GET")
	router.HandleFunc("/GetDataMember", controllers.GetDataMember).Methods("GET")
	router.HandleFunc("/GetStock", controllers.GetDataStock).Methods("GET")
	router.HandleFunc("/InsertPinjam", controllers.InsertDataPinjam).Methods("POST")
	router.HandleFunc("/InsertKembalian", controllers.InsertDataKembali).Methods("POST")
	router.HandleFunc("/GetDetailPinjam", controllers.GetDataDetailPinjam).Methods("GET")
	router.HandleFunc("/GetDetailKembali", controllers.GetDataDetailKembali).Methods("GET")
	router.HandleFunc("/", controllers.Index)
	router.HandleFunc("/upload", controllers.UploadFile)
	router.HandleFunc("/main.go/listfile", controllers.ListFiles)
	router.HandleFunc("/list-files", controllers.HandleListFiles)
	router.HandleFunc("/download", controllers.Download)

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(config.GetString("server.address"), router)
}