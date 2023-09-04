package models

type CardResponse struct {
	Cards []struct {
		Name string `json:"name"`
		Text string `json:"text"`
		Number string `json:"number"`
	} `json:"cards"`
}

type Card struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Deck_id int `json:"deck_id"`
	Text string `json:"text"`
	Number string `json:"number"`
}

var Cards []Card

func AddCards(p Card){
	Cards = append(Cards, p)
}

func RemoveCard(id int) {
	for i := 0; i <= len(Cards)-1; i++ {
		if Cards[i].ID == id {
			Cards = append(Cards[:i], Cards[i+1:]...)
		}
	}
}