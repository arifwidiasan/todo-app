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

func TestCreateTaskFail(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)

	fMock.ExpectQuery(regexp.QuoteMeta("INSERT INTO tasks \\(task_name, task_priority, task_done, activity_id\\) VALUES \\(\\?, \\?, \\?, \\?\\)"))

	data := model.Task{}

	res := repo.CreateTask(data)
	assert.Error(t, res)
	defer dbMock.Close()
}

func TestDeleteAllTask(t *testing.T) {
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

	err := repo.DeleteAllTask(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestGetAlltask(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `tasks`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "task_name", "task_priority", "task_done", "activity_id"}).
			AddRow(1, "task 1", "1", false, 2).
			AddRow(2, "task 2", "2", false, 1))

	res := repo.GetAllTask(1)
	assert.Equal(t, res[0].Task_Name, "task 1")
	defer dbMock.Close()
}

func TestGetTaskByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `tasks`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "task_name", "task_priority", "task_done", "activity_id"}).
			AddRow(1, "task 1", "1", false, 2).
			AddRow(2, "task 2", "2", false, 1))

	res, _ := repo.GetTaskByID(1)
	assert.Equal(t, res.Task_Name, "task 1")
	defer dbMock.Close()
}

func TestGetTaskByIDFail(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `tasks`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "task_name", "task_priority", "task_done", "activity_id"}))

	res, err := repo.GetTaskByID(1)
	assert.Empty(t, res)
	assert.Error(t, err)
	defer dbMock.Close()
}

func TestUpdateTaskByID(t *testing.T) {
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

	err := repo.UpdateTaskByID(1, model.Task{
		Task_Name: "abc",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateTaskByIDFail(t *testing.T) {
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

	err := repo.UpdateTaskByID(1, model.Task{
		Task_Name: "def",
	})
	assert.Error(t, err)
	assert.True(t, true)
}

func TestDeleteTaskByID(t *testing.T) {
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

	err := repo.DeleteTaskByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestDeleteTaskByIDFail(t *testing.T) {
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

	err := repo.DeleteTaskByID(2)
	assert.Error(t, err)
	assert.True(t, true)
}

func TestCompleteTaskByIDFail(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewMysqlRepository(db)

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `tasks`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "task_name", "task_priority", "task_done", "activity_id"}).
			AddRow(1, "task 1", "1", false, 2).
			AddRow(2, "task 2", "2", false, 1))

	res := repo.CompleteTaskByID(1, model.Task{})
	assert.Error(t, res)
	defer dbMock.Close()
}
