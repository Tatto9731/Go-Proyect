package main

import (
	"BEModule/controllers"
	"FEModule/views"
	"net/http"
	"github.com/gorilla/mux"
)

func main(){
	//Crea un enrutador usando Gorilla Mux
	r := mux.NewRouter()

	r.HandleFunc("/", views.MainPage)
	r.HandleFunc("/add", views.AddPage)
	r.HandleFunc("/add/process", controllers.AddForm)
	r.HandleFunc("/{id}/add/deck", views.AddDeckPage)
	r.HandleFunc("/{id}/add/deck/process", controllers.AddDeckForm)
	r.HandleFunc("/{id}/remove", views.RemovePage)
	r.HandleFunc("/{commander}", views.DeckPage)
	r.HandleFunc("/{commander}/remove/deck", views.RemoveDeckPage)
	r.HandleFunc("/{commander}/remove/deck/process", controllers.RemoveDeckForm)
	r.HandleFunc("/{id}/remove/process", controllers.RemoveForm)
	r.HandleFunc("/{id}/update", views.UpdatePage)
	r.HandleFunc("/{id}/update/process", controllers.UpdateForm)
	r.HandleFunc("/{commander}/update/deck", views.UpdateDeckPage)
	r.HandleFunc("/{commander}/update/deck/process", controllers.UpdateDeckForm)
	r.HandleFunc("/{commander}/add/card", views.CardDeckPage)
	r.HandleFunc("/{commander}/add/card/process", controllers.AddCardForm)
	r.HandleFunc("/{commander}/remove/card/{name}", views.RemoveCardPage)
	r.HandleFunc("/{commander}/remove/card/{name}/process", controllers.DeleteCard)
	
	// Inicia el servidor HTTP
	http.ListenAndServe(":8080", r)
}