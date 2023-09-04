package controllers

import (
	"FEModule/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func ConvertirSTR(texto string)int{
	numero, err := strconv.Atoi(texto)
	if err != nil {
		fmt.Println("Error al convertir la cadena a entero:", err)
	} 
	return numero
}

func GetAllUSersAPI(w http.ResponseWriter, r *http.Request) {
    GetAllUsersSQL()
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(models.Users)
}

func GetAllDecksAPI(w http.ResponseWriter, r *http.Request) {
    GetAllDecksSQL()
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(models.Decks)
}

func GetUser_DecksAPI(w http.ResponseWriter, r *http.Request) {
	// Obtiene el ID del usuario desde la URL
	params := mux.Vars(r)
	id := params["id"]

	var decks []models.Deck = GetDecksSQL(ConvertirSTR(id))

	// Responde con los decks del usuario en formato JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(decks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetDeck_CardsAPI(w http.ResponseWriter, r *http.Request) {
	// Obtiene el ID del usuario desde la URL
	params := mux.Vars(r)
	id := params["id"]

	var cards []models.Card = GetCardsSQL(ConvertirSTR(id))

	// Responde con los decks del usuario en formato JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cards); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AddUserAPI(w http.ResponseWriter, r *http.Request) {
	// Decodifica el JSON del cuerpo de la solicitud en una estructura de usuario
	var p models.User
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.ID = AddUserSQL(p.Name, p.Age, p.Email)

	// Responde con el usuario creado en formato JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AddDeckAPI(w http.ResponseWriter, r *http.Request) {
	// Decodifica el JSON del cuerpo de la solicitud en una estructura de usuario
	var p models.Deck
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.ID = AddDeckSQL(p.Commander, p.Powerlvl, p.Colors, p.User_id)

	// Responde con el usuario creado en formato JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AddCardAPI(w http.ResponseWriter, r *http.Request) {
	// Decodifica el JSON del cuerpo de la solicitud en una estructura de usuario
	var card models.Card
	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	card.ID = AddCardSQL(card.Name, card.Deck_id, card.Text, ConvertirSTR(card.Number))

	// Responde con el usuario creado en formato JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(card); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UpdateUserAPI(w http.ResponseWriter, r *http.Request) {

	// Obtiene el ID del usuario desde la URL
	params := mux.Vars(r)
	id := params["id"]
	// Decodifica el JSON del cuerpo de la solicitud en una estructura de usuario
	var p models.User
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.ID = ConvertirSTR(id)
	UpdateUserSQL(p)
	models.UpdateUser(ConvertirSTR(id), p.Name, p.Age, p.Email)

	// Responde con el usuario creado en formato JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UpdateDeckAPI(w http.ResponseWriter, r *http.Request) {

	// Obtiene el ID del usuario desde la URL
	params := mux.Vars(r)
	id := params["id"]
	// Decodifica el JSON del cuerpo de la solicitud en una estructura de usuario
	var p models.Deck
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.ID = ConvertirSTR(id)
	UpdateDeckSQL(p)
	models.UpdateDeck(ConvertirSTR(id), p.Commander, p.Powerlvl, p.Colors)

	// Responde con el usuario creado en formato JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteUserAPI(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
	id := params["id"]

	models.RemoveUser(ConvertirSTR(id))
	DeleteUserSQL(ConvertirSTR(id))

    http.NotFound(w, r)
}

func DeleteDeckAPI(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
	id := params["id"]

	models.RemoveDeck(ConvertirSTR(id))
	DeleteDeckSQL(ConvertirSTR(id))

    http.NotFound(w, r)
}

func DeleteCardAPI(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
	id := params["id"]

	models.RemoveCard(ConvertirSTR(id))
	DeleteCardSQL(ConvertirSTR(id))

    http.NotFound(w, r)
}