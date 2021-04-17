package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

// User: Structure to store users info
type User struct{
	Username string
	Password string
	Token string
}
// Global variable user. All functions are able to access to it
var user User

func main() {
	// User and password registered
	user.Username="PIPO"
	user.Password="123"
	router := mux.NewRouter()
	//All routes for API
	//Need functions
	router.HandleFunc("/Login",login)
	router.HandleFunc("/Logout",logout)
	router.HandleFunc("/Status",status)
	router.HandleFunc("/Upload",addImage)
	router.HandleFunc("/Token",token)
	http.ListenAndServe(":5000",router)
}
func login(w http.ResponseWriter, r *http.Request){

}
func logout(w http.ResponseWriter, r *http.Request){

}
func status(w http.ResponseWriter, r *http.Request){

}
func addImage(w http.ResponseWriter, r *http.Request){

}
func token(){

}