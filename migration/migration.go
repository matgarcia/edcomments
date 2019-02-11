package migration

import (
	"github.com/matgarcia/edcomments/configuration"
	"github.com/matgarcia/edcomments/models"
	_ "github.com/go-sql-driver/mysql"
)

func Migrate() {

	db := configuration.GetConnection()
	defer db.Close()

	db.CreateTable(&models.User{})
	db.CreateTable(&models.Comment{})
	db.CreateTable(&models.Vote{})
	db.Model(&models.Vote{}).AddUniqueIndex("comment_ide_user_id_unique", "comment_id", "user_id")



}
