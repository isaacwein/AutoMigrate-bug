package main

import (
	"log"
)

type Users struct {
	// id and customer are the primary key because if a call goes from one customer to another, it will be a new row in the database
	ID        string `json:"id" gorm:"column:id; size:32; index:id_i; index:pk; primaryKey; autoIncrement:true; not null;"`
	Customer  int    `json:"customer" gorm:"column:customer; index:customer_i; index:pk; primaryKey; autoIncrement:false; not null"`
	FirstName string `json:"first_name" gorm:"column:first_name; size:255; index:first_name_i; not null"`
	LastName  string `json:"last_name" gorm:"column:last_name; size:255; index:last_name_i; not null"`
	// Segments is here so that gorm creates the foreign key between calls and segments.
	Mobiles []*UsersMobiles `json:"parties" gorm:"foreignKey: user_id,customer_id; references: id,customer; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Emails  []*UsersEmails  `json:"tokens" gorm:"foreignKey: user_id,customer_id; references: id,customer; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (*Users) TableName() string {
	return "people"
}

type UsersEmails struct {
	// Primary key is Readonly because the db should create it not gorm.
	ID         uint64 `gorm:"column:id; index:id_i; primaryKey;<-:false"`
	UserID     string `gorm:"column:user_id; index:callid_i; index:fk_i; size:32; not null"`
	CustomerID int    `gorm:"column:customer_id; index:customer_id_i; index:fk_i; not null"`
	Email      string `gorm:"column:email; size:255; index:email_i; not null"`
	//Color        int    `gorm:"column:color; type: mediumint(8) unsigned; not null; index:color_i; "`
}

func (UsersEmails) TableName() string {
	return "people_email"
}

type UsersMobiles struct {
	ID         uint64 `gorm:"column:id; index:id_i; primaryKey;<-:false"`
	UserID     string `gorm:"column:user_id; index:callid_i; index:fk_i; size:32; not null"`
	CustomerID int    `gorm:"column:customer_id; index:customer_id_i; index:fk_i; not null"`
	Mobile     string `gorm:"column:mobile; size:255; index:mobile_i; not null"`
}

func (UsersMobiles) TableName() string {
	return "people_mobile"
}

func main() {

	db := DB

	err := db.AutoMigrate(UsersMobiles{})
	if err != nil {
		log.Fatalln("error code (2.1387) failed to migrate UsersMobiles", err.Error())
	}
	err = db.AutoMigrate(UsersEmails{})
	if err != nil {
		log.Fatalln("error code (2.1387) failed to migrate UsersEmails", err.Error())
	}

	err = db.AutoMigrate(Users{})
	if err != nil {
		log.Fatalln("error code (2.1385) failed to migrate Users", err.Error())
	}
	return

}
