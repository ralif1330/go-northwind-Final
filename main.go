package main

import (
	"log"
	"os"

	"code.id.northwind/config"
	"code.id.northwind/server"
	_ "github.com/lib/pq"

)

func main() {
	log.Println("Starting Northwind REST API!")

	log.Println("Initializing Configuration!")
	config := config.InitConfig(getConfigFileName())

	log.Println("Initializing DataBase!")
	dbHandler := server.InitDatabase(config)
	log.Println(dbHandler)

	log.Println("Initializing HTTP Server!")
	httpServer := server.InitHttpServer(config, dbHandler)

	httpServer.Start()

	// Test insert to category, using goroutine
	// ctx := context.Background()
	// queries := repositories.New(dbHandler)

	// newCategory, err := queries.CreateCategory(ctx,
	// 	repositories.CreateCategoryParams{
	// 		CategoryID:   101,
	// 		CategoryName: "Mainan",
	// 		Description:  "Mainan Anak",
	// 		Picture:      nil,
	// 	},
	// )

	// if err != nil {
	// 	log.Fatal("Error: ", err)
	// }

	// log.Println(newCategory)
}

func getConfigFileName() string {
	env := os.Getenv("ENV")

	if env != "" {
		return "northwind" + env
	}

	return "northwind"
}
