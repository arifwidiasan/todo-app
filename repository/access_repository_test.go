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

func TestCreateAccessFail(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)

	fMock.ExpectQuery(regexp.QuoteMeta("INSERT INTO accesses \\(access_owner, user_id, activity_id\\) VALUES \\(\\?, \\?, \\?\\)"))

	data := model.Access{}

	res := repo.CreateAccess(data)
	assert.Error(t, res)
	defer dbMock.Close()
}

func TestCheckAccess(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `accesses`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "access_owner", "user_id", "activity_id"}))

	_, err := repo.CheckAccess(1, 1)
	assert.Error(t, err)
	defer dbMock.Close()
}

func TestCheckAccessOwner(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `accesses`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "access_owner", "user_id", "activity_id"}))

	_, err := repo.CheckAccessOwner(1, 1)
	assert.Error(t, err)
	defer dbMock.Close()
}

func TestDeleteAllAccess(t *testing.T) {
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

	err := repo.DeleteAllAccess(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestGetAccessUserActivity(t *testing.T) {
	dbMock, _, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)

	res := repo.GetAccessUserActivity(1)
	assert.Equal(t, res, []model.ListAccess([]model.ListAccess{}))
	defer dbMock.Close()
}

func TestDeleteOneAccess(t *testing.T) {
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

	err := repo.DeleteOneAccess(1, 1)
	assert.Error(t, err)
	assert.True(t, true)
}
