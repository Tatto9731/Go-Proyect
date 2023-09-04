package models

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
	Decks []Deck `json:"decks"`
}

var Users []User

func AddUsers(p User) {
	Users = append(Users, p)
}

func GetUSers() []User {
	return Users
}

func RemoveUser(id int) {
	for i := 0; i <= len(Users)-1; i++ {
		if Users[i].ID == id {
			Users = append(Users[:i], Users[i+1:]...)
		}
	}
}

func UpdateUser(p int, nombrep string, age int, email string) {
	for i := 0; i <= len(Users)-1; i++ {
		if Users[i].ID == p {
			Users[i].Name = nombrep
			Users[i].Age = age
			Users[i].Email = email
		}
	}
}

func FindUser(id int) User {
	var p User
	for i := 0; i <= len(Users)-1; i++ {
		if Users[i].ID == id {
			p = Users[i]
		}
	}
	return p
}

func DeleteUsers() {
	Users = []User{}
}