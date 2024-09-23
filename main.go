package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var message string
var DB *gorm.DB

type requestBody struct {
	Message string `json:"message"`
}

func InitDB() {
	// в dsn вводим данные, которые мы указали при создании контейнера
	dsn := "host=localhost user=postgres password=postgres dbname=2task port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
}

func SetMessage(w http.ResponseWriter, r *http.Request) {
	var req requestBody
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message = req.Message
	fmt.Fprintf(w, "message: %s", message)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %s", message)
}

func main() {
	// Вызываем метод InitDB() из файла db.go
	InitDB()

	// Автоматическая миграция модели Message
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	http.ListenAndServe(":8080", router)
}
