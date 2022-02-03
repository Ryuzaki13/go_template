package db

import "fmt"

type User struct {
	Login    string
	Password string
	Name     string
}

func (u *User) Insert() {
	_, e := Link.Exec(`
	INSERT INTO "User"
    ("Login", "Password", "Name")
    VALUES ($1, $2, $3)`, u.Login, u.Password, u.Name)
	if e != nil {
		fmt.Println(e)
	}
}

func (u *User) Select() {
	row := Link.QueryRow(`SELECT "Name" FROM "User" 
WHERE "Login"=$1 AND "Password"=$2`,
		u.Login, u.Password)
	e := row.Scan(&u.Name)
	if e != nil {
		fmt.Println(e)
	}
}

func (u *User) SelectAll() []User {
	rows, e := Link.Query(`SELECT "Login", "Name" FROM "User" ORDER BY "Name"`)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		e = rows.Scan(&u.Login, &u.Name)
		if e != nil {
			fmt.Println(e)
			return nil
		}

		users = append(users, User{
			Login:    u.Login,
			Password: "",
			Name:     u.Name,
		})
	}

	return users
}
