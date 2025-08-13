package product

import (
	"database/sql"

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

func scanRowsIntoQuiz(rows *sql.Rows) (*types.Quiz, error) {
	product := new(types.Quiz)

	err := rows.Scan(
		&product.IDQuiz,
		&product.QuizName,
		&product.Description,
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
		"(idQuiz, quizName, description, creationDate, hasTimer, timer, hasLifeline, idType, idCreator, idCategory, idLanguage, avgRating, isNSFW, uniquePlays, plays)"+
		"VALUES(@idQuiz, @quizName, @description, @creationDate, @hasTimer, @timer, @hasLifeline, @idType, @idCreator, @idCategory, @idLanguage, @avgRating, @isNSFW, @uniquePlays, @plays)",
		sql.Named("idQuiz", quiz.IDQuiz),
		sql.Named("quizName", quiz.QuizName),
		sql.Named("description", quiz.Description),
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
