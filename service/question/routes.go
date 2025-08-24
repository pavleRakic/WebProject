package question

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/pavleRakic/testGoApi/service/auth"
	"github.com/pavleRakic/testGoApi/types"
	"github.com/pavleRakic/testGoApi/utils"
)

type Handler struct {
	store types.QuestionController
}

func NewHandler(store types.QuestionController) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/quiz/{idQuiz}/questions", auth.JWTMiddleware2(h.handleCreateQuestion)).Methods(http.MethodPost)
	router.HandleFunc("/quiz/{idQuiz}/questions", h.handleGetQuestions).Methods(http.MethodGet)
	router.HandleFunc("/quiz/{idQuiz}/questions", h.handleDeleteQuestions).Methods(http.MethodDelete)
	router.HandleFunc("/quiz/question/{idQuestion}", h.handleUpdateQuestion).Methods(http.MethodPut)
}

func (h *Handler) handleGetQuestions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err2 := strconv.Atoi(vars["idQuiz"])

	if err2 != nil {
		utils.WriteError(w, http.StatusBadRequest, err2)
		return
	}
	quiz, err := h.store.GetQuestions(id)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, quiz)
}

func (h *Handler) handleCreateQuestion(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateQuestionPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.CreateQuestion(types.Question{
		IDQuiz:        payload.IDQuiz,
		QuestionText:  payload.QuestionText,
		QuestionImage: payload.QuestionImage,
		QuestionOrder: payload.QuestionOrder,
		IsMultiChoice: payload.IsMultiChoice,
		Timer:         payload.Timer,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleDeleteQuestions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err2 := strconv.Atoi(vars["idQuiz"])

	if err2 != nil {
		utils.WriteError(w, http.StatusBadRequest, err2)
		return
	}
	var payload types.DeleteQuestionsPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.DeleteQuestions(id, payload.IDQuestion)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleUpdateQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idQuestion, err2 := strconv.Atoi(vars["idQuestion"])

	if err2 != nil {
		utils.WriteError(w, http.StatusBadRequest, err2)
		return
	}

	var payload types.CreateQuestionPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.UpdateQuestion(types.Question{
		IDQuestion:    idQuestion,
		QuestionText:  payload.QuestionText,
		QuestionImage: payload.QuestionImage,
		QuestionOrder: payload.QuestionOrder,
		IsMultiChoice: payload.IsMultiChoice,
		Timer:         payload.Timer,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}
