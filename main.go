package main

import (
	"log"
	"main/server"
	"main/server/db"
	"main/server/socket"
	"os"

	"github.com/joho/godotenv"
)

// @title Drag Racing
// @version 2.0
// @description This is the doumentation for drag racing game
// @BasePath /
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	connection := db.InitDB()
	db.Transfer(connection)
	socketServer := socket.SocketInit()
	defer socketServer.Close()
	app := server.NewServer(connection)
	server.ConfigureRoutes(app)

	if err := app.Run(os.Getenv("PORT")); err != nil {
		log.Print(err)
	}
}
