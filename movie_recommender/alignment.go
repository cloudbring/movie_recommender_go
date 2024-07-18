package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Alignment struct {
	gorm.Model
	UserID     uint `gorm:"index"`
	ReviewerID uint `gorm:"index"`
	Score      float64
}

type AlignmentHandler struct {
	DB *gorm.DB
}

func NewAlignmentHandler(db *gorm.DB) *AlignmentHandler {
	return &AlignmentHandler{
		DB: db,
	}
}

func (ah *AlignmentHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	reviewerID := vars["reviewerID"]

	// In a real-world application, you would calculate the alignment score based on the user's questionnaire responses.
	// For simplicity, we're just setting a static score.
	score := 0.5

	alignment := &Alignment{
		UserID:     userID,
		ReviewerID: reviewerID,
		Score:      score,
	}

	if err := ah.DB.Save(alignment).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ah *AlignmentHandler) GetAlignment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	reviewerID := vars["reviewerID"]

	alignment := &Alignment{}
	if err := ah.DB.Where("userID = ? AND reviewerID = ?", userID, reviewerID).First(alignment).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// In a real-world application, you would serialize this data into JSON or XML.
	// For simplicity, we're just writing the score directly to the response.
	w.Write([]byte(alignment.Score))
}

func (ah *AlignmentHandler) AddRoutes(router *mux.Router) {
	router.HandleFunc("/alignment/{userID}/{reviewerID}/update", ah.Update).Methods("POST")
	router.HandleFunc("/alignment/{userID}/{reviewerID}", ah.GetAlignment).Methods("GET")
}
