package configuration

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

type configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

func GetConfiguration() configuration {

	var c configuration
	file, err := os.Open("./config.json")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

// GetConnection Obtiene una conexion a la db
func GetConnection() *gorm.DB {
	c := GetConfiguration()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Server, c.Port, c.Database)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
