package main

import (
	"context"
	"github.com/balloon-chat/topic-recommend/internal/domain/service"
	"log"
	"strings"
)

func main() {
	ctx := context.Background()
	s, err := service.NewTopicService(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// Pickupトピックを作成
	pickups, err := s.GetPickupTopics()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Pickup Topics ", strings.Repeat("=", 10))
	for _, p := range pickups {
		log.Println(*p)
	}
}
