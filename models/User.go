package models

import (
	database "github.com/ricardoassaad/GoBackend/database"
)

type User struct {
	Id   int
	Name string
}

func All() []User {
	db := database.InitiateDatabaseConnection()
	res, err := db.Query("SELECT * FROM users ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	user := User{}
	var users []User
	for res.Next() {
		var id int
		var name string
		err = res.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		user.Id = id
		user.Name = name
		users = append(users, user)
	}
	return users
}

func ShowUser(userId string) User {
	db := database.InitiateDatabaseConnection()
	res, err := db.Query("SELECT * FROM users WHERE id = ?", userId)
	if err != nil {
		panic(err.Error())
	}

	user := User{}
	for res.Next() {
		var id int
		var name string
		err = res.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}

		user.Id = id
		user.Name = name
	}

	return user
}
