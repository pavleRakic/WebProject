package role

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
	store types.RoleController
}

func NewHandler(store types.RoleController) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/roles", h.handleGetRoles).Methods("GET")
	router.HandleFunc("/role", h.handleCreateRole).Methods("POST")
	router.HandleFunc("/role/{idRole}", h.handleDeleteRole).Methods("DELETE")
	router.HandleFunc("/role/{idRole}", h.handleUpdateRole).Methods("PUT")

}

func (h *Handler) handleGetRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := h.store.GetRoles()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, roles)
}

func (h *Handler) handleCreateRole(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateRolePayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.CreateRole(types.Role{
		RoleName: payload.RoleName,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleDeleteRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err2 := strconv.Atoi(vars["idRole"])

	if err2 != nil {
		utils.WriteError(w, http.StatusBadRequest, err2)
		return
	}

	err := h.store.DeleteRole(id)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleUpdateRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err2 := strconv.Atoi(vars["idRole"])

	if err2 != nil {
		utils.WriteError(w, http.StatusBadRequest, err2)
		return
	}

	var payload types.CreateRolePayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.UpdateRole(types.Role{
		IDRole:   id,
		RoleName: payload.RoleName,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}
