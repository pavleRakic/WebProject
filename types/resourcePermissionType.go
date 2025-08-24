package types

type ResourcePermissionController interface {
	GetResourcePermissionByID(id int) (*ResourcePermission, error)
	GetResourcePermissions() ([]ResourcePermission, error)
	DeleteResourcePermission(id int) error
	UpdateResourcePermission(ResourcePermission) error
	CreateResourcePermission(ResourcePermission) error
}

type ResourcePermission struct {
	IDResourcePermission int `json:"idResourcePermission"`
	IDResource           int `json:"idResource"`
	IDPermission         int `json:"idPermission"`
}

type CreateResourcePermissionPayload struct {
	IDResource   int `json:"idResource"`
	IDPermission int `json:"idPermission"`
}
