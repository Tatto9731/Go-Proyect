package views

import (
	"BEModule/controllers"
	"FEModule/models"
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
	var user models.User
	user.ID = controllers.ConvertSTR(ID)
	err := templates.ExecuteTemplate(w, "adddeck.html", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RemovePage(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("views/templates/remove.html"))
	ID := strings.Split(r.URL.Path, "/")[1]
	var user models.User
	user.ID = controllers.ConvertSTR(ID)
	err := templates.ExecuteTemplate(w, "remove.html", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RemoveDeckPage(w http.ResponseWriter, r *http.Request) {
	var deck models.Deck
	var templates = template.Must(template.ParseFiles("views/templates/removedeck.html"))
	deck.Commander = strings.Split(r.URL.Path, "/")[1]
	err := templates.ExecuteTemplate(w, "removedeck.html", deck)
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
	var user models.User
	user.ID = controllers.ConvertSTR(ID)
	err := templates.ExecuteTemplate(w, "update.html",user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UpdateDeckPage(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("views/templates/updatedeck.html"))
	commander := strings.Split(r.URL.Path, "/")[1]
	var deck models.Deck
	deck.Commander = commander
	err := templates.ExecuteTemplate(w, "updatedeck.html",deck)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CardDeckPage(w http.ResponseWriter, r *http.Request) {

	var templates = template.Must(template.ParseFiles("views/templates/card.html"))
	var deck models.Deck
	deck.Commander = strings.Split(r.URL.Path, "/")[1]
	deck.Cards = controllers.GetCardFromAPI(controllers.SearchCard(w,r))
	err := templates.ExecuteTemplate(w, "card.html", deck)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}