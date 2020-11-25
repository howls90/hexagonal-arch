package dataSource

import (
	"gorm.io/gorm"

	"os"

	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func InitMysql() (err error) {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	db := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + db + "?charset=utf8mb4&parseTime=True&loc=Local"
	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err == nil {
		err := DB.AutoMigrate(&UserMsql{})
		return err
	}

	return err
}
