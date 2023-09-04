package controllers

import (
	"FEModule/models"
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllUsersInfoSQL(){

	models.DeleteUsers()

	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/magicpage")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, age, email FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var p models.User
		err := rows.Scan(&p.ID, &p.Name, &p.Age, &p.Email)
		if err != nil {
			panic(err.Error())
		}
		db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/magicpage")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		rowsDecks, err := db.Query("SELECT id, commander, powerlvl, colors, user_id FROM decks WHERE user_id = ?", p.ID)
		if err != nil {
			panic(err.Error())
		}
		defer rowsDecks.Close()

		for rowsDecks.Next() {
			var t models.Deck
			err := rowsDecks.Scan(&t.ID, &t.Commander, &t.Powerlvl, &t.Colors, &t.User_id)
			if err != nil {
				panic(err.Error())
			}

			db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/magicpage")
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			rowsDecks, err := db.Query("SELECT id, name, deck_id, text FROM cards WHERE deck_id = ?", t.ID)
			if err != nil {
				panic(err.Error())
			}
			defer rowsDecks.Close()

			for rowsDecks.Next() {
				var j models.Card
				err := rowsDecks.Scan(&j.ID, &j.Name, &j.Deck_id, &j.Text)
				if err != nil {
					panic(err.Error())
				}
				t.Cards = append(t.Cards, j)
			}

			p.Decks = append(p.Decks, t)
		}

		models.AddUsers(p)
	}
}

func GetAllDecksSQL(){

	models.DeleteDecks()

	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/magicpage")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, commander, powerlvl, colors, user_id FROM decks")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Deck
		err := rows.Scan(&p.ID, &p.Commander, &p.Powerlvl, &p.Colors, &p.User_id)
		if err != nil {
			panic(err.Error())
		}
		models.AddDecks(p)
	}
}

func GetUserSQL(id int)models.User{

	var p models.User

	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/magicpage")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, age, email FROM users WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&p.ID, &p.Name, &p.Age, &p.Email)
		if err != nil {
			panic(err.Error())
		}
	}
	return p
}

func GetDecksSQL(id int)[]models.Deck{

	var ps []models.Deck

	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/magicpage")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, commander, powerlvl, colors, user_id FROM decks WHERE user_id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Deck
		err := rows.Scan(&p.ID, &p.Commander, &p.Powerlvl, &p.Colors, &p.User_id)
		if err != nil {
			panic(err.Error())
		}
		ps = append(ps, p)
	}
	return ps
}

func GetCardsSQL(id int)[]models.Card{

	var cards []models.Card

	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/magicpage")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, deck_id, text, number FROM cards WHERE deck_id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var card models.Card
		err := rows.Scan(&card.ID, &card.Name, &card.Deck_id, &card.Text, &card.Number)
		if err != nil {
			panic(err.Error())
		}
		cards = append(cards, card)
	}
	return cards
}

func AddUserSQL(name string, age int, email string)int{

	var p models.User

	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/magicpage")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO users (name, age, email) VALUES (?,?,?)", name,age,email)
	if err != nil {
		log.Fatal(err)
	}

	lastInsertID, _ := result.LastInsertId()

	p.ID = int(lastInsertID)
	p.Name = name
	p.Age = age
	p.Email = email

	models.AddUsers(p)
	return p.ID
}

func AddDeckSQL(Commander string, Powerlvl int, Colors string, User_id int)int{

	var p models.Deck

	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/magicpage")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO decks (commander, powerlvl, colors, user_id) VALUES (?,?,?,?)", Commander,Powerlvl,Colors,User_id)
	if err != nil {
		log.Fatal(err)
	}

	lastInsertID, _ := result.LastInsertId()

	p.ID = int(lastInsertID)
	p.Commander = Commander
	p.Powerlvl = Powerlvl
	p.Colors = Colors
	p.User_id = User_id

	models.AddDecks(p)
	return p.ID
}

func AddCardSQL(Name string, Deck_id int, Text string, Card_Number int)int{

	var p models.Card

	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/magicpage")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO cards (name, deck_id, text, number) VALUES (?,?,?,?)", Name,Deck_id,Text,Card_Number)
	if err != nil {
		log.Fatal(err)
	}

	lastInsertID, _ := result.LastInsertId()

	p.ID = int(lastInsertID)
	p.Name = Name
	p.Deck_id = Deck_id
	p.Text = Text
	p.Number = strconv.Itoa(Card_Number)

	models.AddCards(p)
	return p.ID
}

func UpdateUserSQL(p models.User){
	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/magicpage")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	updateQuery := "UPDATE users SET name = ?, age = ?, email = ? WHERE id = ?"

	_, err = db.Exec(updateQuery, p.Name, p.Age, p.Email, p.ID)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateDeckSQL(p models.Deck){
	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/magicpage")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	updateQuery := "UPDATE decks SET commander = ?, powerlvl = ?, colors = ? WHERE id = ?"

	_, err = db.Exec(updateQuery, p.Commander, p.Powerlvl, p.Colors, p.ID)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteUserSQL(id int){

	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/magicpage")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	deleteQuery := "DELETE FROM users WHERE id = ?"

	_, err = db.Exec(deleteQuery, id)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteDeckSQL(id int){

	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/magicpage")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	deleteQuery := "DELETE FROM decks WHERE id = ?"

	_, err = db.Exec(deleteQuery, id)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteCardSQL(id int){

	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/magicpage")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	deleteQuery := "DELETE FROM cards WHERE id = ?"

	_, err = db.Exec(deleteQuery, id)
	if err != nil {
		log.Fatal(err)
	}
}