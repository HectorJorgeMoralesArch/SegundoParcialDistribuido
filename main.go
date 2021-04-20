package main

import (
	"math/rand"
	"net/http"
	"strings"
	"time"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"github.com/gorilla/mux"
	"log"
)

// Logged: Structure to store logged in users info
var online = make(map[string]string)

var users = make(map[string]string)

Type Image struct{
	Token,Name string
	x,y int
}
// Global variable user. All functions are able to access to it
func main() {
	router := mux.NewRouter()
	users["username"] = "password"
	//All routes for API
	//Need functions
	router.HandleFunc("/login", login)
	router.HandleFunc("/logout", logout)
	router.HandleFunc("/status", status)
	router.HandleFunc("/upload", addImage)
	http.ListenAndServe(":8080", router)
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	usr, pwd, ok := r.BasicAuth()
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		msg := `
	{
		"Username and Password are required"
	}
`
		w.Write([]byte(msg))
		return
	}
	if !checkPass(usr, pwd) {
		w.WriteHeader(http.StatusUnauthorized)
		msg := `
	{
		"Invalid Username and/or Password"
	}
`
		w.Write([]byte(msg))
		return
	}
	w.WriteHeader(http.StatusOK)
	token := GenerateRandomString(10)
	msg := `
	{
		"message": "Hi ` + usr + ` welcome to the DPIP System"
		"token" "` + token + `"
	}
`
	w.Write([]byte(msg))
	online[token] = usr
	return
}

func checkPass(usr string, pwd string) bool {
	if p, ok := users[usr]; ok {
		return p == pwd
	}
	return false
}

func logout(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if len(token) < 7 {
		w.WriteHeader(http.StatusUnauthorized)
		msg := `
	{
		"Please enter a token" ` + token + `smiles
	}
`
		w.Write([]byte(msg))
		return
	}
	token = token[7:]
	if !loggedIn(token) {
		w.WriteHeader(http.StatusUnauthorized)
		msg := `
	{
		"Please enter a valid token"
	}
`
		w.Write([]byte(msg))
		return
	}
	w.WriteHeader(http.StatusOK)
	msg := `
	{
		"message": "Bye ` + online[token] + `, your token has been revoked"
	}
`
	w.Write([]byte(msg))
	delete(online, token)
	return
}

func loggedIn(token string) bool {
	if _, ok := online[token]; ok {
		return true
	}
	return false
}

func status(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if len(token) < 7 {
		w.WriteHeader(http.StatusUnauthorized)
		msg := `
	{
		"Please enter a token" ` + token + `
	}
`
		w.Write([]byte(msg))
		return
	}
	token = token[7:]
	if !loggedIn(token) {
		w.WriteHeader(http.StatusUnauthorized)
		msg := `
	{
		"Please enter a valid token"
	}
`
		w.Write([]byte(msg))
		return
	}
	w.WriteHeader(http.StatusOK)
	msg := `
	{
		"message": "Hi ` + online[token] + `, the DPIP System is Up and Running "
		"time": "` + time.Now().Format("2006-01-02 15:04:05") + `"
	}
`
	w.Write([]byte(msg))
	return
}
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateRandomString generate a string of random characters of given length
func GenerateRandomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		idx := rand.Int63() % int64(len(letterBytes))
		sb.WriteByte(letterBytes[idx])
	}
	return sb.String()
}














// Receive image and return name and size of it
func addImage(w http.ResponseWriter, r *http.Request) {
	var img Image
	msg:=""
	json.NewDecoder(r.Body).Decode(&img)
	if img.Token == users.Token {
		msg=`
		{
			"Message": "An image has been successfully uploaded",
			"Filename": "`+img.Name+`",
			"Size": "`+getImageSize(img.Name)+`",
			"Width": "`+getImageDimension(img.Name)[0]+`",
			"Height": "`+getImageDimension(img.Name)[1]+`",
		}
		`
	}else{
		msg := `
		{
			"Please enter a valid token"
		}
		`
	}
	w.Write([]byte(msg))
	return
}
// Retrieve image size
func getImageSize(imgPath string) (int64){
	fi, err := os.Stat(imgPath);
	if err != nil {
		log.Fatal(err)
	}
	// Get the size
	size := fi.Size()
	return size
}

// Get image dimensions
func getImageDimension(imgPath string) (int, int) {
	fi, err := os.Open(imgPath)
	defer fi.Close()
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.DecodeConfig(fi)
	if err != nil {
		log.Fatal(err)
	}
	return img.Width, img.Height
}