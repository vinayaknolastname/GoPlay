package main

// import (
// 	"log"
// 	"net"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func main() {

// 	listener, err := net.Listen("tcp", "localhost:8000")

// 	if err != nil {
// 		log.Println("sdddd", err)
// 	}

// 	router := mux.NewRouter()

// 	router.HandleFunc("/lol", lol).Methods("GET")

// 	err = http.Serve(listener, router)

// }

// func lol(r http.ResponseWriter, j *http.Request) {

// 	print("sddd")
// }
