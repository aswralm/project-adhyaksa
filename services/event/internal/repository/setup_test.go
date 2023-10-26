package repository_test

import (
	"fmt"
	"log"
	"project-adhyaksa/pkg/config"
	"project-adhyaksa/pkg/database"
	"project-adhyaksa/services/event/internal/repository/model"
)

func SetupTest() *config.Config {
	//setup config
	config, err := config.NewConfigV2("./")
	if err != nil {
		log.Fatal("config error ", err)
	}

	//setup db
	db := database.ConnectMYSQL("mysql", config)
	config.Db = db

	//setup gorm
	gormDB := database.ConnectGORM("mysql", config)
	config.GormDB = gormDB

	return config
}

func refreshEventTable(config *config.Config, tables map[string]interface{}) {
	db := config.GormDB
	for tableName, table := range tables {
		if err := db.Table(tableName).Where("deleted_at IS NULL").Delete(table).Error; err != nil {
			log.Println(err)
			break
		} else {
			log.Println(fmt.Sprintf("refresh %s success", tableName))
		}
	}

}

func insertTable(config *config.Config, tables map[string]interface{}) {
	db := config.GormDB

	for tableName, table := range tables {
		switch {
		case tableName == model.Branch{}.GetTableName():
			data := table.(model.Branch)
			if err := db.Table(tableName).Create(&data).Error; err != nil {
				log.Println(err)
			} else {
				log.Println("insert data branch success")
			}
		case tableName == model.Event{}.GetTableName():
			data := table.(model.Event)
			if err := db.Table(tableName).Create(&data).Error; err != nil {
				log.Println(err)
			} else {
				log.Println("insert data event success")
			}
		case tableName == model.Documentation{}.GetTableName():
			data := table.(model.Documentation)
			if err := db.Table(tableName).Create(&data).Error; err != nil {
				log.Println(err)
			} else {
				log.Println("insert data documentation success")
			}
		case tableName == model.Photo{}.GetTableName():
			data := table.(model.Photo)
			if err := db.Table(tableName).Create(&data).Error; err != nil {
				log.Println(err)
			} else {
				log.Println("insert data photo success")
			}
		}

	}
}
