package repository

import (
	"database/sql/driver"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arifwidiasan/todo-app/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreateActivitysFail(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)

	fMock.ExpectQuery(regexp.QuoteMeta("INSERT INTO activity \\(username, user_email, user_pass\\) VALUES \\(\\?, \\?, \\?, \\?\\)"))

	data := model.Activity{}

	res := repo.CreateActivity(data)
	assert.Error(t, res)
	defer dbMock.Close()
}

func TestGetLatestActivity(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `activities`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "activity_name", "archived"}).
			AddRow(1, "tes 1", false).
			AddRow(2, "tes 2", false))

	res, _ := repo.GetLatestActivity()
	assert.Equal(t, res.Activity_Name, "tes 1")
	defer dbMock.Close()
}

func TestGetLatestActivityFail(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `activities`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "activity_name", "archived"}))

	res, err := repo.GetLatestActivity()
	assert.Empty(t, res)
	assert.Error(t, err)
	defer dbMock.Close()
}

func TestGetAllActivity(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `activities`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "activity_name", "archived"}).
			AddRow(1, "tes 1", false).
			AddRow(2, "tes 2", false))

	res := repo.GetAllActivity(1)
	assert.Equal(t, res, []model.Activity([]model.Activity{}))
	defer dbMock.Close()
}

func TestGetAllArchiveActivity(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `activities`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "activity_name", "archived"}).
			AddRow(1, "tes 1", false).
			AddRow(2, "tes 2", true))

	res := repo.GetAllArchiveActivity(1)
	assert.Equal(t, res, []model.Activity([]model.Activity{}))
	defer dbMock.Close()
}

func TestGetActivityByName(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `activities`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "activity_name", "archived"}).
			AddRow(1, "tes 1", false).
			AddRow(2, "tes 2", false))

	res, _ := repo.GetActivityByName("tes1")
	assert.Equal(t, res.Activity_Name, "tes 1")
	defer dbMock.Close()
}

func TestGetActivityByNameFail(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `activities`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "activity_name", "archived"}))

	res, err := repo.GetActivityByName("tes1")
	assert.Empty(t, res)
	assert.Error(t, err)
	defer dbMock.Close()
}

func TestArchiveActivityByIDFail(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `activities`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "activity_name", "archived"}).
			AddRow(1, "tes 1", false).
			AddRow(2, "tes 2", false))

	res := repo.ArchiveActivityByID(1, model.Activity{})
	assert.Error(t, res)
	defer dbMock.Close()
}

func TestUpdateActivityByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs("abc", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateActivityByID(1, model.Activity{
		Activity_Name: "abc",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateActivityByIDFail(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs("abc", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateActivityByID(1, model.Activity{
		Activity_Name: "def",
	})
	assert.Error(t, err)
	assert.True(t, true)
}

func TestDeleteActivityByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteActivityByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestDeleteActivityByIDFail(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteActivityByID(10)
	assert.Error(t, err)
	assert.True(t, true)
}
