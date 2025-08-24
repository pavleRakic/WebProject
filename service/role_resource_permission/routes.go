package role_resource_permission

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
	store types.RoleResourcePermissionController
}

func NewHandler(store types.RoleResourcePermissionController) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/roleResourcePermissions", h.handleGetRoleResourcePermissions).Methods("GET")
	router.HandleFunc("/roleResourcePermission", h.handleCreateRoleResourcePermission).Methods("POST")
	router.HandleFunc("/roleResourcePermission/{idRoleResourcePermission}", h.handleDeleteRoleResourcePermission).Methods("DELETE")
	router.HandleFunc("/roleResourcePermission/{idRoleResourcePermission}", h.handleUpdateRoleResourcePermission).Methods("PUT")

}

func (h *Handler) handleGetRoleResourcePermissions(w http.ResponseWriter, r *http.Request) {
	resourcePermission, err := h.store.GetRoleResourcePermissions()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, resourcePermission)
}

func (h *Handler) handleCreateRoleResourcePermission(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateRoleResourcePermissionPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.CreateRoleResourcePermission(types.RoleResourcePermission{
		IDRole:               payload.IDRole,
		IDResourcePermission: payload.IDResourcePermission,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleDeleteRoleResourcePermission(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err2 := strconv.Atoi(vars["idRoleResourcePermission"])

	if err2 != nil {
		utils.WriteError(w, http.StatusBadRequest, err2)
		return
	}

	err := h.store.DeleteRoleResourcePermission(id)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleUpdateRoleResourcePermission(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err2 := strconv.Atoi(vars["idRoleResourcePermission"])

	if err2 != nil {
		utils.WriteError(w, http.StatusBadRequest, err2)
		return
	}

	var payload types.CreateRoleResourcePermissionPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.UpdateRoleResourcePermission(types.RoleResourcePermission{
		IDRoleResourcePermission: id,
		IDRole:                   payload.IDRole,
		IDResourcePermission:     payload.IDResourcePermission,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}
