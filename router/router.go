package router

import (
	"backend-konstruksi/handler"
	"net/http"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Routing
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Konstruksi API"))
	})
	mux.HandleFunc("/kontraktor/tambah", handler.CreateKontraktor)
	mux.HandleFunc("/proyek", handler.GetAllProyek)
	mux.HandleFunc("/proyek/create", handler.CreateProyek)
	mux.HandleFunc("/proyek/", handler.GetProyekByID) // Will handle both GET and DELETE
	mux.HandleFunc("/proyek/update/", handler.UpdateProyek)
	mux.HandleFunc("/proyek/delete/", handler.DeleteProyek)

	return mux
}
