package main

import (
	"context"
	service2 "github.com/balloon-chat/topic-recommend/internal/domain/service"
	"log"
	"strings"
)

func main() {
	ctx := context.Background()
	s, err := service2.NewTopicService(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	pickups, err := s.GetPickupTopics()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Pickup Topics ", strings.Repeat("=", 10))
	for _, p := range pickups {
		log.Println(*p)
	}

	newest, err := s.GetNewestTopics()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Newest Topics ", strings.Repeat("=", 10))
	for _, n := range newest {
		log.Println(*n)
	}
}
