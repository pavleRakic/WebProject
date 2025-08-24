package types

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type RegisterUserPayload struct {
	IDUser           int    `json:"idUser"`
	Username         string `json:"username" validate:"required"`
	Password         string `json:"userPassword" validate:"required,min=3,max=130"`
	Email            string `json:"email" validate:"required,email"`
	IsAdult          bool   `json:"isAdult" validate:"required"`
	CurrentStreak    int    `json:"currentStreak" validate:"required"`
	HighestStreak    int    `json:"highestStreak" validate:"required"`
	QuizzerPoints    int    `json:"quizzerPoints" validate:"required"`
	CreatorPoints    int    `json:"creatorPoints" validate:"required"`
	TranslatorPoints int    `json:"translatorPoints" validate:"required"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"userPassword" validate:"required"`
}

type User struct {
	IDUser           int    `json:"idUser"`
	Username         string `json:"username"`
	Password         string `json:"userPassword"`
	Email            string `json:"email"`
	IsAdult          bool   `json:"isAdult"`
	CurrentStreak    int    `json:"currentStreak"`
	HighestStreak    int    `json:"highestStreak"`
	QuizzerPoints    int    `json:"quizzerPoints"`
	CreatorPoints    int    `json:"creatorPoints"`
	TranslatorPoints int    `json:"translatorPoints"`
}

type ProductStore interface {
	GetProducts() ([]Quiz, error)
	CreateQuiz(Quiz) error
	GetQuizByID(id int) (*Quiz, error)
	GetAllQuestions(id int) ([]Question, error)
	GetOption(id int) ([]Option, error)
}

type Quiz struct {
	IDQuiz            int     `json:"idQuiz"`
	QuizName          string  `json:"quizName"`
	Description       string  `json:"description"`
	QuizImageLocation string  `json:"quizImageLocation"`
	CreationDate      string  `json:"creationDate"`
	HasTimer          bool    `json:"hasTimer"`
	Timer             int     `json:"timer"`
	HasLifeline       bool    `json:"hasLifeline"`
	IDType            int     `json:"idType"`
	IDCreator         int     `json:"idCreator"`
	IDCategory        int     `json:"idCategory"`
	IDLanguage        int     `json:"idLanguage"`
	AvgRating         float32 `json:"avgRating"`
	IsNSFW            bool    `json:"isNSFW"`
	UniquePlays       int     `json:"uniquePlays"`
	Plays             int     `json:"plays"`
}

type QuizFullPayload struct {
	IDQuiz            int            `json:"idQuiz"`
	QuizName          string         `json:"quizName"`
	Description       string         `json:"description"`
	QuizImageLocation string         `json:"quizImageLocation"`
	CreationDate      string         `json:"creationDate"`
	HasTimer          bool           `json:"hasTimer"`
	Timer             int            `json:"timer"`
	HasLifeline       bool           `json:"hasLifeline"`
	IDType            int            `json:"idType"`
	IDCreator         int            `json:"idCreator"`
	IDCategory        int            `json:"idCategory"`
	IDLanguage        int            `json:"idLanguage"`
	AvgRating         float32        `json:"avgRating"`
	IsNSFW            bool           `json:"isNSFW"`
	UniquePlays       int            `json:"uniquePlays"`
	Plays             int            `json:"plays"`
	Questions         []QuestionFull `json"questions"`
}

type CreateQuizPayload struct {
	QuizName          string  `json:"quizName"`
	Description       string  `json:"description"`
	QuizImageLocation string  `json:"quizImageLocation"`
	CreationDate      string  `json:"creationDate"`
	HasTimer          bool    `json:"hasTimer"`
	Timer             int     `json:"timer"`
	HasLifeline       bool    `json:"hasLifeline"`
	IDType            int     `json:"idType"`
	IDCreator         int     `json:"idCreator"`
	IDCategory        int     `json:"idCategory"`
	IDLanguage        int     `json:"idLanguage"`
	AvgRating         float32 `json:"avgRating"`
	IsNSFW            bool    `json:"isNSFW"`
	UniquePlays       int     `json:"uniquePlays"`
	Plays             int     `json:"plays"`
}
