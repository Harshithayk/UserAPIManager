package dbconnection

import (
	"todo/pck/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Dbconnect() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=password dbname=todo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	err = sqlDb.Ping()
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Users{}, &models.FetchByID{}, &models.FetchUser{}, models.Login{}, models.UsersModel{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
