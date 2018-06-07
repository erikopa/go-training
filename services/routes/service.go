package routes

import (
	"database/sql"
	"fmt"
)

// ServiceIface ...
type ServiceIface interface {
	List(limit int, offset int) (Routes, error)
}

// Service ...
type Service struct {
	db *sql.DB
}

// New ..
func New(db *sql.DB) *Service{
	return &Service{
		db: db,
	}
}

// List ...
func (s Service) List(limit int, offset int) (Routes, error){

	routes := Routes{}

	result, err := s.db.Query("select route_key, modality_id, threshold, create_at, update_at from routes limit ? offset ?", limit, offset)

	if err != nil{
		fmt.Println(err.Error())
		return routes, err
	}

	for result.Next(){
		row := Route{}
		err = result.Scan(&row.ID, &row.Modality, &row.Threshold, &row.CreateAt, &row.UpdateAt)

		if err != nil{
			fmt.Println(err.Error())
			return routes, err
		}

		routes = append(routes, row)
	}

	return routes, nil
}
