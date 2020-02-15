package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/madjlzz/madlens/controllers"
	"github.com/madjlzz/madlens/middleware"
	"github.com/madjlzz/madlens/models"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "madlens_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	services, err := models.NewServices(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer services.Close()

	//us.DestructiveReset()
	services.AutoMigrate()

	r := mux.NewRouter()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(services.User)
	galleriesC := controllers.NewGalleries(services.Gallery, r)

	requireUserMiddleware := middleware.RequireUser{
		UserService: services.User,
	}

	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")

	// Gallery routes
	r.Handle("/galleries/new", requireUserMiddleware.Apply(galleriesC.New)).Methods("GET")
	r.HandleFunc("/galleries", requireUserMiddleware.ApplyFn(galleriesC.Create)).Methods("POST")
	r.HandleFunc("/galleries/{id:[0-9]+}", galleriesC.Show).Methods("GET").Name(controllers.ShowGalleryName)
	r.HandleFunc("/galleries/{id:[0-9]+}/edit", requireUserMiddleware.ApplyFn(galleriesC.Edit)).Methods("GET")

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
