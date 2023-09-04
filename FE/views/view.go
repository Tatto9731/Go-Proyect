package views

import (
	"BEModule/controllers"
	"FEModule/models"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	models.Users = controllers.GetUsersPage()
	tmpl := template.Must(template.ParseFiles("views/templates/main.html"))
		err := tmpl.ExecuteTemplate(w,"main.html", models.Users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		
}

func DeckPage(w http.ResponseWriter, r *http.Request) {
	var deck models.Deck
	deck.Commander = strings.Split(r.URL.Path, "/")[1]
	deck.ID =  controllers.GetDeckID(deck.Commander)
	deck.Cards = controllers.GetCardsFromDeck(strconv.Itoa(controllers.GetDeckID(deck.Commander)))
	fmt.Println(deck)
	tmpl := template.Must(template.ParseFiles("views/templates/deck.html"))
		err := tmpl.ExecuteTemplate(w,"deck.html", deck)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
}

func AddPage(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("views/templates/add.html"))
	err := templates.ExecuteTemplate(w, "add.html", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AddDeckPage(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("views/templates/adddeck.html"))
	ID := strings.Split(r.URL.Path, "/")[1]
	var p models.User
	p.ID = controllers.ConvertirSTR(ID)
	err := templates.ExecuteTemplate(w, "adddeck.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RemovePage(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("views/templates/remove.html"))
	ID := strings.Split(r.URL.Path, "/")[1]
	var p models.User
	p.ID = controllers.ConvertirSTR(ID)
	err := templates.ExecuteTemplate(w, "remove.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RemoveDeckPage(w http.ResponseWriter, r *http.Request) {
	var p models.Deck
	var templates = template.Must(template.ParseFiles("views/templates/removedeck.html"))
	p.Commander = strings.Split(r.URL.Path, "/")[1]
	err := templates.ExecuteTemplate(w, "removedeck.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RemoveCardPage(w http.ResponseWriter, r *http.Request) {
	var deck models.Deck
	var card models.Card
	var templates = template.Must(template.ParseFiles("views/templates/removecard.html"))
	deck.Commander = strings.Split(r.URL.Path, "/")[1]
	card.Name = strings.Split(r.URL.Path, "/")[4]
	deck.Cards = append(deck.Cards, card)
	err := templates.ExecuteTemplate(w, "removecard.html", deck)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UpdatePage(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("views/templates/update.html"))
	ID := strings.Split(r.URL.Path, "/")[1]
	var p models.User
	p.ID = controllers.ConvertirSTR(ID)
	err := templates.ExecuteTemplate(w, "update.html",p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UpdateDeckPage(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("views/templates/updatedeck.html"))
	commander := strings.Split(r.URL.Path, "/")[1]
	var p models.Deck
	p.Commander = commander
	err := templates.ExecuteTemplate(w, "updatedeck.html",p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CardDeckPage(w http.ResponseWriter, r *http.Request) {

	var templates = template.Must(template.ParseFiles("views/templates/card.html"))
	var p models.Deck
	p.Commander = strings.Split(r.URL.Path, "/")[1]
	p.Cards = controllers.GetCardAPI(controllers.SearchCard(w,r))
	err := templates.ExecuteTemplate(w, "card.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}