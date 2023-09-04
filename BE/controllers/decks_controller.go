package controllers

import (
	"FEModule/models"
	"net/http"
	"encoding/json"
	"bytes"
	"fmt"
	"strings"
	"strconv"
)

func GetDeckID(commander string)int{
    var ID int
    for i := 0; i < len(models.Users); i++ {
		for j := 0; j < len(models.Users[i].Decks); j++ {
			if models.Users[i].Decks[j].Commander == commander{
				ID = models.Users[i].Decks
			}
		}
	}
    return ID
}

func AddDeckForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Parsea los datos del formulario

	var p models.Deck

	commander := r.FormValue("AddCommander")
	powerlvl := r.FormValue("AddPowerlvl")
	colors := r.FormValue("AddColors")
	user_id := r.FormValue("AddUser_id")

	p.Commander, p.Powerlvl, p.Colors, p.User_id = commander, ConvertirSTR(powerlvl), colors, ConvertirSTR(user_id)

	// Convertir el objeto Usuario a JSON
    jsonBody, err := json.Marshal(p)
    if err != nil {
        fmt.Println("Error al convertir el objeto a JSON:", err)
        return
    }

	// URL de tu API y cuerpo de la solicitud POST
    apiUrl := "http://localhost:8081/decks"
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

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func RemoveDeckForm(w http.ResponseWriter, r *http.Request) {

	commander := strings.Split(r.URL.Path, "/")[1]

    DeleteDeckPage(strconv.Itoa(GetDeckID(commander)))

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func UpdateDeckForm(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    commanderform := r.FormValue("UpdateCommander")
	powerlvl := r.FormValue("UpdatePowerlvl")
	colors := r.FormValue("UpdateColors")
	var deck models.Deck

	commanderURL := strings.Split(r.URL.Path, "/")[1]

	for i := 0; i < len(models.Users); i++ {
		for j := 0; j < len(models.Users[i].Decks); j++ {
			if models.Users[i].Decks[j].Commander == commanderURL{
				deck = models.Users[i].Decks[j]
			}
		}
	}
    deck.Commander = commanderform
    deck.Powerlvl= ConvertirSTR(powerlvl)
    deck.Colors=colors
    PutDeckPage(deck)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func GetDecksFromUser(ID string)[]models.Deck{
	// URL de la API externa
    url := "http://localhost:8081/decks/"+ID

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
    var p []models.Deck
    err = json.NewDecoder(resp.Body).Decode(&p)
    if err != nil {
        fmt.Println("Error al decodificar la respuesta JSON:", err)
    }
	return p
}

func GetAllDecks()[]models.Deck{
	// URL de la API externa
    url := "http://localhost:8081/users/decks"

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
    var p []models.Deck
    err = json.NewDecoder(resp.Body).Decode(&p)
    if err != nil {
        fmt.Println("Error al decodificar la respuesta JSON:", err)
    }
	return p
}

func DeleteDeckPage(ID string){
	// URL de tu API para eliminar un usuario específico
    apiUrl := "http://localhost:8081/decks/" + ID

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

func PutDeckPage(p models.Deck){
	// Convertir el objeto Usuario a JSON
    jsonBody, err := json.Marshal(p)
    if err != nil {
        fmt.Println("Error al convertir el objeto a JSON:", err)
        return
    }

    // URL de tu API para actualizar un usuario específico
    apiUrl := "http://localhost:8081/decks/" + strconv.Itoa(p.ID)

    // Crear una solicitud HTTP PUT con el cuerpo JSON
    req, err := http.NewRequest("PUT", apiUrl, bytes.NewBuffer(jsonBody))
    if err != nil {
        fmt.Println("Error al crear la solicitud PUT:", err)
        return
    }

    // Establecer el encabezado de tipo de contenido JSON
    req.Header.Set("Content-Type", "application/json")

    // Realizar la solicitud PUT
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error al realizar la solicitud PUT:", err)
        return
    }
    defer resp.Body.Close()

    // Verificar el código de estado de la respuesta
    if resp.StatusCode != http.StatusOK {
        fmt.Println("La solicitud PUT devolvió un código de estado no válido:", resp.Status)
        return
    }
}

