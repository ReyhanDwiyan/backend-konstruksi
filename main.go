package main

import (
	"fmt"
	"log"
	"net/http"

	"backend-konstruksi/config"
	"backend-konstruksi/router"
)

func main() {
	// Inisialisasi koneksi database
	config.ConnectDB()

	// Setup router
	r := router.SetupRouter()

	// Jalankan server di port 8080
	port := ":8080"
	fmt.Println("Server berjalan di http://localhost" + port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal("Gagal menjalankan server: ", err)
	}
}
