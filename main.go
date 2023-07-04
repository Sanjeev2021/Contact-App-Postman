package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"contactapp/controller"
)

func main() {
	// Called HandleFunction (to run program)
	HandleFunction()
}

func HandleFunction() {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	router := mux.NewRouter().StrictSlash(true)

	// Made routes (users)
	router.HandleFunc("/createuser", controller.CreateUser).Methods("POST")
	router.HandleFunc("/getuser/{id}", controller.GetUserById).Methods("GET")
	router.HandleFunc("/getusers", controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/updateuser/{id}", controller.UpdateUserById).Methods("POST")
	router.HandleFunc("/deleteuser/{id}", controller.DeleteUserById).Methods("POST")

	router.HandleFunc("/createcontact/{id}", controller.CreateContact).Methods("POST")
	router.HandleFunc("/getcontact/{id}", controller.GetContactById).Methods("GET")
	router.HandleFunc("/getcontacts", controller.GetAllContacts).Methods("GET")
	router.HandleFunc("/updatecontact/{id}", controller.UpdateContact).Methods("POST")
	router.HandleFunc("/deletecontact/{id}", controller.DeleteContact).Methods("POST")

	log.Printf("Server Live on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
