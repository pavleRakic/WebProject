package types

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

/*
type mockUserStore struct {

}
func GetUserByEmail(email string) (*User, error) {
	return nil, nil
}*/

/*type RegisterUserPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}*/

type RegisterUserPayload struct {
	IDUser           int    `json:"idUser"`
	Username         string `json:"username" validate:"required"`
	Password         string `json:"userPassword" validate:"required,min=3,max=130"`
	Email            string `json:"email" validate:"required,email"`
	IsAdult          bool   `json:"isAdult" validate:"required"`
	IDRole           int    `json:"idRole" validate:"required"`
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
	IDRole           int    `json:"idRole"`
	CurrentStreak    int    `json:"currentStreak"`
	HighestStreak    int    `json:"highestStreak"`
	QuizzerPoints    int    `json:"quizzerPoints"`
	CreatorPoints    int    `json:"creatorPoints"`
	TranslatorPoints int    `json:"translatorPoints"`
}
