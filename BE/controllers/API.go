package controllers

import (
	"FEModule/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

//Convert a string into a int
func ConvertSTR(texto string)int{
	numero, err := strconv.Atoi(texto)
	if err != nil {
		fmt.Println("Error trying to convert the string:", err)
	} 
	return numero
}

//Gets all users info from the API
func GetAllUSersAPI(w http.ResponseWriter, r *http.Request) {
    GetAllUsersInfoSQL()
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(models.Users)
}

//Gets all decks info from the API
func GetAllDecksAPI(w http.ResponseWriter, r *http.Request) {
    GetAllDecksSQL()
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(models.Decks)
}

//Gets all decks from a user
func GetUser_DecksAPI(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	var decks []models.Deck = GetDecksSQL(ConvertSTR(id))

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(decks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Gets all cards from a deck
func GetDeck_CardsAPI(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	var cards []models.Card = GetCardsSQL(ConvertSTR(id))

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cards); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Add a user into the data base
func AddUserAPI(w http.ResponseWriter, r *http.Request) {

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.ID = AddUserSQL(user.Name, user.Age, user.Email)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Adds a deck into the data base
func AddDeckAPI(w http.ResponseWriter, r *http.Request) {

	var deck models.Deck
	if err := json.NewDecoder(r.Body).Decode(&deck); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deck.ID = AddDeckSQL(deck.Commander, deck.Powerlvl, deck.Colors, deck.User_id)


	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(deck); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Adds a card into the data base
func AddCardAPI(w http.ResponseWriter, r *http.Request) {

	var card models.Card
	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	card.ID = AddCardSQL(card.Name, card.Deck_id, card.Text, ConvertSTR(card.Number))

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(card); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Updates the user information into the date base
func UpdateUserAPI(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	var p models.User
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.ID = ConvertSTR(id)
	UpdateUserSQL(p)
	models.UpdateUser(ConvertSTR(id), p.Name, p.Age, p.Email)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Updates the deck information into the date base
func UpdateDeckAPI(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	var p models.Deck
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.ID = ConvertSTR(id)
	UpdateDeckSQL(p)
	models.UpdateDeck(ConvertSTR(id), p.Commander, p.Powerlvl, p.Colors)


	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Deletes the user from the data base
func DeleteUserAPI(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
	id := params["id"]

	models.RemoveUser(ConvertSTR(id))
	DeleteUserSQL(ConvertSTR(id))

    http.NotFound(w, r)
}

//Deletes the deck from the data base
func DeleteDeckAPI(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
	id := params["id"]

	models.RemoveDeck(ConvertSTR(id))
	DeleteDeckSQL(ConvertSTR(id))

    http.NotFound(w, r)
}

//Deletes the card from the data base
func DeleteCardAPI(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
	id := params["id"]

	models.RemoveCard(ConvertSTR(id))
	DeleteCardSQL(ConvertSTR(id))

    http.NotFound(w, r)
}