package question

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/pavleRakic/testGoApi/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetQuestions(idQuiz int) ([]types.Question, error) {

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

func (s *Store) GetAllQuestions() ([]types.Question, error) {

	rows, err := s.db.Query("SELECT * FROM WebProject.Question")
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

func (s *Store) GetQuestionByID(idQuestion int) (*types.Question, error) {

	row := s.db.QueryRow("SELECT * FROM WebProject.Question WHERE idQuestion=@idQuestion", sql.Named("idQuestion", idQuestion))

	r := new(types.Question)
	err := row.Scan(&r.IDQuestion, &r.IDQuiz, &r.QuestionText, &r.QuestionOrder, &r.IsMultiChoice, &r.Timer)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("resource permission not found")
		}
		return nil, err
	}

	return r, nil

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

func (s *Store) CreateQuestion(question types.Question) error {

	_, err := s.db.Exec("INSERT INTO WebProject.Question "+
		"(idQuiz, questionText, questionImage, questionOrder, isMultiChoice, timer)"+
		"VALUES(@idQuiz, @questionText, @questionOrder, @isMultiChoice, @timer)",
		sql.Named("idQuiz", question.IDQuiz),
		sql.Named("questionText", question.QuestionText),
		sql.Named("questionImage", question.QuestionImage),
		sql.Named("questionOrder", question.QuestionOrder),
		sql.Named("isMultiChoice", question.IsMultiChoice),
		sql.Named("timer", question.Timer),
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteQuestions(idQuiz int, questions []int) error {
	if len(questions) == 0 {
		return nil
	}

	placeholders := make([]string, len(questions))
	args := []interface{}{idQuiz} // first arg is user ID

	for i, roleID := range questions {
		placeholders[i] = fmt.Sprintf("@p%d", i+2) // @p2, @p3, ...
		args = append(args, roleID)
	}

	query := fmt.Sprintf(
		"DELETE FROM WebProject.Question WHERE idQuiz = @p1 AND idQuestion IN (%s)",
		strings.Join(placeholders, ","),
	)

	_, err := s.db.Exec(query, args...)
	return err
}

func (s *Store) UpdateQuestion(question types.Question) error {
	_, err := s.db.Exec(

		"UPDATE WebProject.Question SET "+
			"questionText=@questionText, questionImage, questionOrder=@questionOrder, isMultiChoice=@isMultiChoice, timer=@timer"+
			" WHERE idQuestion=@idQuestion",
		sql.Named("questionText", question.QuestionText),
		sql.Named("questionImage", question.QuestionImage),
		sql.Named("questionOrder", question.QuestionOrder),
		sql.Named("isMultiChoice", question.IsMultiChoice),
		sql.Named("timer", question.Timer),
		sql.Named("idQuestion", question.IDQuestion),
	)

	if err != nil {
		return err
	}

	return nil
}
