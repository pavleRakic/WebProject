package types

type RoleController interface {
	GetRoleByID(id int) (*Role, error)
	GetRoles() ([]Role, error)
	DeleteRole(id int) error
	UpdateRole(Role) error
	CreateRole(Role) error
}

type Role struct {
	IDRole   int    `json:"idRole"`
	RoleName string `json:"roleName"`
}

type CreateRolePayload struct {
	RoleName string `json:"roleName"`
}
