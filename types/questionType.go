package types

type QuestionController interface {
	GetQuestionByID(id int) (*Question, error)
	GetAllQuestions() ([]Question, error)
	GetQuestions(id int) ([]Question, error)
	DeleteQuestions(id int, questions []int) error
	UpdateQuestion(Question) error
	CreateQuestion(Question) error
}

type Question struct {
	IDQuestion    int    `json:"idQuestion"`
	IDQuiz        int    `json:"idQuiz"`
	QuestionImage string `json:"questionImage"`
	QuestionText  string `json:"questionText"`
	QuestionOrder int    `json:"questionOrder"`
	IsMultiChoice bool   `json:"isMultiChoice"`
	Timer         int    `json:"timer"`
}

type QuestionFull struct {
	IDQuestion    int      `json:"idQuestion"`
	IDQuiz        int      `json:"idQuiz"`
	QuestionImage string   `json:"questionImage"`
	QuestionText  string   `json:"questionText"`
	QuestionOrder int      `json:"questionOrder"`
	IsMultiChoice bool     `json:"isMultiChoice"`
	Timer         int      `json:"timer"`
	Options       []Option `json:"options"`
}

type CreateQuestionPayload struct {
	IDQuiz        int    `json:"idQuiz"`
	QuestionText  string `json:"questionText"`
	QuestionImage string `json:"questionImage"`
	QuestionOrder int    `json:"questionOrder"`
	IsMultiChoice bool   `json:"isMultiChoice"`
	Timer         int    `json:"timer"`
}

type DeleteQuestionsPayload struct {
	IDQuestion []int `json:"idQuestion"`
}
