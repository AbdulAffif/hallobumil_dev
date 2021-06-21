package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"

	//  _ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func SetupCon(Dbdriver string) (db *gorm.DB) {

	var DBURL string
	var err error
	errENV := godotenv.Load()
	if errENV != nil {
		log.Fatalf("Error getting env, not comming through %v", errENV)
	} else {
		fmt.Println("We are getting the env values")
	}
	strUSer := fmt.Sprintf("DB_USER_%s", Dbdriver)
	strPass := fmt.Sprintf("DB_PASSWORD_%s", Dbdriver)
	strHost := fmt.Sprintf("DB_HOST_%s", Dbdriver)
	strPort := fmt.Sprintf("DB_PORT_%s", Dbdriver)
	strName := fmt.Sprintf("DB_NAME_%s", Dbdriver)

	DbUser := os.Getenv(strUSer)
	DbPassword := os.Getenv(strPass)
	DbHost := os.Getenv(strHost)
	DbPort := os.Getenv(strPort)
	DbName := os.Getenv(strName)

	if Dbdriver == "MYSQL" {
		DBURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	}
	if Dbdriver == "POSTGRES" {
		DBURL = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	}
	if Dbdriver == "MSSQL" {
		DBURL = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", DbUser, DbPassword, DbHost, DbPort, DbName)
	}
	driver := strings.ToLower(Dbdriver)
	db, err = gorm.Open(driver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", driver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", driver)
	}
	db.SingularTable(true)
	//db.AutoMigrate(entity.User{}, entity.Ud_access{}, entity.User_data{})
	return db
}

func CloseConDB(db *gorm.DB) {
	db.Close()
}
