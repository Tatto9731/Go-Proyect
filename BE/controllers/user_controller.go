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

func AddForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Parsea los datos del formulario

	var p models.User

	name := r.FormValue("AddName")
	age := r.FormValue("AddAge")
	email := r.FormValue("AddEmail")

	p.Name, p.Age, p.Email = name, ConvertirSTR(age), email

	// Convertir el objeto Usuario a JSON
    jsonBody, err := json.Marshal(p)
    if err != nil {
        fmt.Println("Error al convertir el objeto a JSON:", err)
        return
    }

	// URL de tu API y cuerpo de la solicitud POST
    apiUrl := "http://localhost:8081/users"
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

func RemoveForm(w http.ResponseWriter, r *http.Request) {

    DeleteUserPage(strings.Split(r.URL.Path, "/")[1])

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func GetIDForm(w http.ResponseWriter, r *http.Request)string {
	r.ParseForm() // Parsea los datos del formulario

	ID := r.FormValue("ID")

	return ID
}

func UpdateForm(w http.ResponseWriter, r *http.Request) {
	var p models.User

	r.ParseForm() // Parsea los datos del formulario

	ID := strings.Split(r.URL.Path, "/")[1]
	name := r.FormValue("UpdateName")
	age := r.FormValue("UpdateAge")
	email := r.FormValue("UpdateEmail")

	p.ID, p.Name, p.Age, p.Email = ConvertirSTR(ID), name, ConvertirSTR(age), email
	
	PutUserPage(p)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func GetUsersPage()[]models.User{
	// URL de la API externa
    url := "http://localhost:8081/users"

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
    var p []models.User
    err = json.NewDecoder(resp.Body).Decode(&p)
    if err != nil {
        fmt.Println("Error al decodificar la respuesta JSON:", err)
    }
	return p
}

func DeleteUserPage(ID string){
	// URL de tu API para eliminar un usuario específico
    apiUrl := "http://localhost:8081/users/" + ID

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

func PutUserPage(p models.User){
	// Convertir el objeto Usuario a JSON
    jsonBody, err := json.Marshal(p)
    if err != nil {
        fmt.Println("Error al convertir el objeto a JSON:", err)
        return
    }

    // URL de tu API para actualizar un usuario específico
    apiUrl := "http://localhost:8081/users/" + strconv.Itoa(p.ID)

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
