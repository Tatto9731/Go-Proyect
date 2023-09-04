package main

import (
	"BEModule/controllers"
	//"FEModule/models"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
		//Crea un enrutador usando Gorilla Mux
		r := mux.NewRouter()

		// Manejador para la ruta "/"
		r.HandleFunc("/users", controllers.GetAllUSersAPI).Methods("GET")
		r.HandleFunc("/decks", controllers.GetAllDecksAPI).Methods("GET")
		r.HandleFunc("/decks/{id}", controllers.GetUser_DecksAPI).Methods("GET")
		r.HandleFunc("/decks/{id}/cards", controllers.GetDeck_CardsAPI).Methods("GET")
		r.HandleFunc("/users", controllers.AddUserAPI).Methods("POST")
		r.HandleFunc("/decks", controllers.AddDeckAPI).Methods("POST")
		r.HandleFunc("/cards", controllers.AddCardAPI).Methods("POST")
		r.HandleFunc("/decks/{id}", controllers.UpdateDeckAPI).Methods("PUT")
		r.HandleFunc("/users/{id}", controllers.UpdateUserAPI).Methods("PUT")
		r.HandleFunc("/users/{id}", controllers.DeleteUserAPI).Methods("DELETE")
		r.HandleFunc("/decks/{id}", controllers.DeleteDeckAPI).Methods("DELETE")
		r.HandleFunc("/cards/{id}", controllers.DeleteCardAPI).Methods("DELETE")
		
		// Inicia el servidor HTTP
		http.ListenAndServe(":8081", r)	
}