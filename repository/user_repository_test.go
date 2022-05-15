package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arifwidiasan/todo-app/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetUserByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "user_email", "user_pass"}).
			AddRow(1, "arif", "arif@yahoo.com", "arifws"))

	res, _ := repo.GetUserByID(1)
	assert.Equal(t, res.Username, "arif")
	dbMock.Close()
}

func TestGetUserByIDFail(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "user_email", "user_pass"}))

	res, err := repo.GetUserByID(1)
	assert.Empty(t, res)
	assert.Error(t, err)
	dbMock.Close()
}

func TestGetUserByUsername(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "user_email", "user_pass"}).
			AddRow(1, "arif", "arif@yahoo.com", "arifws"))

	res, _ := repo.GetUserByUsername("arif")
	assert.Equal(t, res.Username, "arif")
	dbMock.Close()
}

func TestGetUserByUsernameFail(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "user_email", "user_pass"}))

	res, err := repo.GetUserByUsername("arif")
	assert.Empty(t, res)
	assert.Error(t, err)
	dbMock.Close()
}

func TestCreateUsersFail(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)

	fMock.ExpectQuery(regexp.QuoteMeta("INSERT INTO users \\(username, user_email, user_pass\\) VALUES \\(\\?, \\?, \\?, \\?\\)"))

	data := model.User{}

	res := repo.CreateUsers(data)
	assert.Error(t, res)
	dbMock.Close()
}
