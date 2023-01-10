package models

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DefaultModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
type DBEnv struct {
	DBUSER string
	DBPASS string
	DBURI  string
	DBNAME string
}

type Model struct {
	*gorm.DB
}

func OpenDatabase() (*Model, error) {
	env := DBEnv{
		DBUSER: os.Getenv("DBUSER"),
		DBPASS: os.Getenv("DBPASS"),
		DBURI:  os.Getenv("DBURI"),
		DBNAME: os.Getenv("DBNAME"),
	}
	connectionURI := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v?charset=utf8mb4&parseTime=True&loc=Local", env.DBUSER, env.DBPASS, env.DBURI, env.DBNAME)
	db, err := gorm.Open(mysql.Open(connectionURI), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.Logger.LogMode(logger.Info)
	return &Model{db}, nil
}

var sharedInstance = newSingle()

func newSingle() *Model {
	db, err := OpenDatabase()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &Model{
		db.DB,
	}
}

func GetDBInstance() *Model {
	return sharedInstance
}
