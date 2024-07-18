package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type Router struct {
	UserHandler         *UserHandler
	ReviewerHandler     *ReviewerHandler
	AlignmentHandler    *AlignmentHandler
	QuestionnaireHandler *QuestionnaireHandler
	RecommendationHandler *RecommendationHandler
}

func NewRouter(db *gorm.DB) *Router {
	return &Router{
		UserHandler:         NewUserHandler(db),
		ReviewerHandler:     NewReviewerHandler(db),
		AlignmentHandler:    NewAlignmentHandler(db),
		QuestionnaireHandler: NewQuestionnaireHandler(db),
		RecommendationHandler: NewRecommendationHandler(db),
	}
}

func (r *Router) AddRoutes(router *mux.Router) {
	r.UserHandler.AddRoutes(router)
	r.ReviewerHandler.AddRoutes(router)
	r.AlignmentHandler.AddRoutes(router)
	r.QuestionnaireHandler.AddRoutes(router)
	r.RecommendationHandler.AddRoutes(router)
}

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=gorm dbname=gorm password=gorm sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	r := NewRouter(db)
	r.AddRoutes(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
