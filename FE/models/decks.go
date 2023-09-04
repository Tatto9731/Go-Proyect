package models

type Deck struct{
	ID int `json:"id"`
	Commander string `json:"commander"`
	Powerlvl int `json:"powerlvl"`
	Colors string `json:"colors"`
	User_id int `json:"user_id"`
	Cards []Card `json:"cards"`
}

var Decks []Deck

func AddDecks(p Deck){
	Decks = append(Decks, p)
}

func GetDecks()[]Deck{
	return Decks
}

func RemoveDeck(id int) {
	for i := 0; i <= len(Decks)-1; i++ {
		if Decks[i].ID == id {
			Decks = append(Decks[:i], Decks[i+1:]...)
		}
	}
}

func UpdateDeck(id int, Commanderp string, Powerlvlp int, Colorsp string) {
	for i := 0; i <= len(Decks)-1; i++ {
		if Decks[i].ID == id {
			Decks[i].Commander = Commanderp
			Decks[i].Powerlvl = Powerlvlp
			Decks[i].Colors = Colorsp

		}
	}
}

func FindDeck(id int) Deck {
	var p Deck
	for i := 0; i <= len(Decks)-1; i++ {
		if Decks[i].ID == id {
			p = Decks[i]
		}
	}
	return p
}

func DeleteDecks(){
	Decks = [] Deck{}
}