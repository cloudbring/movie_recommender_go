package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Questionnaire struct {
	gorm.Model
	UserID uint `gorm:"index"`
	Responses string `gorm:"type:text"`
}

type QuestionnaireHandler struct {
	DB *gorm.DB
}

func NewQuestionnaireHandler(db *gorm.DB) *QuestionnaireHandler {
	return &QuestionnaireHandler{
		DB: db,
	}
}

func (qh *QuestionnaireHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("userID")

	questionnaire := &Questionnaire{
		UserID: userID,
		Responses: "", // Initialize with empty responses
	}

	if err := qh.DB.Create(questionnaire).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (qh *QuestionnaireHandler) Submit(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("userID")
	responses := r.FormValue("responses")

	questionnaire := &Questionnaire{}
	if err := qh.DB.Where("userID = ?", userID).First(questionnaire).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	questionnaire.Responses = responses

	if err := qh.DB.Save(questionnaire).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (qh *QuestionnaireHandler) AddRoutes(router *mux.Router) {
	router.HandleFunc("/questionnaire/create", qh.Create).Methods("POST")
	router.HandleFunc("/questionnaire/submit", qh.Submit).Methods("POST")
}
