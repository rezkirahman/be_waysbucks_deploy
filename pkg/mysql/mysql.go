package mysql

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var DB_HOST = os.Getenv("DB_HOST")
var DB_USER = os.Getenv("DB_USER")
var DB_PASSWORD = os.Getenv("DB_PASSWORD")
var DB_NAME = os.Getenv("DB_NAME")
var DB_PORT = os.Getenv("DB_PORT")

// Connection Database
func DatabaseInit() {
	var err error
	//dsn := "root:@tcp(127.0.0.1:3306)/waysbucks?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
  dsn := "root:eIoIW3ErHby3rxGFR9vK@tcp(containers-us-west-146.railway.app:6648)/railway?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Database")
}
// DB_HOST=containers-us-west-146.railway.app
// DB_USER=root
// DB_PASSWORD=eIoIW3ErHby3rxGFR9vK
// DB_NAME=railway
// DB_PORT=6648