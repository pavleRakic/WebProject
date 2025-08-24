package option

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

func (s *Store) GetAllOption() ([]types.Option, error) {

	rows, err := s.db.Query("SELECT * FROM WebProject.QuestionOption")
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

func (s *Store) GetOptionByID(idOption int) (*types.Option, error) {

	row := s.db.QueryRow("SELECT * FROM WebProject.QuestionOption WHERE idOption=@idOption", sql.Named("idOption", idOption))

	r := new(types.Option)
	err := row.Scan(&r.IDOption, &r.IDQuestion, &r.OptionText, &r.OptionImage, &r.OptionOrder, &r.IsCorrect)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("option was not found")
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

func (s *Store) CreateOption(options types.Option) error {

	_, err := s.db.Exec("INSERT INTO WebProject.QuestionOption "+
		"(idQuestion, optionText, optionImage, optionOrder, isCorrect)"+
		"VALUES(@idQuestion, @optionText, @optionImage, @optionOrder, @isCorrect)",
		sql.Named("idQuestion", options.IDQuestion),
		sql.Named("optionText", options.OptionText),
		sql.Named("optionImage", options.OptionImage),
		sql.Named("optionOrder", options.OptionOrder),
		sql.Named("isCorrect", options.IsCorrect),
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteOption(idQuestion int, options []int) error {
	if len(options) == 0 {
		return nil
	}

	placeholders := make([]string, len(options))
	args := []interface{}{idQuestion} // first arg is user ID

	for i, roleID := range options {
		placeholders[i] = fmt.Sprintf("@p%d", i+2) // @p2, @p3, ...
		args = append(args, roleID)
	}

	query := fmt.Sprintf(
		"DELETE FROM WebProject.QuestionOption WHERE idQuestion = @p1 AND idOption IN (%s)",
		strings.Join(placeholders, ","),
	)

	_, err := s.db.Exec(query, args...)
	return err
}

func (s *Store) UpdateOption(option types.Option) error {
	_, err := s.db.Exec(

		"UPDATE WebProject.QuestionOption SET "+
			"optionText=@optionText, optionImage=@optionImage, optionOrder=@optionOrder, isCorrect=@isCorrect"+
			" WHERE idOption=@idOption",
		sql.Named("optionText", option.OptionText),
		sql.Named("optionImage", option.OptionImage),
		sql.Named("optionOrder", option.OptionOrder),
		sql.Named("isCorrect", option.IsCorrect),
		sql.Named("idOption", option.IDOption),
	)

	if err != nil {
		return err
	}

	return nil
}
