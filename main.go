package main

import (
	"encoding/json"
	"fmt"
	greetingService "go-greeting-app/greeting"
	"net/http"
)

func main() {
	http.HandleFunc("/greeting", func(writer http.ResponseWriter, request *http.Request) {
		greeting := greetingService.GetGreeting()
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(greeting)
	})

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8080/greeting'")
	http.ListenAndServe(":8080", nil)
}
