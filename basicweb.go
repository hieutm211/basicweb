
package main

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/hieutm211/basicweb/login"
	"github.com/hieutm211/basicweb/register"
	"github.com/hieutm211/basicweb/home"
)

func idx_Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "github.com/hieutm211/basicweb/index.html")
}

func idx_Login_Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "github.com/hieutm211/basicweb/login/index.html")
}

func idx_Register_Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "github.com/hieutm211/basicweb/register/index.html")
}

func main(){
	router := mux.NewRouter()

	router.Handle("/", http.HandlerFunc(idx_Handler))

	router.Handle("/login", http.HandlerFunc(idx_Login_Handler))
	router.Handle("/login/check", http.HandlerFunc(login.Handler))

	router.Handle("/register", http.HandlerFunc(idx_Register_Handler))
	router.Handle("/register/check", http.HandlerFunc(register.Handler))

	router.Handle("/home", http.HandlerFunc(home.Handler))

	server := &http.Server {
		Handler: router,
		Addr: ":8080",
		ReadTimeout: 15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	fmt.Println("Server is listening on port 8080")
	log.Fatal(server.ListenAndServe())
}
