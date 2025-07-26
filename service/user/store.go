package user

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/pavleRakic/testGoApi/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM QuizUser WHERE email = @email", sql.Named("email", email))

	if err != nil {
		return nil, err
	}

	u := new(types.User)

	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.IDUser == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.IDUser,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.IsAdult,
		&user.IDRole,
		&user.CurrentStreak,
		&user.HighestStreak,
		&user.CreatorPoints,
		&user.QuizzerPoints,
		&user.TranslatorPoints,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	/*rows, err := s.db.Query("SELECT * FROM QuizUser WHERE idUser = @idUser", sql.Named("idUser", id))
	if err != nil {
		return nil, err
	}

	u := new(types.User)

	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.IDUser == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil*/

	row := s.db.QueryRow("SELECT * FROM QuizUser WHERE idUser = @idUser", sql.Named("idUser", id))

	u := new(types.User)
	err := row.Scan(&u.IDUser, &u.Username, &u.Password, &u.Email, &u.IsAdult, &u.IDRole, &u.CurrentStreak, &u.HighestStreak, &u.QuizzerPoints, &u.CreatorPoints, &u.TranslatorPoints)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return u, nil
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO QuizUser"+
		"(idUser,username,userPassword,email,isAdult,idRole,currentStreak,highestStreak, quizzerPoints, creatorPoints, translatorPoints)"+
		"VALUES(@idUser, @username, @userPassword, @email, @isAdult, @idRole, @currentStreak, @highestStreak, @quizzerPoints, @creatorPoints, @translatorPoints)",
		sql.Named("idUser", user.IDUser),
		sql.Named("username", user.Username),
		sql.Named("userPassword", user.Password),
		sql.Named("email", user.Email),
		sql.Named("isAdult", user.IsAdult),
		sql.Named("idRole", user.IDRole),
		sql.Named("currentStreak", user.CurrentStreak),
		sql.Named("highestStreak", user.HighestStreak),
		sql.Named("quizzerPoints", user.QuizzerPoints),
		sql.Named("creatorPoints", user.CreatorPoints),
		sql.Named("translatorPoints", user.TranslatorPoints),
	)

	if err != nil {
		return err
	}

	return nil
}
