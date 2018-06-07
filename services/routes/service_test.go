package routes_test

import (
	"testing"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestListOk(t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil{
		t.Fatalf("an error '%s' was not expected when open stub database connection", err)
		t.Fail()
	}

	columns := []string{"route_key", "modality_id", "threshold", "create_date", "update_date"}
	mock.ExpectQuery("select ").WithArgs(50,0).WillReturnRows(sqlmock.NewRows(columns).AddRow("400-20000000-23799"))
}
