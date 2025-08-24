package types

type OptionController interface {
	GetOptionByID(id int) (*Option, error)
	GetAllOption() ([]Option, error)
	GetOption(id int) ([]Option, error)
	DeleteOption(id int, option []int) error
	UpdateOption(Option) error
	CreateOption(Option) error
}

type Option struct {
	IDOption    int    `json:"idOption"`
	IDQuestion  int    `json:"idQuestion"`
	OptionText  string `json:"optionText"`
	OptionImage string `json:"optionImage"`
	OptionOrder int    `json:"optionOrder"`
	IsCorrect   bool   `json:"isCorrect"`
}

type CreateOptionPayload struct {
	IDQuestion  int    `json:"idQuestion"`
	OptionText  string `json:"optionText"`
	OptionImage string `json:"optionImage"`
	OptionOrder int    `json:"optionOrder"`
	IsCorrect   bool   `json:"isCorrect"`
}

type DeleteOptionsPayload struct {
	IDOptions []int `json:"idOptions"`
}
