package main

import (
	"fmt"
	"praktikum/config"
	"praktikum/routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	config.InitDB()
	e := routes.New()
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
	fmt.Println("test")
}
