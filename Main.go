package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func main() {
	fmt.Println("Microservice For User Management")
	migrateDB(dba())

}


//SETUP DB
var dba = func() (*gorm.DB){
	//dbs := DBToken{"mysql", "mysql", "21484eb879a418bb","token-service","tcp(172.17.0.2:3306)",nil}
	dbs := DBToken{"mysql", "root", "","service_user","tcp(127.0.0.1:3306)",nil}
	dbs.getConnection()
	return dbs.dbToken
}

//DB MIGRATE
func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&User{},&Additional{},&Role{})
}
