package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/madjlzz/madlens/controllers"
	"github.com/madjlzz/madlens/models"
	"net/http"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "password"
	dbname = "madlens_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()

	us.AutoMigrate()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(us)

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
