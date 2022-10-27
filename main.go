package main

import (
	"fmt"
	"log"
	"net/http"
	c "project/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	http.Handle("/", router)
	router.HandleFunc("/students", c.GetAllStudents).Methods("GET")
	router.HandleFunc("/students/store", c.InsertStudent).Methods("POST")
	router.HandleFunc("/student/{id}", c.GetStudent).Methods("GET")
	router.HandleFunc("/student/{id}/update", c.UpdateStudent).Methods("PUT")
	router.HandleFunc("/student/{id}/delete", c.DeleteStudent).Methods("DELETE")
	fmt.Println("Server berjalan")
	log.Fatal(http.ListenAndServe(":8080", router))
}
