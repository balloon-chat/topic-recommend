package functions

import (
	"github.com/balloon-chat/topic-recommend/internal/interface/api/server/handler"
	"net/http"
)

func RecommendTopics(w http.ResponseWriter, r *http.Request) {
	handler.UpdateRecommendTopics(w, r)
}
