package main

import (
	"fmt"
	"log"
	"net/http"

	"vulnerable-app/config"
	"vulnerable-app/internal/db"
	"vulnerable-app/internal/handlers"
)

func main() {
	// Доступ к уязвимому паролю/ключу захардкоженному в конфиге
	fmt.Printf("Loaded AWS Key: %s\n", config.AWSAccessKey)

	// Инициализация базы данных
	db.InitDB()

	// Регистрация уязвимых роутов
	http.HandleFunc("/user", handlers.GetUser)
	http.HandleFunc("/ping", handlers.PingHandler)
	http.HandleFunc("/download", handlers.DownloadHandler)

	fmt.Println("Intentionally Vulnerable Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
