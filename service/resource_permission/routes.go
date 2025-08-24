package resource_permission

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
	store types.ResourcePermissionController
}

func NewHandler(store types.ResourcePermissionController) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/resourcePermissions", h.handleGetResourcePermissions).Methods("GET")
	router.HandleFunc("/resourcePermission", h.handleCreateResourcePermission).Methods("POST")
	router.HandleFunc("/resourcePermission/{idResourcePermission}", h.handleDeleteResourcePermission).Methods("DELETE")
	router.HandleFunc("/resourcePermission/{idResourcePermission}", h.handleUpdateResourcePermission).Methods("PUT")

}

func (h *Handler) handleGetResourcePermissions(w http.ResponseWriter, r *http.Request) {
	resourcePermission, err := h.store.GetResourcePermissions()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, resourcePermission)
}

func (h *Handler) handleCreateResourcePermission(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateResourcePermissionPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.CreateResourcePermission(types.ResourcePermission{
		IDResource:   payload.IDResource,
		IDPermission: payload.IDPermission,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleDeleteResourcePermission(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err2 := strconv.Atoi(vars["idResourcePermission"])

	if err2 != nil {
		utils.WriteError(w, http.StatusBadRequest, err2)
		return
	}

	err := h.store.DeleteResourcePermission(id)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleUpdateResourcePermission(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err2 := strconv.Atoi(vars["idResourcePermission"])

	if err2 != nil {
		utils.WriteError(w, http.StatusBadRequest, err2)
		return
	}

	var payload types.CreateResourcePermissionPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.UpdateResourcePermission(types.ResourcePermission{
		IDResourcePermission: id,
		IDResource:           payload.IDResource,
		IDPermission:         payload.IDPermission,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}
