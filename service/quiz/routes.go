package product

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	router.HandleFunc("/quiz/{idQuiz}", h.handleGetQuizByID).Methods(http.MethodGet)
	router.HandleFunc("/getFullQuiz/{idQuiz}", h.handleGetFullQuizByID).Methods(http.MethodGet)
}

func (h *Handler) handleGetQuizByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idQuiz, err2 := strconv.Atoi(vars["idQuiz"])

	if err2 != nil {
		utils.WriteError(w, http.StatusBadRequest, err2)
		return
	}

	quiz, err := h.store.GetQuizByID(idQuiz)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, quiz)
}

func (h *Handler) handleGetFullQuizByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idQuiz, err2 := strconv.Atoi(vars["idQuiz"])

	if err2 != nil {
		utils.WriteError(w, http.StatusBadRequest, err2)
		return
	}

	quiz, err := h.store.GetQuizByID(idQuiz)
	questions, err2 := h.store.GetAllQuestions(idQuiz)

	var fullQuestions []types.QuestionFull

	for _, q := range questions {
		log.Println(q.IDQuestion)
		options, err := h.store.GetOption(q.IDQuestion)

		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		fullQuestions = append(fullQuestions, types.QuestionFull{
			IDQuestion:    q.IDQuestion,
			QuestionText:  q.QuestionText,
			QuestionImage: q.QuestionImage,
			QuestionOrder: q.QuestionOrder,
			IDQuiz:        q.IDQuiz,
			IsMultiChoice: q.IsMultiChoice,
			Timer:         q.Timer,
			Options:       options,
		})

	}

	var fullQuiz types.QuizFullPayload
	fullQuiz.IDQuiz = quiz.IDQuiz
	fullQuiz.QuizName = quiz.QuizName
	fullQuiz.Description = quiz.Description
	fullQuiz.QuizImageLocation = quiz.QuizImageLocation
	fullQuiz.CreationDate = quiz.CreationDate
	fullQuiz.HasTimer = quiz.HasTimer
	fullQuiz.Timer = quiz.Timer
	fullQuiz.HasLifeline = quiz.HasLifeline
	fullQuiz.IDType = quiz.IDType
	fullQuiz.IDCreator = quiz.IDCreator
	fullQuiz.IDCategory = quiz.IDCategory
	fullQuiz.IDLanguage = quiz.IDLanguage
	fullQuiz.AvgRating = quiz.AvgRating
	fullQuiz.Plays = quiz.Plays
	fullQuiz.UniquePlays = quiz.UniquePlays
	fullQuiz.IsNSFW = quiz.IsNSFW

	fullQuiz.Questions = fullQuestions

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err2 != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, fullQuiz)
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
		QuizName:          payload.QuizName,
		Description:       payload.Description,
		QuizImageLocation: payload.QuizImageLocation,
		CreationDate:      payload.CreationDate,
		HasTimer:          payload.HasTimer,
		Timer:             payload.Timer,
		HasLifeline:       payload.HasLifeline,
		IDType:            payload.IDType,
		IDCreator:         payload.IDCreator,
		IDCategory:        payload.IDCategory,
		IDLanguage:        payload.IDLanguage,
		AvgRating:         payload.AvgRating,
		IsNSFW:            payload.IsNSFW,
		UniquePlays:       payload.UniquePlays,
		Plays:             payload.Plays,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
