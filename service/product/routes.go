package product

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/pavleRakic/testGoApi/service/auth"
	"github.com/pavleRakic/testGoApi/types"
	"github.com/pavleRakic/testGoApi/utils"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/quizs", auth.JWTMiddleware2(h.handleCreateQuiz)).Methods(http.MethodPost)
	router.HandleFunc("/quizs", h.handleGetQuiz).Methods(http.MethodGet)

}

func (h *Handler) handleGetQuiz(w http.ResponseWriter, r *http.Request) {
	quiz, err := h.store.GetProducts()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, quiz)
}

func (h *Handler) handleCreateQuiz(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateQuizPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.CreateQuiz(types.Quiz{
		IDQuiz:       payload.IDQuiz,
		QuizName:     payload.QuizName,
		Description:  payload.Description,
		CreationDate: payload.CreationDate,
		HasTimer:     payload.HasTimer,
		Timer:        payload.Timer,
		HasLifeline:  payload.HasLifeline,
		IDType:       payload.IDType,
		IDCreator:    payload.IDCreator,
		IDCategory:   payload.IDCategory,
		IDLanguage:   payload.IDLanguage,
		AvgRating:    payload.AvgRating,
		IsNSFW:       payload.IsNSFW,
		UniquePlays:  payload.UniquePlays,
		Plays:        payload.Plays,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
