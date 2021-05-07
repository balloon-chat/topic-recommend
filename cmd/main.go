package main

import (
	"github.com/balloon-chat/topic-recommend/internal/interface/api/server/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.UpdateRecommendTopics)
	log.Println("Listening at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return
	}
}
