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

//Get the cards from a public API and returns a list with all the results
func GetCardFromAPI(name string)[] models.Card {
	apiURL := "https://api.magicthegathering.io/v1/cards?name=" + name

	//Creates the GET request
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error trying to stablish the request:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("The request returns an invalid state:", resp.Status)
	}

	// Decode the JSON
	var cardResponse models.CardResponse
	err = json.NewDecoder(resp.Body).Decode(&cardResponse)
	if err != nil {
		fmt.Println("Error trying to decode the JSON:", err)
	}

	if len(cardResponse.Cards) == 0 {
		fmt.Println("There are no cards with that name")
	}

	// Extract all the required information from the cards
	var cards []models.Card
	var card models.Card
	for _, cardTemp := range cardResponse.Cards {
		card.Name = cardTemp.Name
		card.Text = cardTemp.Text
		card.Number = cardTemp.Number
		cards = append(cards, card)
	}
	return cards
}

//Returns a card name from the view
func SearchCard(w http.ResponseWriter, r *http.Request)string{
	r.ParseForm()

	CardName := r.FormValue("CardName")

	return CardName
}

//Add the selected card into the API from the view
func AddCardForm(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	CardName := r.FormValue("AddCardName")
	CardText := r.FormValue("AddCardText")
	CardNumber := r.FormValue("AddCardNumber")
	var deck models.Deck
	var card models.Card

	commanderURL := strings.Split(r.URL.Path, "/")[1]

	for i := 0; i < len(models.Users); i++ {
		for j := 0; j < len(models.Users[i].Decks); j++ {
			if models.Users[i].Decks[j].Commander == commanderURL{
				deck = models.Users[i].Decks[j]
			}
		}
	}

	card.Name = CardName
	card.Text = CardText
	card.Deck_id = deck.ID
	card.Number = CardNumber

	// Turn the object into a JSON body
    jsonBody, err := json.Marshal(card)
    if err != nil {
        fmt.Println("Error trying to convert into a JSON:", err)
        return
    }

    apiUrl := "http://localhost:8081/cards"
    requestBody := bytes.NewBuffer(jsonBody)

	// Stablish the POST request
    resp, err := http.Post(apiUrl, "application/json", requestBody)
    if err != nil {
        fmt.Println("Error trying to request the POST method:", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Println("The request returns an invalid state:", resp.Status)
        return
    }

	http.Redirect(w, r, "/"+commanderURL, http.StatusSeeOther)
}

//Get all the card from an specific deck using the API
func GetCardsFromDeck(ID string)[]models.Card{

    url := "http://localhost:8081/decks/"+ID+"/cards"

    //Creates the GET request
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error creating the request:", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Println("La solicitud devolvió un código de estado no válido:", resp.Status)
    }

    // Decode the JSON into an object
    var cards []models.Card
    err = json.NewDecoder(resp.Body).Decode(&cards)
    if err != nil {
        fmt.Println("Error decoding the JSON:", err)
    }
	return cards
}

//Get all the card info from the view using a form and the URL in order to delete it
func DeleteCard(w http.ResponseWriter, r *http.Request){

	CardNameTemp := strings.Split(strings.Split(r.URL.Path, "/")[4], "%20")
	var CardName string

	for i := 0; i < len(CardNameTemp); i++ {
		CardName = CardNameTemp[i]+" "
	}
	CardName = CardName[:len(CardName)-1]

	var cards []models.Card
	var ID string

	commanderURL := strings.Split(r.URL.Path, "/")[1]

	for i := 0; i < len(models.Users); i++ {
		for j := 0; j < len(models.Users[i].Decks); j++ {
			if models.Users[i].Decks[j].Commander == commanderURL{
				cards = GetCardsFromDeck(strconv.Itoa(models.Users[i].Decks[j].ID))
				for y := 0; y < len(cards); y++ {
					if cards[y].Name == CardName{
						ID = strconv.Itoa(cards[y].ID)
					}
				}
			}
		}
	}

	DeleteCardPage(ID)

	http.Redirect(w, r, "/"+commanderURL, http.StatusSeeOther)

}

//Deletes a card from the API
func DeleteCardPage(ID string){

    apiUrl := "http://localhost:8081/cards/" + ID

    //Creates a DELETE request
    req, err := http.NewRequest("DELETE", apiUrl, nil)
    if err != nil {
        fmt.Println("Error al crear la solicitud DELETE:", err)
        return
    }

    // Stablish the DELETE request
    client := &http.Client{}
    resp, _ := client.Do(req)
    defer resp.Body.Close()
}