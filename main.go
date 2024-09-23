package main

import (
	"encoding/json"
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
	dsn := "host=localhost user=postgres password=postgres dbname=2task port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var req requestBody
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message = req.Message
	newMessage := Message{Text: message}
	res := DB.Create(&Message{Text: message})

	if err := res.Error; err != nil {
		http.Error(w, "Failed to create message", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newMessage)
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	messages := make([]Message, 1)

	res := DB.Find(&messages)
	if err := res.Error; err != nil {
		log.Fatal("[DB] Ошибка поиска ", err)
		return
	}

	response := map[string][]Message{
		"messages": messages,
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	http.ListenAndServe(":8080", router)
}
