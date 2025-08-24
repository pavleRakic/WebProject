package resource_permission

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

func (s *Store) GetResourcePermissionByID(id int) (*types.ResourcePermission, error) {
	row := s.db.QueryRow("SELECT * FROM WebProject.ResourcePermission WHERE idResourcePermission = @idResourcePermission", sql.Named("idResourcePermission", id))

	r := new(types.ResourcePermission)
	err := row.Scan(&r.IDResourcePermission, &r.IDResource, &r.IDPermission)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("resource permission not found")
		}
		return nil, err
	}

	return r, nil
}

func scanRowsIntoResourcePermission(rows *sql.Rows) (*types.ResourcePermission, error) {
	resourcePermission := new(types.ResourcePermission)

	err := rows.Scan(
		&resourcePermission.IDResourcePermission,
		&resourcePermission.IDResource,
		&resourcePermission.IDPermission,
	)

	if err != nil {
		return nil, err
	}

	return resourcePermission, nil
}

func (s *Store) GetResourcePermissions() ([]types.ResourcePermission, error) {
	rows, err := s.db.Query("SELECT * FROM WebProject.ResourcePermission")
	if err != nil {
		return nil, err
	}

	resourcePermission := make([]types.ResourcePermission, 0)
	for rows.Next() {
		r, err := scanRowsIntoResourcePermission(rows)
		if err != nil {
			return nil, err
		}

		resourcePermission = append(resourcePermission, *r)
	}

	return resourcePermission, nil
}

func (s *Store) CreateResourcePermission(resourcePermission types.ResourcePermission) error {
	_, err := s.db.Exec(

		"INSERT INTO WebProject.ResourcePermission"+
			"(idResource, idPermission)"+
			"VALUES(@idResource, @idPermission)",
		sql.Named("idResource", resourcePermission.IDResource),
		sql.Named("idPermission", resourcePermission.IDPermission),
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteResourcePermission(id int) error {
	_, err := s.db.Exec(

		"DELETE FROM WebProject.ResourcePermission WHERE idResourcePermission=@idResourcePermission",
		sql.Named("idResourcePermission", id),
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateResourcePermission(resourcePermission types.ResourcePermission) error {
	_, err := s.db.Exec(

		"UPDATE WebProject.ResourcePermission SET "+
			"idResource=@idResource, idPermission=@idPermission"+" WHERE idResourcePermission=@idResourcePermission",
		sql.Named("idResource", resourcePermission.IDResource),
		sql.Named("idPermission", resourcePermission.IDPermission),
		sql.Named("idResourcePermission", resourcePermission.IDResourcePermission),
	)

	if err != nil {
		return err
	}

	return nil
}
