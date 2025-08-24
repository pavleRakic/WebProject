package role

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

func (s *Store) GetRoleByID(id int) (*types.Role, error) {
	row := s.db.QueryRow("SELECT * FROM WebProject.Role WHERE idRole = @idRole", sql.Named("idRole", id))

	r := new(types.Role)
	err := row.Scan(&r.IDRole, &r.RoleName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("role not found")
		}
		return nil, err
	}

	return r, nil
}

func scanRowsIntoRole(rows *sql.Rows) (*types.Role, error) {
	role := new(types.Role)

	err := rows.Scan(
		&role.IDRole,
		&role.RoleName,
	)

	if err != nil {
		return nil, err
	}

	return role, nil
}

func (s *Store) GetRoles() ([]types.Role, error) {
	rows, err := s.db.Query("SELECT * FROM WebProject.Role")
	if err != nil {
		return nil, err
	}

	roles := make([]types.Role, 0)
	for rows.Next() {
		r, err := scanRowsIntoRole(rows)
		if err != nil {
			return nil, err
		}

		roles = append(roles, *r)
	}

	return roles, nil
}

func (s *Store) CreateRole(role types.Role) error {
	_, err := s.db.Exec(

		"INSERT INTO WebProject.Role"+
			"(roleName)"+
			"VALUES(@roleName)",
		sql.Named("roleName", role.RoleName),
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteRole(id int) error {
	_, err := s.db.Exec(

		"DELETE FROM WebProject.Role WHERE idRole=@idRole",
		sql.Named("idRole", id),
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateRole(role types.Role) error {
	_, err := s.db.Exec(

		"UPDATE WebProject.Role SET "+
			"roleName=@roleName"+" WHERE idRole=@idRole",
		sql.Named("roleName", role.RoleName),
		sql.Named("idRole", role.IDRole),
	)

	if err != nil {
		return err
	}

	return nil
}
