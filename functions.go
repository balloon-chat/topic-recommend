package functions

import (
	"context"
	"encoding/json"
	service2 "github.com/balloon-chat/topic-recommend/internal/domain/service"
	"log"
	"net/http"
)

func RecommendTopics(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	s, err := service2.NewTopicService(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
	}

	recommend, err := s.SaveRecommendTopics()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
	}

	if err = json.NewEncoder(w).Encode(recommend); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
	}
	w.WriteHeader(http.StatusOK)
}
