package product

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/pavleRakic/testGoApi/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]types.Quiz, error) {

	rows, err := s.db.Query("SELECT * FROM WebProject.Quiz")
	if err != nil {
		return nil, err
	}

	products := make([]types.Quiz, 0)
	for rows.Next() {
		p, err := scanRowsIntoQuiz(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil

}

func (s *Store) GetAllQuestions(idQuiz int) ([]types.Question, error) {

	rows, err := s.db.Query("SELECT * FROM WebProject.Question WHERE idQuiz=@idQuiz", sql.Named("idQuiz", idQuiz))
	if err != nil {
		return nil, err
	}

	questions := make([]types.Question, 0)
	for rows.Next() {
		p, err := scanRowsIntoQuestion(rows)
		if err != nil {
			return nil, err
		}

		questions = append(questions, *p)
	}

	return questions, nil

}
func (s *Store) GetOption(idOption int) ([]types.Option, error) {

	rows, err := s.db.Query("SELECT * FROM WebProject.QuestionOption WHERE idQuestion=@idQuestion", sql.Named("idQuestion", idOption))
	if err != nil {
		return nil, err
	}

	options := make([]types.Option, 0)
	for rows.Next() {
		p, err := scanRowsIntoOption(rows)
		if err != nil {
			return nil, err
		}

		options = append(options, *p)
	}

	return options, nil

}

func (s *Store) GetQuizByID(idQuiz int) (*types.Quiz, error) {

	row := s.db.QueryRow("SELECT * FROM WebProject.Quiz WHERE idQuiz = @idQuiz", sql.Named("idQuiz", idQuiz))
	log.Println("Hey hey hey")
	r := new(types.Quiz)
	err := row.Scan(&r.IDQuiz, &r.QuizName, &r.Description, &r.QuizImageLocation, &r.CreationDate, &r.HasTimer, &r.Timer, &r.HasLifeline, &r.IDType, &r.IDCreator, &r.IDCategory, &r.IDLanguage, &r.AvgRating, &r.IsNSFW, &r.UniquePlays, &r.Plays)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("role not found")
		}
		return nil, err
	}

	return r, nil

}

func scanRowsIntoOption(rows *sql.Rows) (*types.Option, error) {
	options := new(types.Option)

	err := rows.Scan(
		&options.IDOption,
		&options.IDQuestion,
		&options.OptionText,
		&options.OptionImage,
		&options.OptionOrder,
		&options.IsCorrect,
	)

	if err != nil {
		return nil, err
	}

	return options, nil
}

func scanRowsIntoQuestion(rows *sql.Rows) (*types.Question, error) {
	question := new(types.Question)

	err := rows.Scan(
		&question.IDQuiz,
		&question.IDQuestion,
		&question.QuestionImage,
		&question.QuestionText,
		&question.QuestionOrder,
		&question.IsMultiChoice,
		&question.Timer,
	)

	if err != nil {
		return nil, err
	}

	return question, nil
}

func scanRowsIntoQuiz(rows *sql.Rows) (*types.Quiz, error) {
	product := new(types.Quiz)

	err := rows.Scan(
		&product.IDQuiz,
		&product.QuizName,
		&product.Description,
		&product.QuizImageLocation,
		&product.CreationDate,
		&product.HasTimer,
		&product.Timer,
		&product.HasLifeline,
		&product.IDType,
		&product.IDCreator,
		&product.IDCategory,
		&product.IDLanguage,
		&product.AvgRating,
		&product.IsNSFW,
		&product.UniquePlays,
		&product.Plays,
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Store) CreateQuiz(quiz types.Quiz) error {

	_, err := s.db.Exec("INSERT INTO WebProject.Quiz "+
		"(quizName, description,quizImageLocation, creationDate, hasTimer, timer, hasLifeline, idType, idCreator, idCategory, idLanguage, avgRating, isNSFW, uniquePlays, plays)"+
		"VALUES(@quizName, @description, @quizImageLocation, @creationDate, @hasTimer, @timer, @hasLifeline, @idType, @idCreator, @idCategory, @idLanguage, @avgRating, @isNSFW, @uniquePlays, @plays)",
		sql.Named("quizName", quiz.QuizName),
		sql.Named("description", quiz.Description),
		sql.Named("quizImageLocation", quiz.QuizImageLocation),
		sql.Named("creationDate", quiz.CreationDate),
		sql.Named("hasTimer", quiz.HasTimer),
		sql.Named("timer", quiz.Timer),
		sql.Named("hasLifeline", quiz.HasLifeline),
		sql.Named("idType", quiz.IDType),
		sql.Named("idCreator", quiz.IDCreator),
		sql.Named("idCategory", quiz.IDCategory),
		sql.Named("idLanguage", quiz.IDLanguage),
		sql.Named("avgRating", quiz.AvgRating),
		sql.Named("isNSFW", quiz.IsNSFW),
		sql.Named("uniquePlays", quiz.UniquePlays),
		sql.Named("plays", quiz.Plays),
	)

	if err != nil {
		return err
	}

	return nil
}
