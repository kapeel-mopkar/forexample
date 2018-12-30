package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var mySigningKey = []byte("supersecretphrase")
func GenerateJWT() (string, error){
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Kapeel Mopkar"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if(err != nil){
		fmt.Errorf("Somethingwent wrong! :s", err.Error())
		return "", err
	}
	return tokenString, nil
}

type Article struct{
	Title string
	Desc string
	Content string
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request){
	articles := Articles{
		Article{
			Title: "Test Title",
			Desc: "Test Description",
			Content: "Hello World",
		},
		Article{
			Title: "Test Title 2",
			Desc: "Test Description 2",
			Content: "Hello World 2",
		},
	}
	fmt.Println("Endpoint hit: All Articles endpoint")
	json.NewEncoder(w).Encode(articles)
}

func testPostArticles(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Test Post Article endpoint hit")
}

func homePage(w http.ResponseWriter, r *http.Request){
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, validToken)
}

func handleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8888", myRouter))
}

func main(){
	fmt.Println("Simple Client")
	handleRequests()

	/*tokenString, err := GenerateJWT()
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println(tokenString)*/
}