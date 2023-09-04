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

func GetCardAPI(name string)[] models.Card {
	apiURL := "https://api.magicthegathering.io/v1/cards?name=" + name

	// Realizar la solicitud GET a la API
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error al realizar la solicitud:", err)
	}
	defer resp.Body.Close()

	// Verificar el código de estado de la respuesta
	if resp.StatusCode != http.StatusOK {
		fmt.Println("La solicitud devolvió un código de estado no válido:", resp.Status)
	}

	// Decodificar la respuesta JSON
	var cardResponse models.CardResponse
	err = json.NewDecoder(resp.Body).Decode(&cardResponse)
	if err != nil {
		fmt.Println("Error al decodificar la respuesta JSON:", err)
	}

	// Verificar si se encontraron cartas en la respuesta
	if len(cardResponse.Cards) == 0 {
		fmt.Println("No se encontraron cartas en la respuesta JSON")
	}

	// Extraer los nombres de las cartas y mostrarlos
	var cards []models.Card
	var p models.Card
	for _, card := range cardResponse.Cards {
		p.Name = card.Name
		p.Text = card.Text
		p.Number = card.Number
		cards = append(cards, p)
	}
	return cards
}

func SearchCard(w http.ResponseWriter, r *http.Request)string{
	r.ParseForm() // Parsea los datos del formulario

	CardName := r.FormValue("CardName")

	return CardName
}

func AddCardForm(w http.ResponseWriter, r *http.Request){
	r.ParseForm() // Parsea los datos del formulario

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

	// Convertir el objeto Usuario a JSON
    jsonBody, err := json.Marshal(card)
    if err != nil {
        fmt.Println("Error al convertir el objeto a JSON:", err)
        return
    }

	// URL de tu API y cuerpo de la solicitud POST
    apiUrl := "http://localhost:8081/cards"
    requestBody := bytes.NewBuffer(jsonBody)

	// Realizar la solicitud POST
    resp, err := http.Post(apiUrl, "application/json", requestBody)
    if err != nil {
        fmt.Println("Error al realizar la solicitud POST:", err)
        return
    }
    defer resp.Body.Close()

    // Verificar el código de estado de la respuesta
    if resp.StatusCode != http.StatusOK {
        fmt.Println("La solicitud devolvió un código de estado no válido:", resp.Status)
        return
    }

	http.Redirect(w, r, "/"+commanderURL, http.StatusSeeOther)
}

func GetCardsFromDeck(ID string)[]models.Card{
	// URL de la API externa
    url := "http://localhost:8081/decks/"+ID+"/cards"

    // Realiza una solicitud GET a la API
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error al realizar la solicitud:", err)
    }
    defer resp.Body.Close()

    // Verifica el código de estado de la respuesta
    if resp.StatusCode != http.StatusOK {
        fmt.Println("La solicitud devolvió un código de estado no válido:", resp.Status)
    }

    // Decodifica la respuesta JSON en tu objeto Go
    var cards []models.Card
    err = json.NewDecoder(resp.Body).Decode(&cards)
    if err != nil {
        fmt.Println("Error al decodificar la respuesta JSON:", err)
    }
	return cards
}

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

func DeleteCardPage(ID string){
	// URL de tu API para eliminar un usuario específico
    apiUrl := "http://localhost:8081/cards/" + ID

    // Crear una solicitud HTTP DELETE
    req, err := http.NewRequest("DELETE", apiUrl, nil)
    if err != nil {
        fmt.Println("Error al crear la solicitud DELETE:", err)
        return
    }

    // Realizar la solicitud DELETE
    client := &http.Client{}
    resp, _ := client.Do(req)
    defer resp.Body.Close()
}