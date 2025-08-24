package user_role

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
	store types.UserRoleController
}

func NewHandler(store types.UserRoleController) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/user/{idUser}/roles", h.handleGetUserRoles).Methods("GET")
	router.HandleFunc("/user/{idUser}/role", h.handleAssignUserRoles).Methods("POST")
	router.HandleFunc("/user/{idUser}/roles", h.handleUnassignUserRoles).Methods("DELETE")
	router.HandleFunc("/users/roles", h.handleGetUsersRoles).Methods("GET")
	router.HandleFunc("/user/{idUser}/role/{idRole}", h.handleGetUserRole).Methods("GET")
}

func (h *Handler) handleGetUserRoles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idUser, err2 := strconv.Atoi(vars["idUser"])

	if err2 != nil {
		utils.WriteError(w, http.StatusBadRequest, err2)
		return
	}

	userRoles, err := h.store.GetUserRoles(idUser)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, userRoles)
}

func (h *Handler) handleAssignUserRoles(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idUser, err2 := strconv.Atoi(vars["idUser"])

	if err2 != nil {
		utils.WriteError(w, http.StatusBadRequest, err2)
		return
	}

	var payload types.AsignRolesPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.AssignUserRoles(idUser, payload.IDRole)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleUnassignUserRoles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err2 := strconv.Atoi(vars["idUser"])

	if err2 != nil {
		utils.WriteError(w, http.StatusBadRequest, err2)
		return
	}
	var payload types.AsignRolesPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.UnassignUserRoles(id, payload.IDRole)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleGetUsersRoles(w http.ResponseWriter, r *http.Request) {
	resourcePermission, err := h.store.GetUsersRoles()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, resourcePermission)
}

func (h *Handler) handleGetUserRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idUser, err2 := strconv.Atoi(vars["idUser"])
	idRole, err3 := strconv.Atoi(vars["idRole"])

	if err2 != nil || err3 != nil {
		utils.WriteError(w, http.StatusBadRequest, err2)
		return
	}
	resourcePermission, err := h.store.GetUserRoleByID(idUser, idRole)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, resourcePermission)
}
