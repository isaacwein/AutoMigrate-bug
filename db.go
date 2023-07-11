package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func dbSetup() (db *gorm.DB, err error) {

	if db, err = OpenTestConnection(); err != nil {
		log.Printf("failed to connect database, got error %v\n", err)

	} else {
		sqlDB, err := db.DB()
		if err == nil {
			err = sqlDB.Ping()
		}

		if err != nil {
			log.Printf("failed to connect database, got error %v\n", err)
		}

		if db.Dialector.Name() == "sqlite" {
			db.Exec("PRAGMA foreign_keys = ON")
		}

	}
	return
}

func OpenTestConnection() (db *gorm.DB, err error) {

	dbDSN := os.Getenv("GORM_DSN")

	switch os.Getenv("GORM_DIALECT") {
	case "mysql":
		log.Println("testing mysql...")
		if dbDSN == "" {
			dbDSN = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
		}
		db, err = gorm.Open(mysql.Open(dbDSN), &gorm.Config{})
	case "mariadb":
		log.Println("testing mariadb...")
		if dbDSN == "" {
			dbDSN = "gorm:gorm@tcp(localhost:9950)/gorm?charset=utf8&parseTime=True&loc=Local"
		}
		db, err = gorm.Open(mysql.Open(dbDSN), &gorm.Config{})
		//case "postgres":
		//	log.Println("testing postgres...")
		//	if dbDSN == "" {
		//		dbDSN = "user=gorm password=gorm host=localhost dbname=gorm port=9920 sslmode=disable TimeZone=America/New_York"
		//	}
		//	db, err = gorm.Open(postgres.Open(dbDSN), &gorm.Config{})
		//
		//case "sqlserver":
		//	// CREATE LOGIN gorm WITH PASSWORD = 'LoremIpsum86';
		//	// CREATE DATABASE gorm;
		//	// USE gorm;
		//	// CREATE USER gorm FROM LOGIN gorm;
		//	// sp_changedbowner 'gorm';
		//	log.Println("testing sqlserver...")
		//	if dbDSN == "" {
		//		dbDSN = "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
		//	}
		//	db, err = gorm.Open(sqlserver.Open(dbDSN), &gorm.Config{})
		//default:
		//	log.Println("testing sqlite3...")
		//	db, err = gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), &gorm.Config{})
	}

	if debug := os.Getenv("DEBUG"); debug == "true" {
		db.Logger = db.Logger.LogMode(logger.Info)
	} else if debug == "false" {
		db.Logger = db.Logger.LogMode(logger.Silent)
	}

	return
}
