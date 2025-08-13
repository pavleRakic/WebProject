package types

type PermissionController interface {
	GetPermissionByID(id int) (*Permission, error)
	GetPermissions() ([]Permission, error)
	DeletePermission(id int) error
	UpdatePermission(Permission) error
	CreatePermission(Permission) error
}

type Permission struct {
	IDPermission   int    `json:"idPermission"`
	PermissionName string `json:"permissionName"`
}

type CreatePermissionPayload struct {
	PermissionName string `json:"permissionName"`
}
