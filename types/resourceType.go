package types

type ResourceController interface {
	GetResourceByID(id int) (*Resource, error)
	GetResources() ([]Resource, error)
	DeleteResource(id int) error
	UpdateResource(Resource) error
	CreateResource(Resource) error
}

type Resource struct {
	IDResource   int    `json:"idResource"`
	ResourceName string `json:"resourceName"`
}

type CreateResourcePayload struct {
	ResourceName string `json:"resourceName"`
}
