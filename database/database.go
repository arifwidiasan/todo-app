package database

import (
	"fmt"

	"github.com/arifwidiasan/todo-app/config"
	"github.com/arifwidiasan/todo-app/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(conf config.Config) *gorm.DB {

	conectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)
	DB, err := gorm.Open(mysql.Open(conectionString), &gorm.Config{})
	if err != nil {
		fmt.Println("error open conection : ", err)
	}

	priority := DB.Migrator().HasTable(&model.Task_Priority{})
	if !priority {
		DB.Migrator().CreateTable(&model.Task_Priority{})
		DB.Model(&model.Task_Priority{}).Create([]map[string]interface{}{
			{"task_priority_name": "Tidak penting dan tidak mendesak"},
			{"task_priority_name": "Mendesak tapi tidak penting"},
			{"task_priority_name": "Penting tapi tidak mendesak"},
			{"task_priority_name": "Penting dan mendesak"},
		})
	}
	DB.AutoMigrate(&model.User{}, &model.Activity{}, &model.Access{}, &model.Task{})
	return DB
}
