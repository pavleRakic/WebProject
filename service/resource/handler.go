package resource

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

func (s *Store) GetResourceByID(id int) (*types.Resource, error) {
	row := s.db.QueryRow("SELECT * FROM WebProject.Resource WHERE idResource = @idResource", sql.Named("idResource", id))

	r := new(types.Resource)
	err := row.Scan(&r.IDResource, &r.ResourceName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return r, nil
}

func scanRowsIntoResource(rows *sql.Rows) (*types.Resource, error) {
	resource := new(types.Resource)

	err := rows.Scan(
		&resource.IDResource,
		&resource.ResourceName,
	)

	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (s *Store) GetResources() ([]types.Resource, error) {
	rows, err := s.db.Query("SELECT * FROM WebProject.Resource")
	if err != nil {
		return nil, err
	}

	resources := make([]types.Resource, 0)
	for rows.Next() {
		r, err := scanRowsIntoResource(rows)
		if err != nil {
			return nil, err
		}

		resources = append(resources, *r)
	}

	return resources, nil
}

func (s *Store) CreateResource(resource types.Resource) error {
	_, err := s.db.Exec(

		"INSERT INTO WebProject.Resource"+
			"(resourceName)"+
			"VALUES(@resourceName)",
		sql.Named("resourceName", resource.ResourceName),
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteResource(id int) error {
	_, err := s.db.Exec(

		"DELETE FROM WebProject.Resource WHERE idResource=@idResource",
		sql.Named("idResource", id),
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateResource(resource types.Resource) error {
	_, err := s.db.Exec(

		"UPDATE WebProject.Resource SET "+
			"resourceName=@resourceName"+" WHERE idResource=@idResource",
		sql.Named("resourceName", resource.ResourceName),
		sql.Named("idResource", resource.IDResource),
	)

	if err != nil {
		return err
	}

	return nil
}
