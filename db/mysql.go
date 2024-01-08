package db

import (
	"api-instagram/config"

	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getDatabaseConnectURL(cfg *config.Configuration) string {
	return cfg.MySQLUsername + ":" + cfg.MySQLPassword + "@tcp" + "(" + cfg.MySQLHost + ":" + cfg.MySQLPort + ")/" + cfg.MySQLDatabaseName + "?" + "parseTime=true&loc=Local"
}

func GetMySQLInstance(cfg *config.Configuration, migrate bool) *gorm.DB {
	dsn := getDatabaseConnectURL(cfg)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println("MySQL connected")
	}

	if migrate {
		// Run the models migration
		// db.AutoMigrate(
		//
		// )
	}
	return db
}
