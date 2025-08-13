package permission

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

func (s *Store) GetPermissionByID(id int) (*types.Permission, error) {
	row := s.db.QueryRow("SELECT * FROM WebProject.Permission WHERE idPermission = @idPermission", sql.Named("idPermission", id))

	r := new(types.Permission)
	err := row.Scan(&r.IDPermission, &r.PermissionName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("permission not found")
		}
		return nil, err
	}

	return r, nil
}

func scanRowsIntoPermission(rows *sql.Rows) (*types.Permission, error) {
	permission := new(types.Permission)

	err := rows.Scan(
		&permission.IDPermission,
		&permission.PermissionName,
	)

	if err != nil {
		return nil, err
	}

	return permission, nil
}

func (s *Store) GetPermissions() ([]types.Permission, error) {
	rows, err := s.db.Query("SELECT * FROM WebProject.Permission")
	if err != nil {
		return nil, err
	}

	permissions := make([]types.Permission, 0)
	for rows.Next() {
		r, err := scanRowsIntoPermission(rows)
		if err != nil {
			return nil, err
		}

		permissions = append(permissions, *r)
	}

	return permissions, nil
}

func (s *Store) CreatePermission(permission types.Permission) error {
	_, err := s.db.Exec(

		"INSERT INTO WebProject.Permission"+
			"(permissionName)"+
			"VALUES(@permissionName)",
		sql.Named("permissionName", permission.PermissionName),
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeletePermission(id int) error {
	_, err := s.db.Exec(

		"DELETE FROM WebProject.Permission WHERE idPermission=@idPermission",
		sql.Named("idPermission", id),
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdatePermission(permission types.Permission) error {
	_, err := s.db.Exec(

		"UPDATE WebProject.Permission SET "+
			"permissionName=@permissionName"+" WHERE idPermission=@idPermission",
		sql.Named("permissionName", permission.PermissionName),
		sql.Named("idPermission", permission.IDPermission),
	)

	if err != nil {
		return err
	}

	return nil
}
