package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arvindkhoisnam/challenges/04/db"
	"github.com/arvindkhoisnam/challenges/04/models"
	"github.com/arvindkhoisnam/challenges/04/routehandlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	app := gin.Default()
	godotenv.Load()
	config := db.DbConfig{
		Host: os.Getenv("DB_HOST"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName: os.Getenv("DB_NAME"),
		Port: os.Getenv("DB_PORT"),
		SslMode: os.Getenv("DB_SSLMODE"),
	}
	client,err := db.GenerateClient(&config)
	if err != nil {
		fmt.Println(err)
	}
	log.Println("DB CLIENT GENERATED")
	if err := models.MigrateDB(client); err != nil{
		fmt.Println(err)
	}

	log.Println("DB MIGRATED")
	repo := routehandlers.Repository{
		DbClient: client,
	}
	gracefulServer := http.Server{
		Addr: ":8080",
		Handler: app,
	}
	repo.Routes(app)

	quit := make(chan os.Signal,1)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)

	go func ()  {
		if err := gracefulServer.ListenAndServe(); err != nil && err != http.ErrServerClosed{
			log.Fatal(err)
		}
	}()
	<- quit
	fmt.Println("Shutting down")

	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	if err := gracefulServer.Shutdown(ctx);err != nil {
		log.Fatal("A problem occured while shutting down.")
	}
}