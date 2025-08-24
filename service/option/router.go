package option

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/pavleRakic/testGoApi/types"
	"github.com/pavleRakic/testGoApi/utils"
)

type Handler struct {
	store types.OptionController
}

func NewHandler(store types.OptionController) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/question/{idQuestion}/options", h.handleCreateOption).Methods(http.MethodPost)
	router.HandleFunc("/question/{idQuestion}/options", h.handleGetOptions).Methods(http.MethodGet)
	router.HandleFunc("/question/{idQuestion}/options", h.handleDeleteOptions).Methods(http.MethodDelete)
	router.HandleFunc("/question/option/{idOption}", h.handleUpdateOption).Methods(http.MethodPut)
}

func (h *Handler) handleGetOptions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err2 := strconv.Atoi(vars["idQuestion"])

	if err2 != nil {
		utils.WriteError(w, http.StatusBadRequest, err2)
		return
	}
	quiz, err := h.store.GetOption(id)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, quiz)
}

func (h *Handler) handleCreateOption(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateOptionPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.CreateOption(types.Option{
		IDQuestion:  payload.IDQuestion,
		OptionText:  payload.OptionText,
		OptionImage: payload.OptionImage,
		OptionOrder: payload.OptionOrder,
		IsCorrect:   payload.IsCorrect,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleDeleteOptions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err2 := strconv.Atoi(vars["idQuestion"])

	if err2 != nil {
		utils.WriteError(w, http.StatusBadRequest, err2)
		return
	}
	var payload types.DeleteOptionsPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.DeleteOption(id, payload.IDOptions)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleUpdateOption(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idOption, err2 := strconv.Atoi(vars["idOption"])

	if err2 != nil {
		utils.WriteError(w, http.StatusBadRequest, err2)
		return
	}

	var payload types.CreateOptionPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.UpdateOption(types.Option{
		IDOption:    idOption,
		OptionText:  payload.OptionText,
		OptionImage: payload.OptionImage,
		OptionOrder: payload.OptionOrder,
		IsCorrect:   payload.IsCorrect,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}
