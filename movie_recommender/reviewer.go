package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Review struct {
	gorm.Model
	ReviewerID uint
	MovieID    uint
	Content    string `gorm:"type:text"`
}

type Reviewer struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100);unique_index"`
	Reviews []Review
}

type ReviewerHandler struct {
	DB *gorm.DB
}

func NewReviewerHandler(db *gorm.DB) *ReviewerHandler {
	return &ReviewerHandler{
		DB: db,
	}
}

func (rh *ReviewerHandler) GetReviews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	reviewer := &Reviewer{}
	if err := rh.DB.Preload("Reviews").Where("id = ?", id).First(reviewer).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// In a real-world application, you would serialize this data into JSON or XML.
	// For simplicity, we're just writing the reviews directly to the response.
	for _, review := range reviewer.Reviews {
		w.Write([]byte(review.Content))
	}
}

func (rh *ReviewerHandler) AddRoutes(router *mux.Router) {
	router.HandleFunc("/reviewer/{id}/reviews", rh.GetReviews).Methods("GET")
}
