package handler

import (
	"context"
	"encoding/json"
	"github.com/balloon-chat/topic-recommend/internal/domain/service"
	"log"
	"net/http"
)

type UpdateRecommendTopicsResponse struct {
	Pickups []string `json:"pickups"`
}

func UpdateRecommendTopics(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := context.Background()
	if topicService == nil {
		s, err := service.NewTopicService(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		topicService = s
	}

	recommend, err := topicService.SaveRecommendTopics()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	res := UpdateRecommendTopicsResponse{
		Pickups: recommend.Pickup,
	}

	if err = json.NewEncoder(w).Encode(&res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
