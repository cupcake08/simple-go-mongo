package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cupcake08/golang-mongo/controllers"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial(envMongoURL())
	if err != nil {
		fmt.Println("Failed to connect to DB")
		log.Fatal(err)
	}
	fmt.Print("Connection established")
	return s
}

func envMongoURL() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGOURL")
}
