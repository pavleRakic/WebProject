package types

type UserRoleController interface {
	GetUserRoleByID(idUser int, idRole int) (*UserRole, error)
	GetUserRoles(idUser int) ([]UserRole, error)
	GetUsersRoles() ([]UserRole, error)
	UnassignUserRoles(idUser int, roles []string) error
	AssignUserRoles(idUser int, roles []string) error
}

type UserRole struct {
	IDUser int `json:"idUser"`
	IDRole int `json:"idRole"`
}

type CreateUserRolePayload struct {
	IDUser int `json:"idUser"`
	IDRole int `json:"idRole"`
}

type AsignRolesPayload struct {
	IDRole []string `json:"idRole"`
}
