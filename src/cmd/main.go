package main

import (
	"context"
	"github.com/balloon-chat/topic-recommend/src/service"
	"log"
	"strings"
)

func main() {
	ctx := context.Background()
	s, err := service.NewTopicService(ctx)
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
