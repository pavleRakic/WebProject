package types

type RoleResourcePermissionController interface {
	GetRoleResourcePermissionByID(id int) (*RoleResourcePermission, error)
	GetRoleResourcePermissions() ([]RoleResourcePermission, error)
	DeleteRoleResourcePermission(id int) error
	UpdateRoleResourcePermission(RoleResourcePermission) error
	CreateRoleResourcePermission(RoleResourcePermission) error
}

type RoleResourcePermission struct {
	IDRoleResourcePermission int `json:"idRoleResourcePermission"`
	IDRole                   int `json:"idRole"`
	IDResourcePermission     int `json:"idResourcePermission"`
}

type CreateRoleResourcePermissionPayload struct {
	IDRole               int `json:"idRole"`
	IDResourcePermission int `json:"idResourcePermission"`
}
