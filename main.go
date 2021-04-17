package main

import (
	"fmt"
	"crypto/rand"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// User: Structure to store users info
type User struct {
	UID 	 string
	Pass	 string
	Token    string
}

// Global variable user. All functions are able to access to it

//Hola Hector

var user User

func main() {
	// User and password registered
	user.UID = "PIPO"
	user.Pass = "123"
	router := mux.NewRouter()
	//All routes for API
	//Need functions
	router.HandleFunc("/Login", login)
	router.HandleFunc("/Logout", logout)
	router.HandleFunc("/Status", status)
	router.HandleFunc("/Upload", addImage)
	router.HandleFunc("/Token", token)
	http.ListenAndServe(":5000", router)
}
<<<<<<< HEAD
func login() {

}
func logout() {

}
func status() {

}
func addImage() {
=======
func login(w http.ResponseWriter, r *http.Request){
	var userW User
	if user.UID == userW.UID && user.Pass == userW.Pass{
		user.Token=token()
	}else{
		fmt.Println("Wrong Data")
	}
}
func logout(w http.ResponseWriter, r *http.Request){
	user.Token=""
	fmt.Println("Logged out")
}
func status(w http.ResponseWriter, r *http.Request){

}
func addImage(w http.ResponseWriter, r *http.Request){
>>>>>>> 18b11f2c8643e6420f665e468ac7a2523f33ef1e

}
func token() {
	token:=make([]byte, 4)
	return rand.Read(token)
}
