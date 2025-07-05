package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Status:  201,
		Message: "Successfull",
		Data:    map[string]int{"h": 1},
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func main() {
	fmt.Println("server is running in 8080")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
