package repository

//import (
//	"fmt"
//	"github.com/DATA-DOG/go-sqlmock"
//	"github.com/labstack/echo/v4"
//	"github.com/stretchr/testify/require"
//	"regexp"
//	"testing"
//)
//
//func TestPostgres_CheckExistUserByEmail_True(t *testing.T) {
//	// Init sqlmock
//	db, mock, err := sqlmock.New()
//	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
//	defer db.Close()
//
//	// Global input output
//	inputEmail := "CorrectEmail@mail.ru"
//
//	// Input global
//	var ctx echo.Context
//
//	// Create required setup for handling
//	rowMain := sqlmock.NewRows([]string{"exist"})
//
//	rowMain.AddRow(true)
//
//	// Settings mock
//	mock.
//		ExpectQuery(regexp.QuoteMeta(auth.CheckExist)).
//		WithArgs(inputEmail). // Values in query
//		WillReturnRows(rowMain)
//
//	// Init
//	repo := auth.NewAuthDatabase(&sqltools.Database{Connection: db})
//
//	// Check result
//	actual, err := repo.CheckExistUserByEmail(ctx, inputEmail)
//	require.Nil(t, err, fmt.Errorf("unexpected err: %s", err))
//
//	err = mock.ExpectationsWereMet()
//	require.Nil(t, err, fmt.Errorf("there were unfulfilled expectations: %s", err))
//
//	// Check actual
//	require.Equal(t, true, actual)
//}
