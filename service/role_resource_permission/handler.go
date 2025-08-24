package role_resource_permission

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

func (s *Store) GetRoleResourcePermissionByID(id int) (*types.RoleResourcePermission, error) {
	row := s.db.QueryRow("SELECT * FROM WebProject.RoleResourcePermission WHERE idRoleResourcePermission = @idRoleResourcePermission", sql.Named("idRoleResourcePermission", id))

	r := new(types.RoleResourcePermission)
	err := row.Scan(&r.IDRoleResourcePermission, &r.IDRole, &r.IDResourcePermission)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("role resource permission not found")
		}
		return nil, err
	}

	return r, nil
}

func scanRowsIntoRoleResourcePermission(rows *sql.Rows) (*types.RoleResourcePermission, error) {
	roleResourcePermission := new(types.RoleResourcePermission)

	err := rows.Scan(
		&roleResourcePermission.IDRoleResourcePermission,
		&roleResourcePermission.IDRole,
		&roleResourcePermission.IDResourcePermission,
	)

	if err != nil {
		return nil, err
	}

	return roleResourcePermission, nil
}

func (s *Store) GetRoleResourcePermissions() ([]types.RoleResourcePermission, error) {
	rows, err := s.db.Query("SELECT * FROM WebProject.RoleResourcePermission")
	if err != nil {
		return nil, err
	}

	roleResourcePermission := make([]types.RoleResourcePermission, 0)
	for rows.Next() {
		r, err := scanRowsIntoRoleResourcePermission(rows)
		if err != nil {
			return nil, err
		}

		roleResourcePermission = append(roleResourcePermission, *r)
	}

	return roleResourcePermission, nil
}

func (s *Store) CreateRoleResourcePermission(roleResourcePermission types.RoleResourcePermission) error {
	_, err := s.db.Exec(

		"INSERT INTO WebProject.RoleResourcePermission"+
			"(idRole, idResourcePermission)"+
			"VALUES(@idRole, @idResourcePermission)",
		sql.Named("idRole", roleResourcePermission.IDRole),
		sql.Named("idResourcePermission", roleResourcePermission.IDResourcePermission),
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteRoleResourcePermission(id int) error {
	_, err := s.db.Exec(

		"DELETE FROM WebProject.RoleResourcePermission WHERE idRoleResourcePermission=@idRoleResourcePermission",
		sql.Named("idRoleResourcePermission", id),
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateRoleResourcePermission(roleResourcePermission types.RoleResourcePermission) error {
	_, err := s.db.Exec(

		"UPDATE WebProject.RoleResourcePermission SET "+
			"idRole=@idRole, idResourcePermission=@idResourcePermission"+" WHERE idRoleResourcePermission=@idRoleResourcePermission",
		sql.Named("idRole", roleResourcePermission.IDRole),
		sql.Named("idResourcePermission", roleResourcePermission.IDResourcePermission),
		sql.Named("idRoleResourcePermission", roleResourcePermission.IDRoleResourcePermission),
	)

	if err != nil {
		return err
	}

	return nil
}
