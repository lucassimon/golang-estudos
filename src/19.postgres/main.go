package main

import (
	"fmt"

	_ "github.com/lib/pq"

	"github.com/lucassimon/golang-estudos/src/19.postgres/users"
)

func init() {

}

func main() {

	// fmt.Println("Criando usuários")
	// var usuarioFoo = users.User{Name: "Foo", Age: 11, Active: false}
	// users.Create(usuarioFoo)
	// var usuarioBar = users.User{Name: "Bar", Age: 21, Active: true}
	// users.Create(usuarioBar)

	listUsers, _ := users.All(10, 0)
	for index, user := range listUsers {
		fmt.Println("Index:", index, "user: ", user.ID, user.Name, user.Age, user.Active, user.CreatedAt, user.UpdatedAt, user.DeletedAt)

		user.Active = !user.Active
		users.Update(user)
	}

	listUsers, _ = users.All(10, 0)
	for index, user := range listUsers {
		fmt.Println("Index:", index, "user: ", user.ID, user.Name, user.Age, user.Active, user.CreatedAt, user.UpdatedAt, user.DeletedAt)
		users.Delete(user)
	}

	listUsers, _ = users.All(10, 0)
	for index, user := range listUsers {
		fmt.Println("Index:", index, "user: ", user.ID, user.Name, user.Age, user.Active, user.CreatedAt, user.UpdatedAt, user.DeletedAt)
	}

}
