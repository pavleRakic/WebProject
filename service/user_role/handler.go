package user_role

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/pavleRakic/testGoApi/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserRoleByID(idUser int, idRole int) (*types.UserRole, error) {
	row := s.db.QueryRow("SELECT * FROM WebProject.UserRole WHERE idUser = @idUser AND idRole = @idRole", sql.Named("idUser", idUser), sql.Named("idRole", idRole))

	r := new(types.UserRole)
	err := row.Scan(&r.IDUser, &r.IDRole)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user doesn't have that role")
		}
		return nil, err
	}

	return r, nil
}

func scanRowsIntoUserRole(rows *sql.Rows) (*types.UserRole, error) {
	userRole := new(types.UserRole)

	err := rows.Scan(
		&userRole.IDUser,
		&userRole.IDRole,
	)

	if err != nil {
		return nil, err
	}

	return userRole, nil
}

func (s *Store) GetUsersRoles() ([]types.UserRole, error) {
	rows, err := s.db.Query("SELECT * FROM WebProject.UserRole")
	if err != nil {
		return nil, err
	}

	userRole := make([]types.UserRole, 0)
	for rows.Next() {
		r, err := scanRowsIntoUserRole(rows)
		if err != nil {
			return nil, err
		}

		userRole = append(userRole, *r)
	}

	return userRole, nil
}

func (s *Store) GetUserRoles(idUser int) ([]types.UserRole, error) {
	rows, err := s.db.Query("SELECT * FROM WebProject.UserRole WHERE idUser = @idUser", sql.Named("idUser", idUser))
	if err != nil {
		return nil, err
	}

	userRole := make([]types.UserRole, 0)
	for rows.Next() {
		r, err := scanRowsIntoUserRole(rows)
		if err != nil {
			return nil, err
		}

		userRole = append(userRole, *r)
	}

	return userRole, nil
}

func (s *Store) AssignUserRoles(idUser int, roles []string) error {

	values := []string{}
	args := []interface{}{}

	paramIndex := 1
	for _, roleID := range roles {
		values = append(values, fmt.Sprintf("(@p%d, @p%d)", paramIndex, paramIndex+1))
		args = append(args, idUser, roleID)
		paramIndex += 2
	}

	query := "INSERT INTO WebProject.UserRole (idUser, idRole) VALUES " + strings.Join(values, ",")
	_, err := s.db.Exec(query, args...)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *Store) UnassignUserRoles(idUser int, roles []string) error {

	if len(roles) == 0 {
		return nil
	}

	placeholders := make([]string, len(roles))
	args := []interface{}{idUser} // first arg is user ID

	for i, roleID := range roles {
		placeholders[i] = fmt.Sprintf("@p%d", i+2) // @p2, @p3, ...
		args = append(args, roleID)
	}

	query := fmt.Sprintf(
		"DELETE FROM WebProject.UserRole WHERE idUser = @p1 AND idRole IN (%s)",
		strings.Join(placeholders, ","),
	)

	_, err := s.db.Exec(query, args...)
	return err
}
