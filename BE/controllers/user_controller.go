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

//Gets the user information from the view into a form
func AddUserForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() 

	var user models.User

	name := r.FormValue("AddName")
	age := r.FormValue("AddAge")
	email := r.FormValue("AddEmail")

	user.Name, user.Age, user.Email = name, ConvertSTR(age), email

	// Convert the object into a json body
    jsonBody, err := json.Marshal(user)
    if err != nil {
        fmt.Println("Error trying to decode into a JSON:", err)
        return
    }

    apiUrl := "http://localhost:8081/users"
    requestBody := bytes.NewBuffer(jsonBody)

	// Post request
    resp, err := http.Post(apiUrl, "application/json", requestBody)
    if err != nil {
        fmt.Println("Error trying to request a POST:", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Println("The request returns an invalid state:", resp.Status)
        return
    }

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

//Gets the user info from the URL in order to delete the user
func RemoveUser(w http.ResponseWriter, r *http.Request) {

    DeleteUserPage(strings.Split(r.URL.Path, "/")[1])

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

//Gets the user info from the view using a form and the URL
func UpdateForm(w http.ResponseWriter, r *http.Request) {
	var user models.User

	r.ParseForm()

	id := strings.Split(r.URL.Path, "/")[1]
	name := r.FormValue("UpdateName")
	age := r.FormValue("UpdateAge")
	email := r.FormValue("UpdateEmail")

	user.ID, user.Name, user.Age, user.Email = ConvertSTR(id), name, ConvertSTR(age), email
	
	PutUserPage(user)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func GetUsersPage()[]models.User{
    url := "http://localhost:8081/users"

    // Get request
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error trying to complete the request:", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Println("The request returns an invalid state:", resp.Status)
    }

    var users []models.User
    err = json.NewDecoder(resp.Body).Decode(&users)
    if err != nil {
        fmt.Println("Error trying to decode into a JSON:", err)
    }
	return users
}

//Deletes the user requesting a DELETE method to my API
func DeleteUserPage(ID string){

    apiUrl := "http://localhost:8081/users/" + ID

    // Creates a DELETE request
    req, err := http.NewRequest("DELETE", apiUrl, nil)
    if err != nil {
        fmt.Println("Error creating the DELETE request:", err)
        return
    }

    client := &http.Client{}
    resp, _ := client.Do(req)
    defer resp.Body.Close()
}

//Creates a request to update the user info to the API
func PutUserPage(p models.User){
    //Convert the object into a json body
    jsonBody, err := json.Marshal(p)
    if err != nil {
        fmt.Println("Error trying to decode into a JSON:", err)
        return
    }

    apiUrl := "http://localhost:8081/users/" + strconv.Itoa(p.ID)

    // Creates a PUT request
    req, err := http.NewRequest("PUT", apiUrl, bytes.NewBuffer(jsonBody))
    if err != nil {
        fmt.Println("Error creating a PUT request:", err)
        return
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error creating a PUT request:", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Println("The PUT request returns an invalid state:", resp.Status)
        return
    }
}
