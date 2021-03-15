package functions

import (
	"context"
	"encoding/json"
	"github.com/balloon-chat/topic-recommend/src/service"
	"log"
	"net/http"
)

func RecommendTopics(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	s, err := service.NewTopicService(ctx)
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
