package main

import (
	"app/config"
	"app/controller"
	"app/models"
	"app/storage/jsondb"
	"fmt"
	"log"
)

func main() {

	cfg := config.Load()

	strg, err := jsondb.NewConnectionJSON(&cfg)
	if err != nil {
		panic("Failed connect to json:" + err.Error())
	}

	con := controller.NewController(&cfg, strg)

	resp, err := con.UserCreate(&models.CreateUser{
		FirstName: "Asadbek",
		LastName:  "Ergashev",
	})

	if err != nil {
		log.Println("user create err:", err)
		return
	}

	fmt.Println(resp)
}
