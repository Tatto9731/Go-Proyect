package controllers

import (
	"FEModule/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func GetDeckID(commander string) int {
	var ID int
	for i := 0; i < len(models.Users); i++ {
		for j := 0; j < len(models.Users[i].Decks); j++ {
			if models.Users[i].Decks[j].Commander == commander {
				ID = models.Users[i].Decks[j].ID
			}
		}
	}
	return ID
}

//Gets the deck info from the view using a form and add it to the API
func AddDeckForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var deck models.Deck

	commander := r.FormValue("AddCommander")
	powerlvl := r.FormValue("AddPowerlvl")
	colors := r.FormValue("AddColors")
	user_id := r.FormValue("AddUser_id")

	deck.Commander, deck.Powerlvl, deck.Colors, deck.User_id = commander, ConvertSTR(powerlvl), colors, ConvertSTR(user_id)

	// Convert the object into a JSON body
	jsonBody, err := json.Marshal(deck)
	if err != nil {
		fmt.Println("Error trying to convert into JSON:", err)
		return
	}

	apiUrl := "http://localhost:8081/decks"
	requestBody := bytes.NewBuffer(jsonBody)

	// Request a POST method
	resp, err := http.Post(apiUrl, "application/json", requestBody)
	if err != nil {
		fmt.Println("Error trying to create the POST request:", err)
		return
	}
	defer resp.Body.Close()

	// Verificar el código de estado de la respuesta
	if resp.StatusCode != http.StatusOK {
		fmt.Println("The request returns an invalid state:", resp.Status)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

//Gets the commander info from the URL in order to delete the deck
func RemoveDeck(w http.ResponseWriter, r *http.Request) {

	commander := strings.Split(r.URL.Path, "/")[1]

	DeleteDeckPage(strconv.Itoa(GetDeckID(commander)))

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

//Gets the deck info from the view usign a form and the URL in order to update it
func UpdateDeckForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	commanderform := r.FormValue("UpdateCommander")
	powerlvl := r.FormValue("UpdatePowerlvl")
	colors := r.FormValue("UpdateColors")
	var deck models.Deck

	commanderURL := strings.Split(r.URL.Path, "/")[1]

    //Find the deck into the user information
	for i := 0; i < len(models.Users); i++ {
		for j := 0; j < len(models.Users[i].Decks); j++ {
			if models.Users[i].Decks[j].Commander == commanderURL {
				deck = models.Users[i].Decks[j]
			}
		}
	}
	deck.Commander = commanderform
	deck.Powerlvl = ConvertSTR(powerlvl)
	deck.Colors = colors
	PutDeckPage(deck)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

//Get the decks info from the API and returns a slice with all the decks
func GetDecksFromUser(ID string) []models.Deck {

	url := "http://localhost:8081/decks/" + ID

	// Creates a GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error creating the request:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("The request returns an invalid state:", resp.Status)
	}

	var decks []models.Deck
	err = json.NewDecoder(resp.Body).Decode(&decks)
	if err != nil {
		fmt.Println("Error decoding the JSON:", err)
	}
	return decks
}

//Gets all decks from all users from the API and returns a slice
func GetAllDecks() []models.Deck {

	url := "http://localhost:8081/users/decks"

	//Creates the GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error creating the request:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("The request returns an invalid state:", resp.Status)
	}

	var p []models.Deck
	err = json.NewDecoder(resp.Body).Decode(&p)
	if err != nil {
		fmt.Println("Error decoding the JSON:", err)
	}
	return p
}

//Deletes a deck using the API
func DeleteDeckPage(ID string) {

	apiUrl := "http://localhost:8081/decks/" + ID

	//Creates a DELETE request
	req, err := http.NewRequest("DELETE", apiUrl, nil)
	if err != nil {
		fmt.Println("Error creating the DELETE request:", err)
		return
	}

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
}

//Updates the deck information using the API
func PutDeckPage(p models.Deck) {
	// Convert the object into a JSON body
	jsonBody, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error trying to convert into a JSON:", err)
		return
	}

	apiUrl := "http://localhost:8081/decks/" + strconv.Itoa(p.ID)

	//Creates the PUT request
	req, err := http.NewRequest("PUT", apiUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Error creating the PUT request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error trying to stablish the PUT request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("The PUT request returns an invalid state:", resp.Status)
		return
	}
}
